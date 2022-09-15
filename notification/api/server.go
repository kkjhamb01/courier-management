package api

import (
	"context"

	pb "github.com/kkjhamb01/courier-management/notification/proto"
)

func (s *grpcServer) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	return s.service.Register(ctx, in)
}

func (s *grpcServer) Unregister(ctx context.Context, in *pb.UnregisterRequest) (*pb.UnregisterResponse, error) {
	return s.service.Unregister(ctx, in)
}
