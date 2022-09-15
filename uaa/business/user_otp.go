package business

import (
	"context"
	"encoding/json"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	"gitlab.artin.ai/backend/courier-management/party/proto"
	pb "gitlab.artin.ai/backend/courier-management/uaa/proto"
	"gitlab.artin.ai/backend/courier-management/uaa/security"
	"strconv"
	"time"
)

func (s *RegistrationService) UserOtpRegister(ctx context.Context, in *pb.UserOtpRegisterRequest) (*pb.UserOtpRegisterResponse, error){

	var user security.User

	ctxUser := ctx.Value("user")
	if ctxUser != nil {
		user = ctxUser.(security.User)
	}

	otp := s.generateOtp()

	remainingRetryCount := s.config.Otp.MaxRetry

	// Start Redis Operations
	key := "otp_" + in.DeviceId
	fields,_ := s.redis.HGetAll(ctx, key).Result()
	if fields == nil{
		fields = map[string]string{}
	}

	logger.Debugf("Registration request from %v, remainingRetryCount %v", in, remainingRetryCount)

	existingUser,err := s.partyAPI.FindUserByPhoneNumber(in.PhoneNumber, proto.UserType_USER_TYPE_PASSENGER)
	if err != nil{
		logger.Errorf("cannot perform query to find user from party", err)
		return nil, pb.Internal.ErrorMsg("cannot perform query to find user from party")
	}
	if existingUser != nil && existingUser.PhoneNumber != ""{
		return nil, pb.AlreadyExists.ErrorMsg("user already exists")
	}

	fields[REDIS_KEY_REMAINING_RETRY] = strconv.Itoa(int(remainingRetryCount))
	fields[REDIS_KEY_PHONE_NUMBER] = in.PhoneNumber
	fields[REDIS_KEY_OTP] = otp
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

	err = s.sendOtp(in.PhoneNumber, otp)
	logger.Debugf("generated otp ", otp)
	if err != nil{
		logger.Errorf("cannot send otp ", err)
		return nil, pb.Internal.Error(err)
	}

	return &pb.UserOtpRegisterResponse{
	}, nil
}

func (s *RegistrationService) UserOtpLogin(ctx context.Context, in *pb.UserOtpLoginRequest) (*pb.UserOtpLoginResponse, error) {

	otp := s.generateOtp()

	remainingRetryCount := s.config.Otp.MaxRetry

	// Start Redis Operations
	key := "otp_" + in.DeviceId
	fields,_ := s.redis.HGetAll(ctx, key).Result()
	if fields == nil{
		fields = map[string]string{}
	}

	logger.Debugf("Login request from %v, remainingRetryCount %v", in, remainingRetryCount)

	existingUser,err := s.partyAPI.FindUserByPhoneNumber(in.PhoneNumber, proto.UserType_USER_TYPE_PASSENGER)
	if err != nil{
		logger.Errorf("cannot perform query to find user from party", err)
		return nil, pb.Internal.ErrorMsg("cannot perform query to find user from party")
	}
	if existingUser == nil || existingUser.PhoneNumber == ""{
		return nil, pb.NotFound.ErrorMsg("user not found")
	}

	fields[REDIS_KEY_REMAINING_RETRY] = strconv.Itoa(int(remainingRetryCount))
	fields[REDIS_KEY_PHONE_NUMBER] = in.PhoneNumber
	fields[REDIS_KEY_OTP] = otp
	if _,err := s.redis.HMSet(ctx, key, fields).Result(); err != nil{
		return nil, pb.Internal.Error(err)
	}
	if _,err := s.redis.Expire(ctx, key, time.Duration(s.config.Redis.OTPSessionExpireSeconds) * time.Second).Result(); err != nil{
		return nil, pb.Internal.Error(err)
	}
	// Finish Redis Operations

	err = s.sendOtp(in.PhoneNumber, otp)
	logger.Debugf("generated otp ", otp)
	if err != nil{
		logger.Errorf("cannot send otp ", err)
		return nil, pb.Internal.Error(err)
	}

	return &pb.UserOtpLoginResponse{
	}, nil
}

