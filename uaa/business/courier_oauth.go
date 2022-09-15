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

func (s *RegistrationService) CourierOauthRegister(ctx context.Context, in *pb.CourierOauthRegisterRequest) (*pb.CourierOauthRegisterResponse, error) {

	user := ctx.Value("user").(security.User)

	logger.Debugf("CourierOauthRegister oauth_type = %v, device_id = %v, userId = ",
		in.GetOauthType(), in.GetDeviceId(), user.Id)

	if user.Id == "" {
		logger.Debugf("CourierOauthRegister invalid token")
		return nil, pb.Unauthenticated.ErrorMsg("invalid token")
	}

	var authorizedPhoneNumber = false
	for _,claim := range user.Claims {
		if claim.ClaimType == security.CLAIM_TYPE_PHONE_NUMBER{
			authorizedPhoneNumber = true
		}
	}

	if !authorizedPhoneNumber {
		logger.Debugf("CourierOauthRegister token doesn't have any authorized phone number")
		return nil, pb.Unauthenticated.ErrorMsg("token doesn't have any authorized phone number")
	}

	// Start Redis Operations
	key := "oauth_" + in.DeviceId
	fields := map[string]string{}
	fields[REDIS_KEY_OAUTH_TYPE] = strconv.FormatInt(int64(in.OauthType), 10)
	fields[REDIS_KEY_USER_TYPE] = strconv.FormatInt(USER_TYPE_COURIER, 10)
	fields[REDIS_KEY_PHONE_NUMBER] = user.PhoneNumber
	fields[REDIS_KEY_USER_ID] = user.Id
	if _,err := s.redis.HMSet(ctx, key, fields).Result(); err != nil{
		return nil, pb.Internal.Error(err)
	}
	if _,err := s.redis.Expire(ctx, key, time.Duration(s.config.Redis.OTPSessionExpireSeconds) * time.Second).Result(); err != nil{
		return nil, pb.Internal.Error(err)
	}
	// Finish Redis Operations

	var link string
	if in.OauthType == pb.OauthType_OAUTH_GOOGLE{
		link = s.googleCourierConfig.AuthCodeURL("state")
	} else if in.OauthType == pb.OauthType_OAUTH_FACEBOOK{
		link = s.facebookCourierConfig.AuthCodeURL("state")
	} else{
		return nil, pb.InvalidArgument.ErrorMsg("invalid authorization server type")
	}
	return &pb.CourierOauthRegisterResponse{
		FollowLink: link,
	}, nil
}

