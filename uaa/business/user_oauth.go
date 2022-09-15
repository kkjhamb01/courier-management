package business

import (
	"context"
	"encoding/json"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	"gitlab.artin.ai/backend/courier-management/party/proto"
	pb "gitlab.artin.ai/backend/courier-management/uaa/proto"
	"gitlab.artin.ai/backend/courier-management/uaa/security"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)



func (s *RegistrationService) UserOauthRegister(ctx context.Context, in *pb.UserOauthRegisterRequest) (*pb.UserOauthRegisterResponse, error) {

	var user security.User

	ctxUser := ctx.Value("user")
	if ctxUser != nil {
		user = ctxUser.(security.User)
	}

	// Start Redis Operations
	key := "oauth_" + in.DeviceId
	fields := map[string]string{}
	fields[REDIS_KEY_OAUTH_TYPE] = strconv.FormatInt(int64(in.OauthType), 10)
	fields[REDIS_KEY_USER_TYPE] = strconv.FormatInt(USER_TYPE_USER, 10)
	if user.Id != "" {
		b, _ := json.Marshal(user)
		fields[REDIS_KEY_USER] = string(b)
	}
	if _,err := s.redis.HMSet(ctx, key, fields).Result(); err != nil{
		return nil, pb.Internal.Error(err)
	}
	if _,err := s.redis.Expire(ctx, key, time.Duration(s.config.Redis.OTPSessionExpireSeconds) * time.Second).Result(); err != nil{
		return nil, pb.Internal.Error(err)
	}
	// Finish Redis Operations

	var link string
	if in.OauthType == pb.OauthType_OAUTH_GOOGLE{
		link = s.googleUserConfig.AuthCodeURL("state")
	} else if in.OauthType == pb.OauthType_OAUTH_FACEBOOK{
		link = s.facebookUserConfig.AuthCodeURL("state")
	} else{
		return nil, pb.InvalidArgument.ErrorMsg("invalid authorization server type")
	}
	return &pb.UserOauthRegisterResponse{
		FollowLink: link,
	}, nil
}