func (s *RegistrationService) UserOtpAuthenticate(ctx context.Context, in *pb.UserOtpAuthenticateRequest) (*pb.UserOtpAuthenticateResponse, error){
	// Start Redis Operations
	key := "otp_" + in.DeviceId
	fields,err := s.redis.HGetAll(ctx, key).Result()
	if err != nil {
		logger.Error("cannot get redis", err)
		return nil, pb.InvalidPincode.ErrorNoMsg()
	}
	if len(fields) == 0{
		return nil, pb.NotFound.ErrorNoMsg()
	}
	otp := fields[REDIS_KEY_OTP]
	retry,_ := strconv.Atoi(fields[REDIS_KEY_REMAINING_RETRY])
	remainingRetryCount := int32(retry) - 1
	// Finish Redis Operations

	logger.Debugf("User authenticate request remainingRetryCount %v", remainingRetryCount)

	if remainingRetryCount < 0{
		return nil, pb.ExceedsMaximumRetry.ErrorNoMsg()
	}

	fields[REDIS_KEY_REMAINING_RETRY] = strconv.Itoa(int(remainingRetryCount))

	if _,err := s.redis.HMSet(ctx, key, fields).Result(); err != nil{
		return nil, pb.Internal.Error(err)
	}

	storedOtp := ""
	for _, cred := range in.Credentials{
		if cred.Type == pb.CredentialType_CREDENTIAL_TYPE_OTP{
			storedOtp = cred.Value
		}
	}
	if otp != storedOtp{
		return nil, pb.InvalidPincode.ErrorNoMsg()
	}

	phoneNumber := fields[REDIS_KEY_PHONE_NUMBER]
	if phoneNumber == "" {
		return nil, pb.InvalidPincode.ErrorNoMsg()
	}

	existingUser,err := s.partyAPI.FindUserByPhoneNumber(phoneNumber, proto.UserType_USER_TYPE_PASSENGER)
	if err != nil{
		logger.Errorf("cannot perform query to find user from party", err)
	}

	var name string
	var userId string
	var claims []security.Claim
	if existingUser != nil{
		userId = existingUser.Id
		name = existingUser.Name
		claims = existingUser.Claims
	} else {
		claims = []security.Claim{}
		if val, ok := fields[REDIS_KEY_USER]; ok {
			var tokenUser security.User
			json.Unmarshal([]byte(val), &tokenUser)
			if tokenUser.Id != "" {
				userId = tokenUser.Id
				for _,claim := range tokenUser.Claims {
					claims = append(claims, claim)
				}
			}
		} else {
			userId = generateId(false)
		}
	}

	var authorized = false
	for _,claim := range claims{
		if claim.ClaimType == security.CLAIM_TYPE_PHONE_NUMBER && claim.Identifier == phoneNumber{
			authorized = true
			break
		}
	}

	if !authorized {
		claims = append(claims, security.Claim{
			ClaimType:  security.CLAIM_TYPE_PHONE_NUMBER,
			Identifier: phoneNumber,
		})
	}

	if existingUser != nil{
		phoneNumber = existingUser.PhoneNumber
	}

	user := security.User{
		Id: userId,
		DeviceID: in.DeviceId,
		PhoneNumber: phoneNumber,
		Roles: []security.Role{security.Role_CLIENT},
		Name: name,
		Claims: claims,
	}

	user.Roles = []security.Role{security.Role_CLIENT}

	token, err := s.jwtUtils.GenerateToken(user)

	return &pb.UserOtpAuthenticateResponse{
		Token: token,
	}, err
}

func (s *RegistrationService) UserOtpRetry(ctx context.Context, in *pb.UserOtpRetryRequest) (*pb.UserOtpRetryResponse, error){
	otp := s.generateOtp()

	remainingRetryCount := s.config.Otp.MaxRetry

	// Start Redis Operations
	key := "otp_" + in.DeviceId
	fields,_ := s.redis.HGetAll(ctx, key).Result()
	if fields == nil{
		fields = map[string]string{}
	} else if len(fields) > 0 {
		retry,_ := strconv.Atoi(fields[REDIS_KEY_REMAINING_RETRY])
		remainingRetryCount = int32(retry) - 1
	}

	logger.Debugf("Retry request from %v, remainingRetryCount %v", in, remainingRetryCount)

	if remainingRetryCount < 0{
		return nil, pb.ExceedsMaximumRetry.ErrorNoMsg()
	}

	fields[REDIS_KEY_REMAINING_RETRY] = strconv.Itoa(int(remainingRetryCount))
	fields[REDIS_KEY_PHONE_NUMBER] = in.PhoneNumber
	fields[REDIS_KEY_OTP] = otp
	if _,err := s.redis.HMSet(ctx, key, fields).Result(); err != nil{
		return nil, err
	}
	if _,err := s.redis.Expire(ctx, key, time.Duration(s.config.Redis.OTPSessionExpireSeconds) * time.Second).Result(); err != nil{
		return nil, pb.Internal.Error(err)
	}
	// Finish Redis Operations

	err := s.sendOtp(in.PhoneNumber, otp)
	logger.Debugf("generated otp ", otp)
	if err != nil{
		logger.Errorf("cannot send otp ", err)
		return nil, pb.Internal.Error(err)
	}

	return &pb.UserOtpRetryResponse{
	}, nil
}

func (s *RegistrationService) UserOtpReclaim(ctx context.Context, in *pb.UserOtpReclaimRequest) (*pb.UserOtpReclaimResponse, error){

	otp := s.generateOtp()

	remainingRetryCount := s.config.Otp.MaxRetry

	// Start Redis Operations
	key := "otp_" + in.DeviceId
	fields,_ := s.redis.HGetAll(ctx, key).Result()
	if fields == nil{
		fields = map[string]string{}
	}

	logger.Debugf("Reclaim request from %v, remainingRetryCount %v", in, remainingRetryCount)

	existingUser,err := s.partyAPI.FindUserByPhoneNumber(in.PhoneNumber, proto.UserType_USER_TYPE_PASSENGER)
	if err != nil{
		logger.Errorf("cannot perform query to find user from party", err)
		return nil, pb.Internal.ErrorMsg("cannot perform query to find user from party")
	}
	if existingUser != nil && existingUser.PhoneNumber != ""{
		return nil, pb.AlreadyExists.ErrorMsg("user already exists")
	}

	fields[REDIS_KEY_REMAINING_RETRY] = strconv.Itoa(int(remainingRetryCount))
	fields[REDIS_KEY_PHONE_NUMBER] = in.PhoneNumber
	fields[REDIS_KEY_OTP] = otp
	if _,err := s.redis.HMSet(ctx, key, fields).Result(); err != nil{
		return nil, pb.Internal.Error(err)
	}
	if _,err := s.redis.Expire(ctx, key, time.Duration(s.config.Redis.OTPSessionExpireSeconds) * time.Second).Result(); err != nil{
		return nil, pb.Internal.Error(err)
	}
	// Finish Redis Operations

	err = s.sendOtp(in.PhoneNumber, otp)
	logger.Debugf("generated otp ", otp)
	if err != nil{
		logger.Errorf("cannot send otp ", err)
		return nil, pb.Internal.Error(err)
	}

	return &pb.UserOtpReclaimResponse{
	}, nil
}