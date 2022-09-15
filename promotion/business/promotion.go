package business

import (
	"context"
	"database/sql"
	"fmt"
	"math"
	"time"

	"github.com/kkjhamb01/courier-management/common/config"
	"github.com/kkjhamb01/courier-management/common/logger"
	"github.com/kkjhamb01/courier-management/party/db"
	"github.com/kkjhamb01/courier-management/promotion/domain"
	pb "github.com/kkjhamb01/courier-management/promotion/proto"
	"github.com/kkjhamb01/courier-management/uaa/proto"
	"github.com/kkjhamb01/courier-management/uaa/security"
	"gorm.io/gorm"
)

type Service struct {
	config     config.PromotionData
	db         *gorm.DB
	jwtUtils   *security.JWTUtils
	dateFormat string
}

func (s *Service) CreatePromotion(ctx context.Context, in *pb.CreatePromotionRequest) (*pb.CreatePromotionResponse, error) {
	logger.Infof("CreatePromotion name = %v, start_date = %v, exp_date = %v, type = %v, discount_percentage = %v, discount_value = %v",
		in.GetName(), in.GetStartDate(), in.GetExpDate(),
		in.GetType(), in.GetDiscountPercentage(), in.GetDiscountValue())

	startDate, err := time.Parse(s.dateFormat, in.GetStartDate())

	if err != nil {
		logger.Debugf("CreatePromotion invalid start date")
		return nil, proto.InvalidArgument.ErrorMsg("invalid start date")
	}

	expDate, err := time.Parse(s.dateFormat, in.GetExpDate())

	if err != nil {
		logger.Debugf("CreatePromotion invalid exp date")
		return nil, proto.InvalidArgument.ErrorMsg("invalid exp date")
	}

	promotion := &domain.Promotion{
		Name:               in.GetName(),
		StartDate:          sql.NullTime{Time: startDate, Valid: in.GetStartDate() != ""},
		ExpDate:            sql.NullTime{Time: expDate, Valid: in.GetExpDate() != ""},
		DiscountPercentage: sql.NullFloat64{Float64: in.GetDiscountPercentage(), Valid: in.GetDiscountPercentage() > 0},
		DiscountValue:      sql.NullFloat64{Float64: in.GetDiscountValue(), Valid: in.GetDiscountValue() > 0},
		Type:               domain.PromotionType(in.GetType()),
	}

	err = s.db.Transaction(func(tx *gorm.DB) error {
		// insert promotion
		if err := tx.Create(promotion).Error; err != nil {
			logger.Errorf("CreatePromotion cannot perform update", err)
			return proto.Internal.Error(err)
		}
		return nil
	})

	if err != nil {
		logger.Errorf("CreatePromotion error in creating promotion", err)
		return nil, proto.Internal.Error(err)
	}

	logger.Debugf("CreatePromotion promotion created successfully %v", promotion.Id)

	return &pb.CreatePromotionResponse{
		PromotionId: promotion.Id,
	}, nil
}

func (s *Service) AssignPromotionToUser(ctx context.Context, in *pb.AssignPromotionToUserRequest) (*pb.AssignPromotionToUserResponse, error) {
	logger.Infof("AssignPromotionToUser promotion_id = %v, user_id = %v", in.GetPromotionId(), in.GetUserIds())

	promotion := domain.Promotion{}

	if err := s.db.Model(&domain.Promotion{}).Where("id = ?", in.GetPromotionId()).Find(&promotion).Error; err != nil {
		logger.Errorf("AssignPromotionToUser error in query", err)
		return nil, proto.Internal.Error(err)
	}

	if promotion.Id == 0 {
		logger.Debugf("AssignPromotionToUser promotion not found")
		return nil, proto.NotFound.ErrorMsg("promotion not found")
	}

	promotionUsers := make([]*domain.PromotionUser, len(in.GetUserIds()))

	for i, userId := range in.GetUserIds() {
		promotionUsers[i] = &domain.PromotionUser{
			PromotionId: in.GetPromotionId(),
			UserId:      userId,
			Status:      domain.PROMOTION_STATUS_AVAILABLE,
		}
	}

	err := s.db.Transaction(func(tx *gorm.DB) error {
		// insert promotion-user
		if err := tx.CreateInBatches(promotionUsers, len(promotionUsers)).Error; err != nil {
			logger.Errorf("AssignPromotionToUser cannot perform update", err)
			return proto.Internal.Error(err)
		}
		return nil
	})

	if err != nil {
		logger.Errorf("AssignPromotionToUser error in creating promotion-user-assignment", err)
		return nil, proto.Internal.Error(err)
	}

	logger.Debugf("AssignPromotionToUser promotion created successfully")

	return &pb.AssignPromotionToUserResponse{}, nil
}

