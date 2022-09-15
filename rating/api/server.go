package api

import (
	"context"
	pb "gitlab.artin.ai/backend/courier-management/rating/proto"
)

func (s *grpcServer) CreateCourierRating(ctx context.Context, in *pb.CreateCourierRatingRequest) (*pb.CreateCourierRatingResponse, error){
	return s.service.CreateCourierRating(ctx, in)
}

func (s *grpcServer) CreateClientRating(ctx context.Context, in *pb.CreateClientRatingRequest) (*pb.CreateClientRatingResponse, error){
	return s.service.CreateClientRating(ctx, in)
}

func (s *grpcServer) GetCourierRating(ctx context.Context, in *pb.GetCourierRatingRequest) (*pb.GetCourierRatingResponse, error){
	return s.service.GetCourierRating(ctx, in)
}

func (s *grpcServer) GetClientRating(ctx context.Context, in *pb.GetClientRatingRequest) (*pb.GetClientRatingResponse, error){
	return s.service.GetClientRating(ctx, in)
}

func (s *grpcServer) GetCourierRatingStat(ctx context.Context, in *pb.GetCourierRatingStatRequest) (*pb.GetCourierRatingStatResponse, error){
	return s.service.GetCourierRatingStat(ctx, in)
}

func (s *grpcServer) GetClientRatingStat(ctx context.Context, in *pb.GetClientRatingStatRequest) (*pb.GetClientRatingStatResponse, error){
	return s.service.GetClientRatingStat(ctx, in)
}

func (s *grpcServer) GetCourierRatingStatByToken(ctx context.Context, in *pb.GetCourierRatingStatByTokenRequest) (*pb.GetCourierRatingStatByTokenResponse, error){
	return s.service.GetCourierRatingStatByToken(ctx, in)
}

func (s *grpcServer) GetClientRatingStatByToken(ctx context.Context, in *pb.GetClientRatingStatByTokenRequest) (*pb.GetClientRatingStatByTokenResponse, error){
	return s.service.GetClientRatingStatByToken(ctx, in)
}