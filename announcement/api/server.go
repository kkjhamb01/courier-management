package api

import (
	"context"
	pb "gitlab.artin.ai/backend/courier-management/announcement/proto"
)

func (s *grpcServer) CreateAnnouncement(ctx context.Context, in *pb.CreateAnnouncementRequest) (*pb.CreateAnnouncementResponse, error){
	return s.service.CreateAnnouncement(ctx, in)
}

func (s *grpcServer) AssignAnnouncementToUser(ctx context.Context, in *pb.AssignAnnouncementToUserRequest) (*pb.AssignAnnouncementToUserResponse, error){
	return s.service.AssignAnnouncementToUser(ctx, in)
}

func (s *grpcServer) GetAnnouncements(ctx context.Context, in *pb.GetAnnouncementsRequest) (*pb.GetAnnouncementsResponse, error){
	return s.service.GetAnnouncements(ctx, in)
}

func (s *grpcServer) GetUsers(ctx context.Context, in *pb.GetUsersRequest) (*pb.GetUsersResponse, error){
	return s.service.GetUsers(ctx, in)
}

func (s *grpcServer) GetAnnouncementsOfUser(ctx context.Context, in *pb.GetAnnouncementsOfUserRequest) (*pb.GetAnnouncementsOfUserResponse, error){
	return s.service.GetAnnouncementsOfUser(ctx, in)
}
