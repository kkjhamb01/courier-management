package api

import (
	"context"
	pb "gitlab.artin.ai/backend/courier-management/uaa/proto"
)

func (s *grpcServer) CourierOtpRegister(ctx context.Context, in *pb.CourierOtpRegisterRequest) (*pb.CourierOtpRegisterResponse, error){
	return s.registrationService.CourierOtpRegister(ctx, in)
}

func (s *grpcServer) CourierOtpLogin(ctx context.Context, in *pb.CourierOtpLoginRequest) (*pb.CourierOtpLoginResponse, error){
	return s.registrationService.CourierOtpLogin(ctx, in)
}

func (s *grpcServer) CourierOtpAuthenticate(ctx context.Context, in *pb.CourierOtpAuthenticateRequest) (*pb.CourierOtpAuthenticateResponse, error){
	return s.registrationService.CourierOtpAuthenticate(ctx, in)
}

func (s *grpcServer) CourierOtpRetry(ctx context.Context, in *pb.CourierOtpRetryRequest) (*pb.CourierOtpRetryResponse, error){
	return s.registrationService.CourierOtpRetry(ctx, in)
}

func (s *grpcServer) CourierOauthRegister(ctx context.Context, in *pb.CourierOauthRegisterRequest) (*pb.CourierOauthRegisterResponse, error) {
	return s.registrationService.CourierOauthRegister(ctx, in)
}

func (s *grpcServer) CourierOauthRegisterVerify(ctx context.Context, in *pb.CourierOauthRegisterVerifyRequest) (*pb.CourierOauthRegisterVerifyResponse, error) {
	return s.registrationService.CourierOauthRegisterVerify(ctx, in)
}

func (s *grpcServer) CourierOauthLogin(ctx context.Context, in *pb.CourierOauthLoginRequest) (*pb.CourierOauthLoginResponse, error) {
	return s.registrationService.CourierOauthLogin(ctx, in)
}

func (s *grpcServer) CourierOauthLoginVerify(ctx context.Context, in *pb.CourierOauthLoginVerifyRequest) (*pb.CourierOauthLoginVerifyResponse, error) {
	return s.registrationService.CourierOauthLoginVerify(ctx, in)
}

func (s *grpcServer) CourierOtpReclaim(ctx context.Context, in *pb.CourierOtpReclaimRequest) (*pb.CourierOtpReclaimResponse, error){
	return s.registrationService.CourierOtpReclaim(ctx, in)
}


func (s *grpcServer) UserOtpRegister(ctx context.Context, in *pb.UserOtpRegisterRequest) (*pb.UserOtpRegisterResponse, error){
	return s.registrationService.UserOtpRegister(ctx, in)
}

func (s *grpcServer) UserOtpLogin(ctx context.Context, in *pb.UserOtpLoginRequest) (*pb.UserOtpLoginResponse, error){
	return s.registrationService.UserOtpLogin(ctx, in)
}

func (s *grpcServer) UserOtpAuthenticate(ctx context.Context, in *pb.UserOtpAuthenticateRequest) (*pb.UserOtpAuthenticateResponse, error){
	return s.registrationService.UserOtpAuthenticate(ctx, in)
}

func (s *grpcServer) UserOtpRetry(ctx context.Context, in *pb.UserOtpRetryRequest) (*pb.UserOtpRetryResponse, error){
	return s.registrationService.UserOtpRetry(ctx, in)
}

func (s *grpcServer) UserOauthRegister(ctx context.Context, in *pb.UserOauthRegisterRequest) (*pb.UserOauthRegisterResponse, error) {
	return s.registrationService.UserOauthRegister(ctx, in)
}

func (s *grpcServer) UserOauthRegisterVerify(ctx context.Context, in *pb.UserOauthRegisterVerifyRequest) (*pb.UserOauthRegisterVerifyResponse, error) {
	return s.registrationService.UserOauthRegisterVerify(ctx, in)
}

func (s *grpcServer) UserOauthLogin(ctx context.Context, in *pb.UserOauthLoginRequest) (*pb.UserOauthLoginResponse, error) {
	return s.registrationService.UserOauthLogin(ctx, in)
}

func (s *grpcServer) UserOauthLoginVerify(ctx context.Context, in *pb.UserOauthLoginVerifyRequest) (*pb.UserOauthLoginVerifyResponse, error) {
	return s.registrationService.UserOauthLoginVerify(ctx, in)
}

func (s *grpcServer) UserOtpReclaim(ctx context.Context, in *pb.UserOtpReclaimRequest) (*pb.UserOtpReclaimResponse, error){
	return s.registrationService.UserOtpReclaim(ctx, in)
}