package api

import (
	"context"
	pb "gitlab.artin.ai/backend/courier-management/notification/proto"
)

func (s *grpcServer) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error){
	return s.service.Register(ctx, in)
}

func (s *grpcServer) Unregister(ctx context.Context, in *pb.UnregisterRequest) (*pb.UnregisterResponse, error){
	return s.service.Unregister(ctx, in)
}