func (s *RegistrationService) UserOauthRegisterVerify(ctx context.Context, in *pb.UserOauthRegisterVerifyRequest) (*pb.UserOauthRegisterVerifyResponse, error) {

	// Start Redis Operations
	key := "oauth_" + in.DeviceId
	redisMap,err := s.redis.HGetAll(ctx, key).Result()
	if err != nil {
		return nil, pb.InvalidCode.ErrorNoMsg()
	}
	oauthType,err := strconv.ParseInt(redisMap[REDIS_KEY_OAUTH_TYPE], 10, 32)
	userType,err := strconv.ParseInt(redisMap[REDIS_KEY_USER_TYPE], 10, 32)

	if err != nil {
		return nil, pb.InvalidCode.ErrorNoMsg()
	}
	// Finish Redis Operations

	if int(userType) != USER_TYPE_USER {
		return nil, pb.InvalidArgument.ErrorMsg("user is not of type user")
	}

	var tokenUser security.User

	if val, ok := redisMap[REDIS_KEY_USER]; ok {
		json.Unmarshal([]byte(val), &tokenUser)
	}

	var user *security.User

	var userOauthInfo *pb.UserOauthInfo

	var facebookId, googleId string

	var exchangeCodeErr error

	if int(oauthType) == int(pb.OauthType_OAUTH_GOOGLE){
		code, _ := url.QueryUnescape(in.Code)

		tok, err := s.googleUserConfig.Exchange(ctx, code)

		if err!=nil{
			logger.Errorf("Invalid exchange code", err)
			exchangeCodeErr = pb.InvalidCode.ErrorNoMsg()
		} else {
			idToken := tok.Extra("id_token").(string)

			user, err = s.jwtUtils.ValidateUnsigned(idToken, false)
			if err != nil {
				logger.Errorf("google token validation failed ", err)
				exchangeCodeErr = pb.InvalidCode.ErrorMsg(err.Error())
			} else {
				googleId = user.Id
				userOauthInfo = &pb.UserOauthInfo{
					Name:      user.Name,
					Email:     user.Email,
					FirstName: user.FirstName,
					LastName:  user.LastName,
				}
				user = &security.User{
					Id:    googleId,
					Email: userOauthInfo.Email,
					Name:  userOauthInfo.Name,
					FirstName: userOauthInfo.FirstName,
					LastName:  userOauthInfo.LastName,
				}
			}
		}
	} else 	if int(oauthType) == int(pb.OauthType_OAUTH_FACEBOOK) {
		code, _ := url.QueryUnescape(in.Code)

		tok, err := s.facebookUserConfig.Exchange(ctx, code)

		if err!=nil{
			logger.Errorf("Invalid exchange code", err)
			exchangeCodeErr = pb.InvalidCode.ErrorNoMsg()
		} else {
			resp, err := http.Get("https://graph.facebook.com/me?fields=email,name,gender&access_token=" + url.QueryEscape(tok.AccessToken))
			defer resp.Body.Close()
			if err != nil {
				logger.Errorf("Error in fetching info from facebook", err)
				exchangeCodeErr = pb.Internal.ErrorMsg("cannot fetch user info from Facebook")
			} else {
				response, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					logger.Errorf("Error in fetching info from facebook", err)
					exchangeCodeErr = pb.Internal.ErrorMsg("cannot fetch user info from Facebook")
				} else {
					logger.Debugf("fetched user info : %v", string(response))

					var userInfo facebookUserInfo

					err = json.Unmarshal(response, &userInfo)

					if err != nil {
						logger.Errorf("Error in decoding message from facebook", err)
						exchangeCodeErr = pb.Internal.ErrorMsg("cannot decode user info from Facebook")
					} else {
						facebookId = userInfo.Id
						userOauthInfo = &pb.UserOauthInfo{
							Name:      userInfo.Name,
							Email:     userInfo.Email,
							FirstName: userInfo.FirstName,
							LastName:  userInfo.LastName,
						}

						user = &security.User{
							Id:    userInfo.Id,
							Email: userInfo.Email,
							Name:  userInfo.Name,
							FirstName: userInfo.FirstName,
							LastName:  userInfo.LastName,
						}
					}
				}
			}
		}
	} else {
		return nil, pb.InvalidArgument.ErrorMsg("invalid authorization server type")
	}

	if exchangeCodeErr != nil {
		oauthResponse, err := s.getOauthResponseFromCache(ctx, in.GetDeviceId())
		if err != nil{
			return nil, err
		}
		if oauthResponse == nil{
			return nil, exchangeCodeErr
		}
		if int(oauthType) == int(pb.OauthType_OAUTH_GOOGLE){
			googleId = oauthResponse.Identifier
		} else if int(oauthType) == int(pb.OauthType_OAUTH_FACEBOOK){
			facebookId = oauthResponse.Identifier
		}
		user = &security.User{
			Id:    oauthResponse.Identifier,
			Email: oauthResponse.Info.Email,
			Name:  oauthResponse.Info.Name,
			FirstName: oauthResponse.Info.FirstName,
			LastName: oauthResponse.Info.LastName,
		}
		userOauthInfo = oauthResponse.Info
	}

	if int(oauthType) == int(pb.OauthType_OAUTH_GOOGLE) {
		user.Claims = []security.Claim{{
			ClaimType:  security.CLAIM_TYPE_GOOGLE_ID,
			Identifier: googleId,
		}}
	} else if int(oauthType) == int(pb.OauthType_OAUTH_FACEBOOK) {
		user.Claims = []security.Claim{{
			ClaimType:  security.CLAIM_TYPE_FACEBOOK_ID,
			Identifier: facebookId,
		}}
	}

	var newUser = true

	if tokenUser.Id == "" {
		user.Id = generateId(false)
	} else {
		user.Id = tokenUser.Id
		newUser = false
		for _,claim := range tokenUser.Claims {
			user.Claims = append(user.Claims, claim)
		}
	}

	for _,claim := range user.Claims{
		if claim.ClaimType == security.CLAIM_TYPE_GOOGLE_ID  {
			u,_ := s.partyAPI.FindUserByGoogleId(claim.Identifier, proto.UserType_USER_TYPE_PASSENGER)
			if u != nil && u.Id != ""{
				if u != nil && u.Id != ""{
					if err := s.insertOauthResponseIntoCache(ctx, oauthResponse{
						DeviceId: in.GetDeviceId(),
						OauthType: int64(pb.OauthType_OAUTH_GOOGLE),
						UserType: USER_TYPE_USER,
						Identifier: googleId,
						Info: userOauthInfo,
					}); err != nil {
						return nil, err
					}
					return nil, pb.AlreadyExists.ErrorMsg("already registered with this google id")
				}
				return nil, pb.AlreadyExists.ErrorMsg("already registered with this google id")
			}
		} else if claim.ClaimType == security.CLAIM_TYPE_FACEBOOK_ID  {
			u,_ := s.partyAPI.FindUserByFacebookId(claim.Identifier, proto.UserType_USER_TYPE_PASSENGER)
			if u != nil && u.Id != ""{
				if err := s.insertOauthResponseIntoCache(ctx, oauthResponse{
					DeviceId: in.GetDeviceId(),
					OauthType: int64(pb.OauthType_OAUTH_FACEBOOK),
					UserType: USER_TYPE_USER,
					Identifier: facebookId,
					Info: userOauthInfo,
				}); err != nil {
					return nil, err
				}
				return nil, pb.AlreadyExists.ErrorMsg("already registered with this facebook id")
			}
		}
	}

	user.Roles = []security.Role{security.Role_CLIENT}

	token, err := s.jwtUtils.GenerateToken(*user)

	if err != nil{
		logger.Errorf("cannot generate token", err)
		return nil, pb.Internal.Error(err)
	}

	if !newUser {
		validUser,_ := s.partyAPI.GetUserByUserId(user.Id, proto.UserType_USER_TYPE_PASSENGER)
		if validUser != nil && validUser.Id != "" {
			var identifier string
			var claimType proto.ClaimType

			if facebookId != "" {
				identifier = facebookId
				claimType = proto.ClaimType_CLAIM_TYPE_FACEBOOK_ID
			} else if googleId != "" {
				identifier = googleId
				claimType = proto.ClaimType_CLAIM_TYPE_GOOGLE_ID
			}

			if err := s.partyAPI.RegisterClaim(user.Id, claimType, identifier, proto.UserType_USER_TYPE_PASSENGER); err != nil{
				logger.Errorf("cannot register claim to party api", err)
				return nil, pb.Internal.ErrorMsg("cannot register claim to party api " + err.Error())
			}
		}
	}

	return &pb.UserOauthRegisterVerifyResponse{
		Token: token,
		Info: userOauthInfo,
	}, nil
}

