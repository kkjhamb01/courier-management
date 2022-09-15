package api

import (
	"context"

	pb "github.com/kkjhamb01/courier-management/party/proto"
)

func (s *grpcServer) OpenGetCourierAccount(ctx context.Context, in *pb.OpenGetCourierAccountRequest) (*pb.GetCourierAccountResponse, error) {
	return s.service.OpenGetCourierAccount(ctx, in)
}

func (s *grpcServer) OpenGetProfileAdditionalInfo(ctx context.Context, in *pb.OpenGetProfileAdditionalInfoRequest) (*pb.GetProfileAdditionalInfoResponse, error) {
	return s.service.OpenGetProfileAdditionalInfo(ctx, in)
}

func (s *grpcServer) OpenGetDocumentsOfUser(ctx context.Context, in *pb.OpenGetDocumentsOfUserRequest) (*pb.GetDocumentsOfUserResponse, error) {
	return s.service.OpenGetDocumentsOfUser(ctx, in)
}

func (s *grpcServer) OpenGetUserAccount(ctx context.Context, in *pb.OpenGetUserAccountRequest) (*pb.GetUserAccountResponse, error) {
	return s.service.OpenGetUserAccount(ctx, in)
}

func (s *grpcServer) OpenGetCourierPublicInfo(ctx context.Context, in *pb.OpenGetCourierPublicInfoRequest) (*pb.OpenGetCourierPublicInfoResponse, error) {
	return s.service.OpenGetCourierPublicInfo(ctx, in)
}
