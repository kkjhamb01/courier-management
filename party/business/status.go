package business

import (
	"context"
	"encoding/json"
	"time"

	"github.com/kkjhamb01/courier-management/common/logger"
	"github.com/kkjhamb01/courier-management/common/messaging"
	"github.com/kkjhamb01/courier-management/party/domain"
	pb "github.com/kkjhamb01/courier-management/party/proto"
	"github.com/kkjhamb01/courier-management/uaa/proto"
	"github.com/kkjhamb01/courier-management/uaa/security"
)

func (s *Service) GetCourierUserStatus(ctx context.Context, in *pb.GetCourierUserStatusRequest) (*pb.GetCourierUserStatusResponse, error) {
	logger.Infof("GetCourierUserStatus userId = %v", in.GetUserId())

	user := domain.CourierUser{}

	if err := s.db.Model(&domain.CourierUser{}).Where("id = ?", in.UserId).Find(&user).Error; err != nil {
		logger.Infof("GetCourierUserStatus userId = %v, error = %v", in.GetUserId(), err)
		return nil, proto.Internal.Error(err)
	}

	if user.ID == "" {
		logger.Infof("GetCourierUserStatus user not found %v", in.GetUserId())
		return nil, proto.NotFound.ErrorMsg("user not found")
	}

	logger.Infof("GetCourierUserStatus result = %v", user.Status)

	return &pb.GetCourierUserStatusResponse{
		Status: pb.UserStatus(user.Status),
	}, nil
}

func (s *Service) UpdateCourierUserStatus(ctx context.Context, in *pb.UpdateCourierUserStatusRequest) (*pb.UpdateCourierUserStatusResponse, error) {
	logger.Infof("UpdateCourierUserStatus userId = %v, status = %v", in.GetUserId(), in.GetStatus())

	result := s.db.Model(&domain.CourierUser{}).Where("id = ?", in.UserId).Update("status", in.GetStatus())

	if result.Error != nil {
		logger.Errorf("UpdateCourierUserStatus error in updating status", result.Error)
		return nil, proto.Internal.Error(result.Error)
	}

	if result.RowsAffected == 0 {
		logger.Debugf("UpdateCourierUserStatus user not found %v", in.UserId)
		return nil, proto.NotFound.ErrorMsg("user not found or update doesn't change a record")
	}

	s.publishUpdateCourierStatusEvent(&UpdateCourierStatusEvent{
		CourierId: in.GetUserId(),
		Status:    in.GetStatus(),
		Time:      time.Now().Format("2006-01-02 15:04:05"),
	})

	return &pb.UpdateCourierUserStatusResponse{}, nil
}

func (s *Service) GetClientUserStatus(ctx context.Context, in *pb.GetClientUserStatusRequest) (*pb.GetClientUserStatusResponse, error) {
	logger.Infof("GetClientUserStatus userId = %v", in.GetUserId())

	user := domain.ClientUser{}

	if err := s.db.Model(&domain.ClientUser{}).Where("id = ?", in.UserId).Find(&user).Error; err != nil {
		logger.Errorf("GetClientUserStatus error", err)
		return nil, proto.Internal.Error(err)
	}

	if user.ID == "" {
		return nil, proto.NotFound.ErrorMsg("user not found")
	}

	logger.Infof("GetClientUserStatus result = %v", user.Status)

	return &pb.GetClientUserStatusResponse{
		Status: pb.UserStatus(user.Status),
	}, nil
}

func (s *Service) UpdateClientUserStatus(ctx context.Context, in *pb.UpdateClientUserStatusRequest) (*pb.UpdateClientUserStatusResponse, error) {
	logger.Infof("UpdateClientUserStatus userId = %v, status = %v", in.GetUserId(), in.GetStatus())

	result := s.db.Model(&domain.ClientUser{}).Where("id = ?", in.UserId).Update("status", in.GetStatus())

	if result.Error != nil {
		logger.Errorf("UpdateClientUserStatus error in updating status", result.Error)
		return nil, proto.Internal.Error(result.Error)
	}

	if result.RowsAffected == 0 {
		logger.Debugf("UpdateClientUserStatus user not found %v", in.UserId)
		return nil, proto.NotFound.ErrorMsg("user not found or update doesn't change a record")
	}

	return &pb.UpdateClientUserStatusResponse{}, nil
}

func (s *Service) GetCourierUserStatusByToken(ctx context.Context, in *pb.GetCourierUserStatusByTokenRequest) (*pb.GetCourierUserStatusResponse, error) {
	tokenUser := ctx.Value("user").(security.User)

	logger.Infof("GetCourierUserStatusByToken userId = %v", tokenUser.Id)

	return s.GetCourierUserStatus(ctx, &pb.GetCourierUserStatusRequest{
		UserId: tokenUser.Id,
	})
}

func (s *Service) UpdateCourierUserStatusByToken(ctx context.Context, in *pb.UpdateCourierUserStatusByTokenRequest) (*pb.UpdateCourierUserStatusResponse, error) {

	tokenUser := ctx.Value("user").(security.User)

	logger.Infof("UpdateCourierUserStatusByToken userId = %v, status = %v", tokenUser.Id, in.Status)

	return s.UpdateCourierUserStatus(ctx, &pb.UpdateCourierUserStatusRequest{
		UserId: tokenUser.Id,
		Status: in.Status,
	})
}

func (s *Service) GetClientUserStatusByToken(ctx context.Context, in *pb.GetClientUserStatusByTokenRequest) (*pb.GetClientUserStatusResponse, error) {

	tokenUser := ctx.Value("user").(security.User)

	logger.Infof("GetClientUserStatusByToken userId = %v", tokenUser.Id)

	return s.GetClientUserStatus(ctx, &pb.GetClientUserStatusRequest{
		UserId: tokenUser.Id,
	})
}

func (s *Service) UpdateClientUserStatusByToken(ctx context.Context, in *pb.UpdateClientUserStatusByTokenRequest) (*pb.UpdateClientUserStatusResponse, error) {

	tokenUser := ctx.Value("user").(security.User)

	logger.Infof("UpdateClientUserStatusByToken userId = %v, status = %v", tokenUser.Id, in.Status)

	return s.UpdateClientUserStatus(ctx, &pb.UpdateClientUserStatusRequest{
		UserId: tokenUser.Id,
		Status: in.Status,
	})
}

type UpdateCourierStatusEvent struct {
	CourierId string
	Status    pb.UserStatus
	Time      string
}

func (s *Service) publishUpdateCourierStatusEvent(event *UpdateCourierStatusEvent) error {
	client := messaging.NatsClient()
	serialized, _ := json.Marshal(event)
	logger.Debugf("publishUpdateCourierStatusEvent %v", string(serialized))
	return client.Publish(messaging.TopicCourierStatusUpdate, serialized)
}
