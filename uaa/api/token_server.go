package api

import (
	"context"

	pb "github.com/kkjhamb01/courier-management/uaa/proto"
)

func (s *grpcServer) RefreshToken(ctx context.Context, in *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error) {
	return s.tokenService.RefreshToken(ctx, in)
}

func (s *grpcServer) GetJwks(ctx context.Context, in *pb.GetJwksRequest) (*pb.GetJwksResponse, error) {
	return s.tokenService.GetJwks(ctx, in)
}
