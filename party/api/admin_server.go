package api

import (
	"context"
	pb "gitlab.artin.ai/backend/courier-management/party/proto"
)

func (s *grpcServer) ServiceGetCourierRegistrationStat(ctx context.Context, in *pb.ServiceGetCourierRegistrationStatRequest) (*pb.GetRegistrationStatResponse, error){
	return s.service.ServiceGetCourierRegistrationStat(ctx, in)
}

func (s *grpcServer) ServiceGetClientRegistrationStat(ctx context.Context, in *pb.ServiceGetClientRegistrationStatRequest) (*pb.GetRegistrationStatResponse, error){
	return s.service.ServiceGetClientRegistrationStat(ctx, in)
}

func (s *grpcServer) ServiceGetCourierRegistration(ctx context.Context, in *pb.ServiceGetCourierRegistrationRequest) (*pb.GetCourierRegistrationResponse, error){
	return s.service.ServiceGetCourierRegistration(ctx, in)
}

func (s *grpcServer) ServiceGetClientRegistration(ctx context.Context, in *pb.ServiceGetClientRegistrationRequest) (*pb.GetClientRegistrationResponse, error){
	return s.service.ServiceGetClientRegistration(ctx, in)
}

func (s *grpcServer) ServiceGetDocumentsOfUser(ctx context.Context, in *pb.ServiceGetDocumentsOfUserRequest) (*pb.GetDocumentsOfUserResponse, error){
	return s.service.ServiceGetDocumentsOfUser(ctx, in)
}

func (s *grpcServer) ServiceGetProfileAdditionalInfo(ctx context.Context, in *pb.ServiceGetProfileAdditionalInfoRequest) (*pb.GetProfileAdditionalInfoResponse, error){
	return s.service.ServiceGetProfileAdditionalInfo(ctx, in)
}

func (s *grpcServer) ServiceUpdateProfileStatus(ctx context.Context, in *pb.ServiceUpdateProfileStatusRequest) (*pb.UpdateProfileAdditionalInfoStatusResponse, error){
	return s.service.ServiceUpdateProfileStatus(ctx, in)
}