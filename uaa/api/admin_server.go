package api

import (
	"context"
	pb "gitlab.artin.ai/backend/courier-management/uaa/proto"
)

func (s *grpcServer) AdminLogin(ctx context.Context, in *pb.AdminLoginRequest) (*pb.AdminLoginResponse, error){
	return s.adminService.AdminLogin(ctx, in)
}
