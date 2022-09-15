package api

import (
	"context"
	pb "gitlab.artin.ai/backend/courier-management/party/proto"
)

func (s *grpcServer) InterServiceGetProfileAdditionalInfo(ctx context.Context, in *pb.InterServiceGetProfileAdditionalInfoRequest) (*pb.GetProfileAdditionalInfoResponse, error) {
	return s.service.InterServiceGetProfileAdditionalInfo(ctx, in)
}

func (s *grpcServer) InterServiceUpdateProfileAdditionalInfo(ctx context.Context, in *pb.InterServiceUpdateProfileAdditionalInfoRequest) (*pb.UpdateProfileAdditionalInfoResponse, error) {
	return s.service.InterServiceUpdateProfileAdditionalInfo(ctx, in)
}

func (s *grpcServer) InterServiceFindCourierAccounts(ctx context.Context, in *pb.InterServiceFindCourierAccountsRequest) (*pb.FindCourierAccountsResponse, error) {
	return s.service.InterServiceFindCourierAccounts(ctx, in)
}

func (s *grpcServer) InterServiceGetDocumentsOfUser(ctx context.Context, in *pb.InterServiceGetDocumentsOfUserRequest) (*pb.GetDocumentsOfUserResponse, error) {
	return s.service.InterServiceGetDocumentsOfUser(ctx, in)
}