func (s *RegistrationService) CourierOauthRegisterVerify(ctx context.Context, in *pb.CourierOauthRegisterVerifyRequest) (*pb.CourierOauthRegisterVerifyResponse, error) {

	logger.Debugf("CourierOauthRegisterVerify device_id = %v, code = %v",
		in.GetDeviceId(), in.GetCode())

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

	if int(userType) != USER_TYPE_COURIER {
		return nil, pb.InvalidArgument.ErrorMsg("user is not of type courier")
	}

	userId := redisMap[REDIS_KEY_USER_ID]

	var facebookId, googleId string

	var userOauthInfo *pb.UserOauthInfo

	if int(oauthType) == int(pb.OauthType_OAUTH_GOOGLE){
		code, _ := url.QueryUnescape(in.Code)

		tok, err := s.googleCourierConfig.Exchange(ctx, code)

		if err!=nil{
			logger.Errorf("Invalid exchange code", err)
			return nil, pb.InvalidCode.ErrorNoMsg()
		}

		idToken := tok.Extra("id_token").(string)

		user, err := s.jwtUtils.ValidateUnsigned(idToken, false)
		if err != nil{
			logger.Errorf("google token validation failed ", err)
			return nil, pb.InvalidCode.ErrorMsg(err.Error())
		}
		googleId = user.Id
		userOauthInfo = &pb.UserOauthInfo{
			Name: user.Name,
			Email: user.Email,
			FirstName: user.FirstName,
			LastName: user.LastName,
		}
	} else 	if int(oauthType) == int(pb.OauthType_OAUTH_FACEBOOK) {
		code, _ := url.QueryUnescape(in.Code)

		tok, err := s.facebookCourierConfig.Exchange(ctx, code)

		if err!=nil{
			logger.Errorf("Invalid exchange code", err)
			return nil, pb.InvalidCode.ErrorNoMsg()
		}

		resp, err := http.Get("https://graph.facebook.com/me?fields=email,name,first_name,last_name&access_token=" + url.QueryEscape(tok.AccessToken))
		if err != nil {
			logger.Errorf("Error in fetching info from facebook", err)
			return nil, pb.Internal.ErrorMsg("cannot fetch user info from Facebook")
		}
		defer resp.Body.Close()

		response, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logger.Errorf("Error in fetching info from facebook", err)
			return nil, pb.Internal.ErrorMsg("cannot fetch user info from Facebook")
		}

		logger.Debugf("fetched user info : %v", string(response))

		var userInfo facebookUserInfo

		err = json.Unmarshal(response, &userInfo)

		if err != nil {
			logger.Errorf("Error in decoding message from facebook", err)
			return nil, pb.Internal.ErrorMsg("cannot decode user info from Facebook")
		}

		facebookId = userInfo.Id
		userOauthInfo = &pb.UserOauthInfo{
			Name: userInfo.Name,
			Email: userInfo.Email,
			FirstName: userInfo.FirstName,
			LastName: userInfo.LastName,
		}

	} else {
		return nil, pb.InvalidArgument.ErrorMsg("invalid authorization server type")
	}

	if googleId != "" {
		u,_ := s.partyAPI.FindUserByGoogleId(googleId, proto.UserType_USER_TYPE_CURIOUR)
		if u != nil && u.Id != ""{
			return nil, pb.AlreadyExists.ErrorMsg("already registered with this google id")
		}
	} else if facebookId != "" {
		u,_ := s.partyAPI.FindUserByFacebookId(facebookId, proto.UserType_USER_TYPE_CURIOUR)
		if u != nil && u.Id != ""{
			return nil, pb.AlreadyExists.ErrorMsg("already registered with this facebook id")
		}
	}

	var identifier string
	var claimType proto.ClaimType

	if facebookId != "" {
		identifier = facebookId
		claimType = proto.ClaimType_CLAIM_TYPE_FACEBOOK_ID
	} else if googleId != "" {
		identifier = googleId
		claimType = proto.ClaimType_CLAIM_TYPE_GOOGLE_ID
	}

	if err := s.partyAPI.RegisterClaim(userId, claimType, identifier, proto.UserType_USER_TYPE_CURIOUR); err != nil{
		logger.Errorf("cannot register claim to party api", err)
		return nil, pb.Internal.ErrorMsg("cannot register claim to party api " + err.Error())
	}

	return &pb.CourierOauthRegisterVerifyResponse{
		Info: userOauthInfo,
	}, nil
}

func (s *RegistrationService) CourierOauthLogin(ctx context.Context, in *pb.CourierOauthLoginRequest) (*pb.CourierOauthLoginResponse, error) {

	// Start Redis Operations
	key := "oauth_" + in.DeviceId
	fields := map[string]string{}
	fields[REDIS_KEY_OAUTH_TYPE] = strconv.FormatInt(int64(in.OauthType), 10)
	fields[REDIS_KEY_USER_TYPE] = strconv.FormatInt(USER_TYPE_COURIER, 10)
	if _,err := s.redis.HMSet(ctx, key, fields).Result(); err != nil{
		return nil, pb.Internal.Error(err)
	}
	if _,err := s.redis.Expire(ctx, key, time.Duration(s.config.Redis.OTPSessionExpireSeconds) * time.Second).Result(); err != nil{
		return nil, pb.Internal.Error(err)
	}
	// Finish Redis Operations

	var link string
	if in.OauthType == pb.OauthType_OAUTH_GOOGLE{
		link = s.googleCourierConfig.AuthCodeURL("state")
	} else if in.OauthType == pb.OauthType_OAUTH_FACEBOOK{
		link = s.facebookCourierConfig.AuthCodeURL("state")
	} else{
		return nil, pb.InvalidArgument.ErrorMsg("invalid authorization server type")
	}
	return &pb.CourierOauthLoginResponse{
		FollowLink: link,
	}, nil
}

