package api

import (
	"context"

	pb "github.com/kkjhamb01/courier-management/promotion/proto"
)

func (s *grpcServer) CreatePromotion(ctx context.Context, in *pb.CreatePromotionRequest) (*pb.CreatePromotionResponse, error) {
	return s.service.CreatePromotion(ctx, in)
}

func (s *grpcServer) AssignPromotionToUser(ctx context.Context, in *pb.AssignPromotionToUserRequest) (*pb.AssignPromotionToUserResponse, error) {
	return s.service.AssignPromotionToUser(ctx, in)
}

func (s *grpcServer) GetPromotions(ctx context.Context, in *pb.GetPromotionsRequest) (*pb.GetPromotionsResponse, error) {
	return s.service.GetPromotions(ctx, in)
}

func (s *grpcServer) GetUsers(ctx context.Context, in *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	return s.service.GetUsers(ctx, in)
}

func (s *grpcServer) AssignUserReferral(ctx context.Context, in *pb.AssignUserReferralRequest) (*pb.AssignUserReferralResponse, error) {
	return s.service.AssignUserReferral(ctx, in)
}

func (s *grpcServer) ApplyPromotion(ctx context.Context, in *pb.ApplyPromotionRequest) (*pb.ApplyPromotionResponse, error) {
	return s.service.ApplyPromotion(ctx, in)
}

func (s *grpcServer) GetPromotionsOfUser(ctx context.Context, in *pb.GetPromotionsOfUserRequest) (*pb.GetPromotionsOfUserResponse, error) {
	return s.service.GetPromotionsOfUser(ctx, in)
}
