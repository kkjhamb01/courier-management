package business

import (
	"context"
	"errors"
	"fmt"
	"gitlab.artin.ai/backend/courier-management/announcement/domain"
	pb "gitlab.artin.ai/backend/courier-management/announcement/proto"
	"gitlab.artin.ai/backend/courier-management/common/config"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	"gitlab.artin.ai/backend/courier-management/common/messaging"
	"gitlab.artin.ai/backend/courier-management/common/push"
	"gitlab.artin.ai/backend/courier-management/party/db"
	partypb "gitlab.artin.ai/backend/courier-management/party/proto"
	"gitlab.artin.ai/backend/courier-management/uaa/proto"
	"gitlab.artin.ai/backend/courier-management/uaa/security"
	"gorm.io/gorm"
	"time"
)

type Service struct {
	config                config.AnnouncementData
	db					  *gorm.DB
	jwtUtils      		  *security.JWTUtils
	partyApi      		  PartyAPI
}

func (s *Service) CreateAnnouncement(ctx context.Context, in *pb.CreateAnnouncementRequest) (*pb.CreateAnnouncementResponse, error){
	logger.Infof("CreateAnnouncement title = %v, text = %v, type = %v",
		in.GetTitle(), in.GetText(), in.GetType())

	announcement := &domain.Announcement{
		Title:        		in.GetTitle(),
		Text:   			in.GetText(),
		Type:      			domain.AnnouncementType(in.GetType()),
		MessageType:      	domain.AnnouncementMessageType(in.GetMessageType()),
		CreationTime: 		time.Now(),
	}

	err := s.db.Transaction(func(tx *gorm.DB) error {
		// insert announcement
		if err := tx.Create(announcement).Error; err != nil {
			logger.Errorf("CreateAnnouncement cannot perform update", err)
			return proto.Internal.Error(err)
		}
		return nil
	})

	if err != nil {
		logger.Errorf("CreateAnnouncement error in creating announcement", err)
		return nil, proto.Internal.Error(err)
	}

	logger.Debugf("CreateAnnouncement announcement created successfully %v", announcement.Id)

	return &pb.CreateAnnouncementResponse{
		AnnouncementId: announcement.Id,
	}, nil
}

func (s *Service) sendPushEvent(announcement domain.Announcement, userId string) error {
	logger.Debugf("sendPushEvent announcement = %v, userId = %v", announcement.Id, userId)
	user,err := s.partyApi.GetUserByUserId(userId, partypb.UserType_USER_TYPE_ALL)
	if err != nil{
		return err
	}
	if user == nil || user.Id == "" {
		return errors.New(fmt.Sprintf("cannot find user %v", userId))
	}

	evt1Data := &push.AnnouncementReceived{
		UserPhoneNumber: user.PhoneNumber,
		Id: announcement.Id,
		Title: announcement.Title,
		Text: announcement.Text,
		Type: int32(announcement.Type),
		MessageType: int32(announcement.MessageType),
		Time: announcement.CreationTime.Format("2006-01-02 15:04:05"),
	}
	pushClient := messaging.NatsClient()

	if pushClient == nil{
		logger.Debugf("cannot connect to nats")
		return errors.New("cannot connect to nats")
	} else {
		pushData, err := push.Encode(evt1Data)
		err = pushClient.Publish(messaging.TopicPushNotification, pushData)
		if err != nil {
			return err
		}
	}

	logger.Debugf("sendPushEvent sent successfully announcement = %v to userId = %v", announcement.Id, userId)

	return nil
}

func (s *Service) AssignAnnouncementToUser(ctx context.Context, in *pb.AssignAnnouncementToUserRequest) (*pb.AssignAnnouncementToUserResponse, error){
	logger.Infof("AssignAnnouncementToUser announcement_id = %v, user_id = %v", in.GetAnnouncementId(), in.GetUserIds())

	if in.GetUserIds() == nil || len(in.GetUserIds()) == 0 {
		logger.Infof("AssignAnnouncementToUser empty user ids")
		return nil, proto.InvalidArgument.ErrorMsg("empty user ids")
	}

	announcement := domain.Announcement{}

	if err := s.db.Model(&domain.Announcement{}).Where("id = ?", in.GetAnnouncementId()).Find(&announcement).Error; err != nil{
		logger.Errorf("AssignAnnouncementToUser error in query", err)
		return nil, proto.Internal.Error(err)
	}

	if announcement.Id == 0 {
		logger.Debugf("AssignAnnouncementToUser announcement not found")
		return nil, proto.NotFound.ErrorMsg("announcement not found")
	}

	announcementUsers := make([]*domain.AnnouncementUser, len(in.GetUserIds()))

	for i, userId := range in.GetUserIds(){
		announcementUsers[i] = &domain.AnnouncementUser{
			AnnouncementId:        	in.GetAnnouncementId(),
			UserId:   				userId,
		}
	}

	err := s.db.Transaction(func(tx *gorm.DB) error {
		// insert announcement-user
		if err := tx.CreateInBatches(announcementUsers, len(announcementUsers)).Error; err != nil {
			logger.Errorf("AssignAnnouncementToUser cannot perform update", err)
			return proto.Internal.Error(err)
		}
		return nil
	})

	if err != nil {
		logger.Errorf("AssignAnnouncementToUser error in creating announcement-user-assignment", err)
		return nil, proto.Internal.Error(err)
	}

	logger.Debugf("AssignAnnouncementToUser announcement created successfully")

	for _, userId := range in.GetUserIds(){
		if err := s.sendPushEvent(announcement, userId); err != nil{
			logger.Errorf("AssignAnnouncementToUser error in sending push event %v", err)
			return nil, proto.Internal.Error(err)
		}
	}

	return &pb.AssignAnnouncementToUserResponse{

	}, nil
}