func (s *RegistrationService) UserOauthLogin(ctx context.Context, in *pb.UserOauthLoginRequest) (*pb.UserOauthLoginResponse, error) {
	// Start Redis Operations
	key := "oauth_" + in.DeviceId
	fields := map[string]string{}
	fields[REDIS_KEY_OAUTH_TYPE] = strconv.FormatInt(int64(in.OauthType), 10)
	fields[REDIS_KEY_USER_TYPE] = strconv.FormatInt(USER_TYPE_USER, 10)
	if _,err := s.redis.HMSet(ctx, key, fields).Result(); err != nil{
		return nil, pb.Internal.Error(err)
	}
	if _,err := s.redis.Expire(ctx, key, time.Duration(s.config.Redis.OTPSessionExpireSeconds) * time.Second).Result(); err != nil{
		return nil, pb.Internal.Error(err)
	}
	// Finish Redis Operations

	var link string
	if in.OauthType == pb.OauthType_OAUTH_GOOGLE{
		link = s.googleUserConfig.AuthCodeURL("state")
	} else if in.OauthType == pb.OauthType_OAUTH_FACEBOOK{
		link = s.facebookUserConfig.AuthCodeURL("state")
	} else{
		return nil, pb.InvalidArgument.ErrorMsg("invalid authorization server type")
	}
	return &pb.UserOauthLoginResponse{
		FollowLink: link,
	}, nil
}

