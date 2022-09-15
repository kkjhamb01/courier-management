package api

import (
	"context"

	pb "github.com/kkjhamb01/courier-management/party/proto"
)

func (s *grpcServer) GetCourierUserStatus(ctx context.Context, in *pb.GetCourierUserStatusRequest) (*pb.GetCourierUserStatusResponse, error) {
	return s.service.GetCourierUserStatus(ctx, in)
}

func (s *grpcServer) UpdateCourierUserStatus(ctx context.Context, in *pb.UpdateCourierUserStatusRequest) (*pb.UpdateCourierUserStatusResponse, error) {
	return s.service.UpdateCourierUserStatus(ctx, in)
}

func (s *grpcServer) GetClientUserStatus(ctx context.Context, in *pb.GetClientUserStatusRequest) (*pb.GetClientUserStatusResponse, error) {
	return s.service.GetClientUserStatus(ctx, in)
}

func (s *grpcServer) UpdateClientUserStatus(ctx context.Context, in *pb.UpdateClientUserStatusRequest) (*pb.UpdateClientUserStatusResponse, error) {
	return s.service.UpdateClientUserStatus(ctx, in)
}

func (s *grpcServer) GetCourierUserStatusByToken(ctx context.Context, in *pb.GetCourierUserStatusByTokenRequest) (*pb.GetCourierUserStatusResponse, error) {
	return s.service.GetCourierUserStatusByToken(ctx, in)
}

func (s *grpcServer) UpdateCourierUserStatusByToken(ctx context.Context, in *pb.UpdateCourierUserStatusByTokenRequest) (*pb.UpdateCourierUserStatusResponse, error) {
	return s.service.UpdateCourierUserStatusByToken(ctx, in)
}

func (s *grpcServer) GetClientUserStatusByToken(ctx context.Context, in *pb.GetClientUserStatusByTokenRequest) (*pb.GetClientUserStatusResponse, error) {
	return s.service.GetClientUserStatusByToken(ctx, in)
}

func (s *grpcServer) UpdateClientUserStatusByToken(ctx context.Context, in *pb.UpdateClientUserStatusByTokenRequest) (*pb.UpdateClientUserStatusResponse, error) {
	return s.service.UpdateClientUserStatusByToken(ctx, in)
}