func (s *Service) GetAnnouncements(ctx context.Context, in *pb.GetAnnouncementsRequest) (*pb.GetAnnouncementsResponse, error){

	logger.Infof("GetAnnouncements pagination = %v", in.GetPagination())

	var announcements []domain.Announcement

	var offset, limit, page int32
	var order = ""

	if in.GetPagination() != nil {
		page = in.GetPagination().GetPage()
		limit = in.GetPagination().GetLimit()
		if in.GetPagination().GetSort() != ""{
			order = in.GetPagination().GetSort()
			if in.GetPagination().GetSortType() == pb.SortType_SORT_TYPE_DESC {
				order = order + " desc"
			} else{
				order = order + " asc"
			}
		}
	}
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 20
	}
	offset = (page - 1) * limit
	if order == ""{
		order = "creation_time desc"
	}

	err := s.db.Limit(int(limit)).Offset(int(offset)).Order(order).Model(&domain.Announcement{}).Scan(&announcements).Error

	if err != nil{
		logger.Errorf("GetAnnouncements error in query", err)
		return nil, proto.Internal.Error(err)
	}

	var result = make([]*pb.Announcement, len(announcements))

	for i,announcement := range announcements{
		result[i] = s.announcementToDto(announcement)
	}

	logger.Infof("GetAnnouncements result = %v", result)

	return &pb.GetAnnouncementsResponse{
		Announcements: result,
	}, nil
}

func (s *Service) announcementToDto(announcement domain.Announcement) *pb.Announcement{
	return &pb.Announcement{
		Id: announcement.Id,
		Title: announcement.Title,
		Text: announcement.Text,
		Type: pb.AnnouncementType(announcement.Type),
		MessageType: pb.AnnouncementMessageType(announcement.MessageType),
		Time: announcement.CreationTime.Format("2006-01-02 15:04:05"),
	}
}

func (s *Service) GetUsers(ctx context.Context, in *pb.GetUsersRequest) (*pb.GetUsersResponse, error){

	logger.Infof("GetUsers user_id = %v, announcement_id = %v", in.GetUserId(), in.GetAnnouncementId())

	var users []domain.AnnouncementUser

	var query string

	if in.GetUserId() != "" {
		query = fmt.Sprintf("user_id = %v", in.GetUserId())
	} else if in.GetAnnouncementId() > 0 {
		query = fmt.Sprintf("announcement_id = %v", in.GetAnnouncementId())
	}

	err := s.db.Preload("Announcement").Where(query).Find(&users).Error

	if err != nil{
		logger.Errorf("GetUsers error in query", err)
		return nil, proto.Internal.Error(err)
	}

	var result = make([]*pb.AnnouncementUser, len(users))

	for i,user := range users {
		result[i] = &pb.AnnouncementUser{
			UserId : user.UserId,
			Announcement: s.announcementToDto(user.Announcement),
		}
	}

	logger.Infof("GetUsers result = %v", result)

	return &pb.GetUsersResponse{
		Users: result,
	}, nil
}

func (s *Service) GetAnnouncementsOfUser(ctx context.Context, in *pb.GetAnnouncementsOfUserRequest) (*pb.GetAnnouncementsOfUserResponse, error){

	tokenUser := ctx.Value("user").(security.User)

	logger.Infof("GetAnnouncementsOfUser user_id = %v", tokenUser.Id)

	var users []domain.AnnouncementUser

	var err = s.db.Preload("Announcement").Order("creation_time desc").Where("user_id = ?", tokenUser.Id).Find(&users).Error

	if err != nil{
		logger.Errorf("GetAnnouncementsOfUser error in query", err)
		return nil, proto.Internal.Error(err)
	}

	var result = make([]*pb.Announcement, len(users))

	for i,user := range users{
		result[i] = s.announcementToDto(user.Announcement)
	}

	logger.Infof("GetAnnouncementsOfUser result = %v", result)

	return &pb.GetAnnouncementsOfUserResponse{
		Announcements: result,
	}, nil
}


func NewService(config config.Data) *Service {
	dbInstance,err := db.NewOrm(config.Announcement.Database)
	if err != nil{
		logger.Fatalf("cannot connect to database", err)
	}
	jwtUtils, err := security.NewJWTUtils(config.Jwt)
	if err != nil{
		logger.Fatalf("cannot create jwtutils ", err)
	}
	partyApi := NewPartyAPI(config.Uaa)
	return &Service{
		config: config.Announcement,
		db: dbInstance,
		jwtUtils: &jwtUtils,
		partyApi: partyApi,
	}
}