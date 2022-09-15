package business

import (
	"context"
	"strconv"
	"time"

	"github.com/kkjhamb01/courier-management/common/logger"
	"github.com/kkjhamb01/courier-management/party/proto"
	pb "github.com/kkjhamb01/courier-management/uaa/proto"
	"github.com/kkjhamb01/courier-management/uaa/security"
)

func (s *RegistrationService) CourierOtpRegister(ctx context.Context, in *pb.CourierOtpRegisterRequest) (*pb.CourierOtpRegisterResponse, error) {

	otp := s.generateOtp()

	remainingRetryCount := s.config.Otp.MaxRetry

	// Start Redis Operations
	key := "otp_" + in.DeviceId
	fields, _ := s.redis.HGetAll(ctx, key).Result()
	if fields == nil {
		fields = map[string]string{}
	}

	logger.Debugf("Registration request from %v, remainingRetryCount %v", in, remainingRetryCount)

	existingUser, err := s.partyAPI.FindUserByPhoneNumber(in.PhoneNumber, proto.UserType_USER_TYPE_CURIOUR)
	if err != nil {
		logger.Errorf("cannot perform query to find user from party", err)
		return nil, pb.Internal.ErrorMsg("cannot perform query to find user from party")
	}
	if existingUser != nil && existingUser.PhoneNumber != "" {
		return nil, pb.AlreadyExists.ErrorMsg("user already exists")
	}

	fields[REDIS_KEY_REMAINING_RETRY] = strconv.Itoa(int(remainingRetryCount))
	fields[REDIS_KEY_PHONE_NUMBER] = in.PhoneNumber
	fields[REDIS_KEY_OTP] = otp
	if _, err := s.redis.HMSet(ctx, key, fields).Result(); err != nil {
		return nil, pb.Internal.Error(err)
	}
	if _, err := s.redis.Expire(ctx, key, time.Duration(s.config.Redis.OTPSessionExpireSeconds)*time.Second).Result(); err != nil {
		return nil, pb.Internal.Error(err)
	}
	// Finish Redis Operations

	err = s.sendOtp(in.PhoneNumber, otp)
	logger.Debugf("generated otp ", otp)
	if err != nil {
		logger.Errorf("cannot send otp ", err)
		return nil, pb.Internal.Error(err)
	}

	return &pb.CourierOtpRegisterResponse{}, nil
}

func (s *RegistrationService) CourierOtpLogin(ctx context.Context, in *pb.CourierOtpLoginRequest) (*pb.CourierOtpLoginResponse, error) {

	otp := s.generateOtp()

	remainingRetryCount := s.config.Otp.MaxRetry

	// Start Redis Operations
	key := "otp_" + in.DeviceId
	fields, _ := s.redis.HGetAll(ctx, key).Result()
	if fields == nil {
		fields = map[string]string{}
	}

	logger.Debugf("Login request from %v, remainingRetryCount %v", in, remainingRetryCount)

	existingUser, err := s.partyAPI.FindUserByPhoneNumber(in.PhoneNumber, proto.UserType_USER_TYPE_CURIOUR)
	if err != nil {
		logger.Errorf("cannot perform query to find user from party", err)
		return nil, pb.Internal.ErrorMsg("cannot perform query to find user from party")
	}
	if existingUser == nil || existingUser.PhoneNumber == "" {
		return nil, pb.NotFound.ErrorMsg("user not found")
	}

	fields[REDIS_KEY_REMAINING_RETRY] = strconv.Itoa(int(remainingRetryCount))
	fields[REDIS_KEY_PHONE_NUMBER] = in.PhoneNumber
	fields[REDIS_KEY_OTP] = otp
	if _, err := s.redis.HMSet(ctx, key, fields).Result(); err != nil {
		return nil, pb.Internal.Error(err)
	}
	if _, err := s.redis.Expire(ctx, key, time.Duration(s.config.Redis.OTPSessionExpireSeconds)*time.Second).Result(); err != nil {
		return nil, pb.Internal.Error(err)
	}
	// Finish Redis Operations

	err = s.sendOtp(in.PhoneNumber, otp)
	logger.Debugf("generated otp ", otp)
	if err != nil {
		logger.Errorf("cannot send otp ", err)
		return nil, pb.Internal.Error(err)
	}

	return &pb.CourierOtpLoginResponse{}, nil
}

