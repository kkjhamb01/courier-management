package api

import (
	"context"

	pb "github.com/kkjhamb01/courier-management/party/proto"
)

func (s *grpcServer) CreateCourierAccount(ctx context.Context, in *pb.CreateCourierAccountRequest) (*pb.CreateCourierAccountResponse, error) {
	return s.service.CreateCourierAccount(ctx, in)
}

func (s *grpcServer) UpdateCourierAccount(ctx context.Context, in *pb.UpdateCourierAccountRequest) (*pb.UpdateCourierAccountResponse, error) {
	return s.service.UpdateCourierAccount(ctx, in)
}

func (s *grpcServer) UpdateProfileAdditionalInfo(ctx context.Context, in *pb.UpdateProfileAdditionalInfoRequest) (*pb.UpdateProfileAdditionalInfoResponse, error) {
	return s.service.UpdateProfileAdditionalInfo(ctx, in)
}

func (s *grpcServer) GetProfileAdditionalInfo(ctx context.Context, in *pb.GetProfileAdditionalInfoRequest) (*pb.GetProfileAdditionalInfoResponse, error) {
	return s.service.GetProfileAdditionalInfo(ctx, in)
}

func (s *grpcServer) FindAccount(ctx context.Context, in *pb.FindAccountRequest) (*pb.FindAccountResponse, error) {
	return s.service.FindAccount(ctx, in)
}

func (s *grpcServer) RegisterClaim(ctx context.Context, in *pb.RegisterClaimRequest) (*pb.RegisterClaimResponse, error) {
	return s.service.RegisterClaim(ctx, in)
}

func (s *grpcServer) GetCourierAccount(ctx context.Context, in *pb.GetCourierAccountRequest) (*pb.GetCourierAccountResponse, error) {
	return s.service.GetCourierAccount(ctx, in)
}

func (s *grpcServer) FindCourierAccounts(ctx context.Context, in *pb.FindCourierAccountsRequest) (*pb.FindCourierAccountsResponse, error) {
	return s.service.FindCourierAccounts(ctx, in)
}

func (s *grpcServer) DeleteProfileAdditionalInfo(ctx context.Context, in *pb.DeleteProfileAdditionalInfoRequest) (*pb.DeleteProfileAdditionalInfoResponse, error) {
	return s.service.DeleteProfileAdditionalInfo(ctx, in)
}

func (s *grpcServer) GetProfileAdditionalInfoStatus(ctx context.Context, in *pb.GetProfileAdditionalInfoStatusRequest) (*pb.GetProfileAdditionalInfoStatusResponse, error) {
	return s.service.GetProfileAdditionalInfoStatus(ctx, in)
}

func (s *grpcServer) UpdateProfileAdditionalInfoStatus(ctx context.Context, in *pb.UpdateProfileAdditionalInfoStatusRequest) (*pb.UpdateProfileAdditionalInfoStatusResponse, error) {
	return s.service.UpdateProfileAdditionalInfoStatus(ctx, in)
}

func (s *grpcServer) GetProfileStatus(ctx context.Context, in *pb.GetProfileStatusRequest) (*pb.GetProfileStatusResponse, error) {
	return s.service.GetProfileStatus(ctx, in)
}

func (s *grpcServer) CreateUserAccount(ctx context.Context, in *pb.CreateUserAccountRequest) (*pb.CreateUserAccountResponse, error) {
	return s.service.CreateUserAccount(ctx, in)
}

func (s *grpcServer) GetUserAccount(ctx context.Context, in *pb.GetUserAccountRequest) (*pb.GetUserAccountResponse, error) {
	return s.service.GetUserAccount(ctx, in)
}

func (s *grpcServer) FindUserAccounts(ctx context.Context, in *pb.FindUserAccountsRequest) (*pb.FindUserAccountsResponse, error) {
	return s.service.FindUserAccounts(ctx, in)
}

func (s *grpcServer) UpdateUserAccount(ctx context.Context, in *pb.UpdateUserAccountRequest) (*pb.UpdateUserAccountResponse, error) {
	return s.service.UpdateUserAccount(ctx, in)
}

func (s *grpcServer) UpdateUserCard(ctx context.Context, in *pb.UpdateUserCardRequest) (*pb.UpdateUserCardResponse, error) {
	return s.service.UpdateUserCard(ctx, in)
}

func (s *grpcServer) DeleteUserCard(ctx context.Context, in *pb.DeleteUserCardRequest) (*pb.DeleteUserCardResponse, error) {
	return s.service.DeleteUserCard(ctx, in)
}

func (s *grpcServer) GetUserCard(ctx context.Context, in *pb.GetUserCardRequest) (*pb.GetUserCardResponse, error) {
	return s.service.GetUserCard(ctx, in)
}

func (s *grpcServer) SearchMot(ctx context.Context, in *pb.SearchMotRequest) (*pb.SearchMotResponse, error) {
	return s.service.SearchMot(ctx, in)
}

func (s *grpcServer) UpdateCourierPhoneNumber(ctx context.Context, in *pb.UpdateCourierPhoneNumberRequest) (*pb.UpdateCourierPhoneNumberResponse, error) {
	return s.service.UpdateCourierPhoneNumber(ctx, in)
}

func (s *grpcServer) UpdateUserPhoneNumber(ctx context.Context, in *pb.UpdateUserPhoneNumberRequest) (*pb.UpdateUserPhoneNumberResponse, error) {
	return s.service.UpdateUserPhoneNumber(ctx, in)
}

func (s *grpcServer) GetUserAddress(ctx context.Context, in *pb.GetUserAddressRequest) (*pb.GetUserAddressResponse, error) {
	return s.service.GetUserAddress(ctx, in)
}

func (s *grpcServer) UpdateUserAddress(ctx context.Context, in *pb.UpdateUserAddressRequest) (*pb.UpdateUserAddressResponse, error) {
	return s.service.UpdateUserAddress(ctx, in)
}