func (s *RegistrationService) UserOauthLoginVerify(ctx context.Context, in *pb.UserOauthLoginVerifyRequest) (*pb.UserOauthLoginVerifyResponse, error) {
	// Start Redis Operations
	key := "oauth_" + in.DeviceId
	redisMap,err := s.redis.HGetAll(ctx, key).Result()
	if err != nil {
		return nil, pb.InvalidCode.ErrorNoMsg()
	}
	oauthType,err := strconv.ParseInt(redisMap[REDIS_KEY_OAUTH_TYPE], 10, 32)
	userType,err := strconv.ParseInt(redisMap[REDIS_KEY_USER_TYPE], 10, 32)

	if err != nil {
		return nil, pb.InvalidCode.ErrorNoMsg()
	}
	// Finish Redis Operations

	if int(userType) != USER_TYPE_USER {
		return nil, pb.InvalidArgument.ErrorMsg("user is not of type user")
	}

	var facebookId, googleId string

	var claims []security.Claim

	var userOauthInfo *pb.UserOauthInfo

	var exchangeCodeErr error

	if int(oauthType) == int(pb.OauthType_OAUTH_GOOGLE){
		code, _ := url.QueryUnescape(in.Code)

		tok, err := s.googleUserConfig.Exchange(ctx, code)

		if err!=nil{
			logger.Errorf("Invalid exchange code", err)
			exchangeCodeErr = pb.InvalidCode.ErrorNoMsg()
		} else {
			idToken := tok.Extra("id_token").(string)

			user, err := s.jwtUtils.ValidateUnsigned(idToken, false)

			if err != nil {
				logger.Errorf("google token validation failed ", err)
				exchangeCodeErr = pb.InvalidCode.ErrorMsg(err.Error())
			} else {
				googleId = user.Id
				userOauthInfo = &pb.UserOauthInfo{
					Name:  user.Name,
					Email: user.Email,
				}
			}
		}
	} else 	if int(oauthType) == int(pb.OauthType_OAUTH_FACEBOOK) {
		code, _ := url.QueryUnescape(in.Code)

		tok, err := s.facebookUserConfig.Exchange(ctx, code)

		if err!=nil{
			logger.Errorf("Invalid exchange code", err)
			exchangeCodeErr = pb.InvalidCode.ErrorNoMsg()
		} else {
			resp, err := http.Get("https://graph.facebook.com/me?fields=email,name,gender&access_token=" + url.QueryEscape(tok.AccessToken))
			defer resp.Body.Close()
			if err != nil {
				logger.Errorf("Error in fetching info from facebook", err)
				exchangeCodeErr = pb.Internal.ErrorMsg("cannot fetch user info from Facebook")
			} else {
				response, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					logger.Errorf("Error in fetching info from facebook", err)
					exchangeCodeErr = pb.Internal.ErrorMsg("cannot fetch user info from Facebook")
				} else {
					logger.Debugf("fetched user info : %v", string(response))

					var userInfo facebookUserInfo

					err = json.Unmarshal(response, &userInfo)

					if err != nil {
						logger.Errorf("Error in decoding message from facebook", err)
						exchangeCodeErr = pb.Internal.ErrorMsg("cannot decode user info from Facebook")
					} else {
						facebookId = userInfo.Id
						userOauthInfo = &pb.UserOauthInfo{
							Name:  userInfo.Name,
							Email: userInfo.Email,
						}
					}
				}
			}
		}
	} else {
		return nil, pb.InvalidArgument.ErrorMsg("invalid authorization server type")
	}

	if exchangeCodeErr != nil {
		oauthResponse, err := s.getOauthResponseFromCache(ctx, in.GetDeviceId())
		if err != nil{
			return nil, err
		}
		if oauthResponse == nil {
			return nil, exchangeCodeErr
		}
		if int(oauthType) == int(pb.OauthType_OAUTH_GOOGLE){
			googleId = oauthResponse.Identifier
		} else if int(oauthType) == int(pb.OauthType_OAUTH_FACEBOOK){
			facebookId = oauthResponse.Identifier
		}
		userOauthInfo = oauthResponse.Info
	}

	if int(oauthType) == int(pb.OauthType_OAUTH_GOOGLE) {
		claims = []security.Claim{{
			ClaimType:  security.CLAIM_TYPE_GOOGLE_ID,
			Identifier: googleId,
		}}
	} else if int(oauthType) == int(pb.OauthType_OAUTH_FACEBOOK) {
		claims = []security.Claim{{
			ClaimType:  security.CLAIM_TYPE_FACEBOOK_ID,
			Identifier: facebookId,
		}}
	}

	var user *security.User

	var selectedClaim security.Claim

	for _,claim := range claims{
		if claim.ClaimType == security.CLAIM_TYPE_GOOGLE_ID  ||
			claim.ClaimType == security.CLAIM_TYPE_FACEBOOK_ID  {
			selectedClaim = claim
		}
	}

	if selectedClaim.ClaimType == security.CLAIM_TYPE_GOOGLE_ID  {
		user,_ = s.partyAPI.FindUserByGoogleId(selectedClaim.Identifier, proto.UserType_USER_TYPE_PASSENGER)
	} else if selectedClaim.ClaimType == security.CLAIM_TYPE_FACEBOOK_ID  {
		user,_ = s.partyAPI.FindUserByFacebookId(selectedClaim.Identifier, proto.UserType_USER_TYPE_PASSENGER)
	}

	if user == nil || user.Id == ""{
		if selectedClaim.ClaimType == security.CLAIM_TYPE_GOOGLE_ID {
			if err := s.insertOauthResponseIntoCache(ctx, oauthResponse{
				DeviceId:   in.GetDeviceId(),
				OauthType:  int64(pb.OauthType_OAUTH_GOOGLE),
				UserType:   USER_TYPE_USER,
				Identifier: googleId,
				Info:       userOauthInfo,
			}); err != nil {
				return nil, err
			}
		} else if selectedClaim.ClaimType == security.CLAIM_TYPE_FACEBOOK_ID {
			if err := s.insertOauthResponseIntoCache(ctx, oauthResponse{
				DeviceId:   in.GetDeviceId(),
				OauthType:  int64(pb.OauthType_OAUTH_FACEBOOK),
				UserType:   USER_TYPE_USER,
				Identifier: facebookId,
				Info:       userOauthInfo,
			}); err != nil {
				return nil, err
			}
		}
		return nil, pb.NotFound.ErrorMsg("user not found")
	}

	user.Roles = []security.Role{security.Role_CLIENT}

	for _,claim := range claims{
		var has = false
		for _,claim2 := range user.Claims{
			if claim.ClaimType == claim2.ClaimType && claim.Identifier == claim2.Identifier {
				has = true
				break
			}
		}
		if !has {
			user.Claims = append(user.Claims, claim)
		}
	}

	token, err := s.jwtUtils.GenerateToken(*user)

	if err != nil{
		logger.Errorf("error in generating token", err)
		return nil, pb.Internal.ErrorMsg("error in generating token")
	}

	return &pb.UserOauthLoginVerifyResponse{
		Token: token,
		Info: userOauthInfo,
	}, nil
}