func (s *Service) GetPromotions(ctx context.Context, in *pb.GetPromotionsRequest) (*pb.GetPromotionsResponse, error) {

	logger.Infof("GetPromotions pagination = %v", in.GetPagination())

	var promotions []domain.Promotion

	var offset, limit, page int32
	var order = ""

	if in.GetPagination() != nil {
		page = in.GetPagination().GetPage()
		limit = in.GetPagination().GetLimit()
		if in.GetPagination().GetSort() != "" {
			order = in.GetPagination().GetSort()
			if in.GetPagination().GetSortType() == pb.SortType_SORT_TYPE_DESC {
				order = order + " desc"
			} else {
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
	if order == "" {
		order = "creation_time desc"
	}

	err := s.db.Limit(int(limit)).Offset(int(offset)).Order(order).Model(&domain.Promotion{}).Scan(&promotions).Error

	if err != nil {
		logger.Errorf("GetPromotions error in query", err)
		return nil, proto.Internal.Error(err)
	}

	if len(promotions) == 0 {
		logger.Debugf("GetPromotions no promotion found")
		return nil, proto.NotFound.ErrorMsg("no promotion found")
	}

	var result = make([]*pb.Promotion, len(promotions))

	for i, promotion := range promotions {
		result[i] = s.promotionToDto(promotion)
	}

	logger.Infof("GetPromotions result = %v", result)

	return &pb.GetPromotionsResponse{
		Promotions: result,
	}, nil
}

func (s *Service) promotionToDto(promotion domain.Promotion) *pb.Promotion {
	var startDate, expDate string

	if promotion.StartDate.Valid {
		startDate = promotion.StartDate.Time.Format(s.dateFormat)
	}
	if promotion.ExpDate.Valid {
		expDate = promotion.ExpDate.Time.Format(s.dateFormat)
	}

	return &pb.Promotion{
		Id:                 promotion.Id,
		Name:               promotion.Name,
		StartDate:          startDate,
		ExpDate:            expDate,
		Type:               pb.PromotionType(promotion.Type),
		DiscountPercentage: promotion.DiscountPercentage.Float64,
		DiscountValue:      promotion.DiscountValue.Float64,
	}
}

func (s *Service) GetUsers(ctx context.Context, in *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {

	logger.Infof("GetUsers user_id = %v, promotion_id = %v", in.GetUserId(), in.GetPromotionId())

	var users []domain.PromotionUser

	var query string

	if in.GetUserId() != "" {
		query = fmt.Sprintf("user_id = %v", in.GetUserId())
	} else if in.GetPromotionId() > 0 {
		query = fmt.Sprintf("promotion_id = %v", in.GetPromotionId())
	}

	err := s.db.Preload("Promotion").Where(query).Find(&users).Error

	if err != nil {
		logger.Errorf("GetUsers error in query", err)
		return nil, proto.Internal.Error(err)
	}

	if len(users) == 0 {
		logger.Debugf("GetUsers no promotion-user found")
		return nil, proto.NotFound.ErrorMsg("no promotion-user found")
	}

	var result = make([]*pb.PromotionUser, len(users))

	for i, user := range users {
		var promotionHistory = domain.PromotionHistory{}
		var transaction *pb.Transaction

		s.db.Where("promotion_id = ? AND user_id = ?", user.PromotionId, user.UserId).Find(&promotionHistory)

		if promotionHistory.PromotionId != 0 {
			transaction = &pb.Transaction{
				TransactionId: promotionHistory.TransactionId,
				Date:          promotionHistory.Date.String(),
			}
		}

		result[i] = &pb.PromotionUser{
			UserId:      user.UserId,
			Metadata:    user.Metadata.String,
			Status:      pb.PromotionUserStatus(user.Status),
			Promotion:   s.promotionToDto(user.Promotion),
			Transaction: transaction,
		}
	}

	logger.Infof("GetUsers result = %v", result)

	return &pb.GetUsersResponse{
		Users: result,
	}, nil
}

func (s *Service) AssignUserReferral(ctx context.Context, in *pb.AssignUserReferralRequest) (*pb.AssignUserReferralResponse, error) {

	logger.Infof("AssignUserReferral user_id = %v, referral = %v", in.GetUserId(), in.GetReferral())

	startDate := time.Now()

	promotion := &domain.Promotion{
		Name:               s.config.Referral.PromotionName,
		StartDate:          sql.NullTime{Time: startDate, Valid: true},
		DiscountPercentage: sql.NullFloat64{Float64: s.config.Referral.DiscountPercentage, Valid: s.config.Referral.DiscountPercentage > 0},
		DiscountValue:      sql.NullFloat64{Float64: s.config.Referral.DiscountValue, Valid: s.config.Referral.DiscountValue > 0},
		Type:               domain.PROMOTION_TYPE_REFERRAL,
	}

	err := s.db.Transaction(func(tx *gorm.DB) error {
		// insert promotion
		if err := tx.Create(promotion).Error; err != nil {
			logger.Errorf("AssignUserReferral cannot perform update", err)
			return proto.Internal.Error(err)
		}
		return nil
	})

	if err != nil {
		logger.Errorf("AssignUserReferral error in creating promotion", err)
		return nil, proto.Internal.Error(err)
	}

	promotionUser := &domain.PromotionUser{
		PromotionId: promotion.Id,
		UserId:      in.GetUserId(),
		Metadata:    sql.NullString{String: in.GetReferral(), Valid: in.GetReferral() != ""},
		Status:      domain.PROMOTION_STATUS_AVAILABLE,
	}

	err = s.db.Transaction(func(tx *gorm.DB) error {
		// insert promotion-user
		if err := tx.Create(promotionUser).Error; err != nil {
			logger.Errorf("AssignUserReferral cannot perform update", err)
			return proto.Internal.Error(err)
		}
		return nil
	})

	if err != nil {
		logger.Errorf("AssignUserReferral error in creating promotion-user", err)
		return nil, proto.Internal.Error(err)
	}

	logger.Debugf("AssignUserReferral promotion created successfully %v", promotion.Id)

	return &pb.AssignUserReferralResponse{
		PromotionId: promotion.Id,
	}, nil
}

func (s *Service) ApplyPromotion(ctx context.Context, in *pb.ApplyPromotionRequest) (*pb.ApplyPromotionResponse, error) {

	logger.Infof("ApplyPromotion promotion_id = %v, user_id = %v, transaction_id = %v, total_payment = %v",
		in.GetPromotionId(), in.GetUserId(), in.GetTransactionId(), in.GetTotalPayment())

	promotion := domain.Promotion{}

	if err := s.db.Model(&domain.Promotion{}).Where("id = ?", in.GetPromotionId()).Find(&promotion).Error; err != nil {
		logger.Errorf("ApplyPromotion error in query", err)
		return nil, proto.Internal.Error(err)
	}

	if promotion.Id == 0 {
		logger.Debugf("ApplyPromotion promotion not found")
		return nil, proto.NotFound.ErrorMsg("promotion not found")
	}

	today := time.Now().Round(0)

	if promotion.StartDate.Valid {
		if promotion.StartDate.Time.After(today) {
			logger.Debugf("ApplyPromotion promotion will start in the future")
			return nil, proto.PromotionWillStartInTheFuture.ErrorNoMsg()
		}
	}

	if promotion.ExpDate.Valid {
		if promotion.ExpDate.Time.Before(today) {
			logger.Debugf("ApplyPromotion promotion is expired")
			return nil, proto.PromotionIsExpired.ErrorNoMsg()
		}
	}

	promotionUser := domain.PromotionUser{}

	if err := s.db.Model(&domain.PromotionUser{}).Where("promotion_id = ? AND user_id = ?",
		in.GetPromotionId(), in.GetUserId()).Find(&promotionUser).Error; err != nil {
		logger.Errorf("ApplyPromotion error in query", err)
		return nil, proto.Internal.Error(err)
	}

	if promotionUser.UserId == "" {
		logger.Debugf("ApplyPromotion this user has not such promotion")
		return nil, proto.NotFound.ErrorMsg("this user has not such promotion")
	}

	if promotionUser.Status != domain.PROMOTION_STATUS_AVAILABLE {
		logger.Debugf("ApplyPromotion this promotion is not available for this user")
		return nil, proto.PromotionNotAvailable.ErrorMsg("this promotion is not available for this user")
	}

	var totalPayment = in.GetTotalPayment()

	var d1 = promotion.DiscountValue.Float64
	var d2 = promotion.DiscountPercentage.Float64 * totalPayment / float64(100)
	var discount float64

	if d1 < 1 {
		discount = d2
	} else if d2 < 1 {
		discount = d1
	} else {
		discount = math.Min(d1, d2)
	}

	if discount < 1 {
		logger.Debugf("ApplyPromotion discount is zero")
		return nil, proto.ZeroDiscount.ErrorMsg("discount is zero")
	}

	if discount >= totalPayment {
		logger.Debugf("ApplyPromotion discount is greater than total payment")
		return nil, proto.DiscountIsGreaterThanTotalPayment.ErrorNoMsg()
	}

	err := s.db.Transaction(func(tx *gorm.DB) error {
		// update promotion-user

		updatedPromotionUser := domain.PromotionUser{
			PromotionId: in.PromotionId,
			UserId:      in.UserId,
			Status:      domain.PROMOTION_STATUS_CONSUMED,
		}

		if err := tx.Model(&updatedPromotionUser).Updates(&updatedPromotionUser).Error; err != nil {
			logger.Errorf("ApplyPromotion cannot perform update user promotion", err)
			return err
		}

		// insert promotion-history

		if err := tx.Create(&domain.PromotionHistory{
			PromotionId:   in.PromotionId,
			UserId:        in.UserId,
			TransactionId: in.TransactionId,
			Date:          time.Now(),
		}).Error; err != nil {
			logger.Errorf("ApplyPromotion cannot perform insert history", err)
			return proto.Internal.Error(err)
		}

		return nil
	})

	if err != nil {
		logger.Errorf("ApplyPromotion error in applying discount", err)
		return nil, proto.Internal.Error(err)
	}

	return &pb.ApplyPromotionResponse{
		DiscountPercentage: float64(100) * discount / totalPayment,
		DiscountValue:      discount,
	}, nil
}

func (s *Service) GetPromotionsOfUser(ctx context.Context, in *pb.GetPromotionsOfUserRequest) (*pb.GetPromotionsOfUserResponse, error) {

	logger.Infof("GetPromotionsOfUser type = %v", in.GetType())

	tokenUser := ctx.Value("user").(security.User)

	logger.Infof("GetPromotionsOfUser user_id = %v", tokenUser.Id)

	var users []domain.PromotionUser

	var err error

	if in.GetType() == pb.PromotionUserStatus_UNKNOWN_PROMOTION_USER_STATUS {
		err = s.db.Preload("Promotion").Where("user_id = ?", tokenUser.Id).Find(&users).Error
	} else {
		err = s.db.Preload("Promotion").Where("user_id = ? AND status = ?",
			tokenUser.Id, domain.PromotionStatus(in.GetType())).Find(&users).Error
	}

	if err != nil {
		logger.Errorf("GetPromotionsOfUser error in query", err)
		return nil, proto.Internal.Error(err)
	}

	if len(users) == 0 {
		logger.Debugf("GetPromotionsOfUser no promotion found")
		return nil, proto.NotFound.ErrorMsg("no promotion found")
	}

	var result = make([]*pb.Promotion, len(users))

	for i, user := range users {
		result[i] = s.promotionToDto(user.Promotion)
	}

	logger.Infof("GetPromotionsOfUser result = %v", result)

	return &pb.GetPromotionsOfUserResponse{
		Promotions: result,
	}, nil
}

func NewService(config config.PromotionData, jwtConfig config.JwtData) *Service {
	dbInstance, err := db.NewOrm(config.Database)
	if err != nil {
		logger.Fatalf("cannot connect to database", err)
	}
	jwtUtils, err := security.NewJWTUtils(jwtConfig)
	if err != nil {
		logger.Fatalf("cannot create jwtutils ", err)
	}
	return &Service{
		config:     config,
		db:         dbInstance,
		jwtUtils:   &jwtUtils,
		dateFormat: "2006-01-02",
	}
}
