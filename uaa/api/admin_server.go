package api

import (
	"context"

	pb "github.com/kkjhamb01/courier-management/uaa/proto"
)

func (s *grpcServer) AdminLogin(ctx context.Context, in *pb.AdminLoginRequest) (*pb.AdminLoginResponse, error) {
	return s.adminService.AdminLogin(ctx, in)
}