func (s *RegistrationService) insertOauthResponseIntoCache(ctx context.Context, oauthResponse oauthResponse) error {
	key := "oauth_res_" + oauthResponse.DeviceId
	fields := map[string]string{}
	b, err := json.Marshal(oauthResponse)
	if err != nil {
		return pb.Internal.Error(err)
	}
	fields[REDIS_KEY_OAUTH_RESPONSE] = string(b)
	if _,err := s.redis.HMSet(ctx, key, fields).Result(); err != nil{
		return pb.Internal.Error(err)
	}
	if _,err := s.redis.Expire(ctx, key, time.Duration(s.config.Redis.OTPSessionExpireSeconds) * time.Second).Result(); err != nil{
		return pb.Internal.Error(err)
	}
	return nil
}

func (s *RegistrationService) getOauthResponseFromCache(ctx context.Context, deviceId string) (*oauthResponse, error) {
	key := "oauth_res_" + deviceId
	redisMap,err := s.redis.HGetAll(ctx, key).Result()
	if err != nil {
		return nil, pb.InvalidCode.ErrorNoMsg()
	}
	jsonString := redisMap[REDIS_KEY_OAUTH_RESPONSE]
	var oauthResponse oauthResponse
	err = json.Unmarshal([]byte(jsonString), &oauthResponse)
	if err != nil {
		return nil, pb.Internal.Error(err)
	}
	return &oauthResponse, nil
}

type oauthResponse struct {
	DeviceId string
	OauthType int64
	UserType int64
	Identifier string
	Info *pb.UserOauthInfo
}