func (s *RegistrationService) CourierOtpAuthenticate(ctx context.Context, in *pb.CourierOtpAuthenticateRequest) (*pb.CourierOtpAuthenticateResponse, error) {

	// Start Redis Operations
	key := "otp_" + in.DeviceId
	fields, err := s.redis.HGetAll(ctx, key).Result()
	if err != nil {
		logger.Error("cannot get redis", err)
		return nil, pb.InvalidPincode.ErrorNoMsg()
	}
	if len(fields) == 0 {
		return nil, pb.NotFound.ErrorNoMsg()
	}
	retry, _ := strconv.Atoi(fields[REDIS_KEY_REMAINING_RETRY])
	remainingRetryCount := int32(retry) - 1

	logger.Debugf("Courier authenticate request remainingRetryCount %v", remainingRetryCount)
	if remainingRetryCount < 0 {
		return nil, pb.ExceedsMaximumRetry.ErrorNoMsg()
	}

	fields[REDIS_KEY_REMAINING_RETRY] = strconv.Itoa(int(remainingRetryCount))

	if _, err := s.redis.HMSet(ctx, key, fields).Result(); err != nil {
		return nil, pb.Internal.Error(err)
	}

	otp := fields[REDIS_KEY_OTP]
	// Finish Redis Operations

	storedOtp := ""
	for _, cred := range in.Credentials {
		if cred.Type == pb.CredentialType_CREDENTIAL_TYPE_OTP {
			storedOtp = cred.Value
		}
	}
	if otp != storedOtp {
		return nil, pb.InvalidPincode.ErrorNoMsg()
	}

	phoneNumber := fields[REDIS_KEY_PHONE_NUMBER]
	if phoneNumber == "" {
		return nil, pb.InvalidPincode.ErrorNoMsg()
	}

	existingUser, err := s.partyAPI.FindUserByPhoneNumber(phoneNumber, proto.UserType_USER_TYPE_CURIOUR)
	if err != nil {
		logger.Errorf("cannot perform query to find user from party", err)
	}

	var name string
	var userId string
	var claims []security.Claim
	if existingUser != nil {
		userId = existingUser.Id
		name = existingUser.Name
		claims = existingUser.Claims
	} else {
		if oldUserId, ok := fields[REDIS_KEY_USER_ID]; ok {
			userId = oldUserId
		} else {
			userId = generateId(true)
		}
		claims = []security.Claim{}
	}

	var authorized = false
	for _, claim := range claims {
		if claim.ClaimType == security.CLAIM_TYPE_PHONE_NUMBER && claim.Identifier == phoneNumber {
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

	user := security.User{
		Id:          userId,
		DeviceID:    in.DeviceId,
		PhoneNumber: phoneNumber,
		Roles:       []security.Role{security.Role_COURIER},
		Name:        name,
		Claims:      claims,
	}

	user.Roles = []security.Role{security.Role_COURIER}

	token, err := s.jwtUtils.GenerateToken(user)

	return &pb.CourierOtpAuthenticateResponse{
		Token: token,
	}, err
}

func (s *RegistrationService) CourierOtpRetry(ctx context.Context, in *pb.CourierOtpRetryRequest) (*pb.CourierOtpRetryResponse, error) {
	otp := s.generateOtp()

	remainingRetryCount := s.config.Otp.MaxRetry

	// Start Redis Operations
	key := "otp_" + in.DeviceId
	fields, _ := s.redis.HGetAll(ctx, key).Result()
	if fields == nil {
		fields = map[string]string{}
	} else if len(fields) > 0 {
		retry, _ := strconv.Atoi(fields[REDIS_KEY_REMAINING_RETRY])
		remainingRetryCount = int32(retry) - 1
	}

	logger.Debugf("Retry request from %v, remainingRetryCount %v", in, remainingRetryCount)

	if remainingRetryCount < 0 {
		return nil, pb.ExceedsMaximumRetry.ErrorNoMsg()
	}

	fields[REDIS_KEY_REMAINING_RETRY] = strconv.Itoa(int(remainingRetryCount))
	fields[REDIS_KEY_PHONE_NUMBER] = in.PhoneNumber
	fields[REDIS_KEY_OTP] = otp
	if _, err := s.redis.HMSet(ctx, key, fields).Result(); err != nil {
		return nil, err
	}
	if _, err := s.redis.Expire(ctx, key, time.Duration(s.config.Redis.OTPSessionExpireSeconds)*time.Second).Result(); err != nil {
		return nil, pb.Internal.Error(err)
	}
	// Finish Redis Operations

	err := s.sendOtp(in.PhoneNumber, otp)
	logger.Debugf("generated otp ", otp)
	if err != nil {
		logger.Errorf("cannot send otp ", err)
		return nil, pb.Internal.Error(err)
	}

	return &pb.CourierOtpRetryResponse{}, nil
}

func (s *RegistrationService) CourierOtpReclaim(ctx context.Context, in *pb.CourierOtpReclaimRequest) (*pb.CourierOtpReclaimResponse, error) {

	otp := s.generateOtp()

	remainingRetryCount := s.config.Otp.MaxRetry

	// Start Redis Operations
	key := "otp_" + in.DeviceId
	fields, _ := s.redis.HGetAll(ctx, key).Result()
	if fields == nil {
		fields = map[string]string{}
	}

	logger.Debugf("Reclaim request from %v, remainingRetryCount %v", in, remainingRetryCount)

	existingUser, err := s.partyAPI.FindUserByPhoneNumber(in.PhoneNumber, proto.UserType_USER_TYPE_CURIOUR)
	if err != nil {
		logger.Errorf("cannot perform query to find user from party", err)
		return nil, pb.Internal.ErrorMsg("cannot perform query to find user from party")
	}
	if existingUser != nil && existingUser.PhoneNumber != "" {
		return nil, pb.AlreadyExists.ErrorMsg("user already exists")
	}

	oldUser, err := s.partyAPI.FindUserByPhoneNumber(in.OldPhoneNumber, proto.UserType_USER_TYPE_CURIOUR)
	if err != nil {
		logger.Errorf("cannot perform query to find user from party", err)
		return nil, pb.Internal.ErrorMsg("cannot perform query to find user from party")
	}
	if oldUser == nil || oldUser.PhoneNumber == "" {
		return nil, pb.NotFound.ErrorMsg("user does not exist")
	}

	fields[REDIS_KEY_REMAINING_RETRY] = strconv.Itoa(int(remainingRetryCount))
	fields[REDIS_KEY_PHONE_NUMBER] = in.PhoneNumber
	fields[REDIS_KEY_USER_ID] = oldUser.Id
	fields[REDIS_KEY_OTP] = otp
	if _, err := s.redis.HMSet(ctx, key, fields).Result(); err != nil {
		return nil, pb.Internal.Error(err)
	}
	if _, err := s.redis.Expire(ctx, key, time.Duration(s.config.Redis.OTPSessionExpireSeconds)*time.Second).Result(); err != nil {
		return nil, pb.Internal.Error(err)
	}
	// Finish Redis Operations

	err = s.sendOtp(in.PhoneNumber, otp)
	logger.Debugf("generated otp ", otp)
	if err != nil {
		logger.Errorf("cannot send otp ", err)
		return nil, pb.Internal.Error(err)
	}

	return &pb.CourierOtpReclaimResponse{}, nil
}