func (s *RegistrationService) CourierOauthLoginVerify(ctx context.Context, in *pb.CourierOauthLoginVerifyRequest) (*pb.CourierOauthLoginVerifyResponse, error) {

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

	if int(userType) != USER_TYPE_COURIER {
		return nil, pb.InvalidArgument.ErrorMsg("user is not of type courier")
	}

	var facebookId, googleId string
	var claims []security.Claim

	var userOauthInfo *pb.UserOauthInfo

	if int(oauthType) == int(pb.OauthType_OAUTH_GOOGLE){
		code, _ := url.QueryUnescape(in.Code)

		tok, err := s.googleCourierConfig.Exchange(ctx, code)

		if err!=nil{
			logger.Errorf("Invalid exchange code", err)
			return nil, pb.InvalidCode.ErrorNoMsg()
		}

		idToken := tok.Extra("id_token").(string)

		user, err := s.jwtUtils.ValidateUnsigned(idToken, false)
		if err != nil{
			logger.Errorf("google token validation failed ", err)
			return nil, pb.InvalidCode.ErrorMsg(err.Error())
		}
		googleId = user.Id
		userOauthInfo = &pb.UserOauthInfo{
			Name: user.Name,
			Email: user.Email,
		}
		claims = []security.Claim{{
			ClaimType:  security.CLAIM_TYPE_GOOGLE_ID,
			Identifier: googleId,
		}}
	} else 	if int(oauthType) == int(pb.OauthType_OAUTH_FACEBOOK) {
		code, _ := url.QueryUnescape(in.Code)

		tok, err := s.facebookCourierConfig.Exchange(ctx, code)

		if err!=nil{
			logger.Errorf("Invalid exchange code", err)
			return nil, pb.InvalidCode.ErrorNoMsg()
		}

		resp, err := http.Get("https://graph.facebook.com/me?fields=email,name,gender&access_token=" + url.QueryEscape(tok.AccessToken))
		if err != nil {
			logger.Errorf("Error in fetching info from facebook", err)
			return nil, pb.Internal.ErrorMsg("cannot fetch user info from Facebook")
		}
		defer resp.Body.Close()

		response, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logger.Errorf("Error in fetching info from facebook", err)
			return nil, pb.Internal.ErrorMsg("cannot fetch user info from Facebook")
		}

		logger.Debugf("fetched user info : %v", string(response))

		var userInfo facebookUserInfo

		err = json.Unmarshal(response, &userInfo)

		if err != nil {
			logger.Errorf("Error in decoding message from facebook", err)
			return nil, pb.Internal.ErrorMsg("cannot decode user info from Facebook")
		}

		facebookId = userInfo.Id
		userOauthInfo = &pb.UserOauthInfo{
			Name: userInfo.Name,
			Email: userInfo.Email,
		}
		claims = []security.Claim{{
			ClaimType:  security.CLAIM_TYPE_FACEBOOK_ID,
			Identifier: facebookId,
		}}

	} else {
		return nil, pb.InvalidArgument.ErrorMsg("invalid authorization server type")
	}

	var user *security.User

	for _,claim := range claims{
		if claim.ClaimType == security.CLAIM_TYPE_GOOGLE_ID  {
			user,_ = s.partyAPI.FindUserByGoogleId(claim.Identifier, proto.UserType_USER_TYPE_CURIOUR)
		} else if claim.ClaimType == security.CLAIM_TYPE_FACEBOOK_ID  {
			user,_ = s.partyAPI.FindUserByFacebookId(claim.Identifier, proto.UserType_USER_TYPE_CURIOUR)
		}
	}

	if user == nil || user.Id == ""{
		return nil, pb.NotFound.ErrorMsg("user not found")
	}

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

	user.Roles = []security.Role{security.Role_COURIER}

	token, err := s.jwtUtils.GenerateToken(*user)

	if err != nil{
		logger.Errorf("error in generating token", err)
		return nil, pb.Internal.ErrorMsg("error in generating token")
	}

	return &pb.CourierOauthLoginVerifyResponse{
		Token: token,
		Info: userOauthInfo,
	}, nil
}
