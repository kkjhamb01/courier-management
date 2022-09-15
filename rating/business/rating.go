package business

import (
	"context"
	"database/sql"
	"fmt"
	"gitlab.artin.ai/backend/courier-management/common/config"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	delivery "gitlab.artin.ai/backend/courier-management/grpc/delivery/go"
	"gitlab.artin.ai/backend/courier-management/party/db"
	"gitlab.artin.ai/backend/courier-management/rating/domain"
	pb "gitlab.artin.ai/backend/courier-management/rating/proto"
	proto "gitlab.artin.ai/backend/courier-management/uaa/proto"
	"gitlab.artin.ai/backend/courier-management/uaa/security"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"math"
	"time"
)

type Service struct {
	config                config.RatingData
	db					  *gorm.DB
	jwtUtils      		  *security.JWTUtils
}

func (s *Service) CreateCourierRating(ctx context.Context, in *pb.CreateCourierRatingRequest) (*pb.CreateCourierRatingResponse, error){

	tokenUser := ctx.Value("user").(security.User)

	logger.Infof("CreateCourierRating clientId = %v, ride_id = %v, rating = %v ",
		tokenUser.Id, in.GetRideId(), in.GetRating())

	rideInfo,err := s.getRideDetails(in.GetRideId(), in.GetAccessToken())

	if err != nil{
		logger.Debugf("CreateCourierRating error in retrieving ride %v", err)
		if e, ok := status.FromError(err); ok {
			if e.Code() == 2{
				return nil, proto.NotFound.ErrorMsg("ride not found")
			}
		}
		return nil, err
	}

	if rideInfo.ClientId != tokenUser.Id{
		logger.Debugf("CreateCourierRating client %v does not match to the token, client id of the ride is %v",
			tokenUser.Id, rideInfo.ClientId)
		return nil, proto.InvalidArgument.ErrorMsg("client does not match to the token")
	}

	rate := &domain.RateItem{
		Rater:   		tokenUser.Id,
		Rated:    		rideInfo.CourierId,
		Ride:     		in.RideId,
		RaterType: 		domain.RATEE_TYPE_CLIENT,
		RatedType:      domain.RATEE_TYPE_COURIER,
		RateValue:   	int32(in.Rating.Value),
		Message:     	sql.NullString{String: in.Rating.Message, Valid: in.Rating.Message != ""},
	}

	err = s.db.Transaction(func(tx *gorm.DB) error {
		// insert rate item
		if err := tx.Create(rate).Error; err != nil {
			logger.Errorf("CreateCourierRating cannot perform update", err)
			return proto.Internal.Error(err)
		}

		// insert feedbacks
		var feedbacks []*domain.RateItemFeedback

		if len(in.GetRating().GetPositiveFeedbacks()) > 0{
			for _,feedback := range in.GetRating().GetPositiveFeedbacks(){
				feedbacks = append(feedbacks, &domain.RateItemFeedback{
					RateId:   		rate.ID,
					Feedback:    	int(feedback),
					Positive: domain.FEEDBACK_TYPE_POSITIVE,
				})
			}
		}

		if len(in.GetRating().GetNegativeFeedbacks()) > 0{
			for _,feedback := range in.GetRating().GetNegativeFeedbacks(){
				feedbacks = append(feedbacks, &domain.RateItemFeedback{
					RateId:   		rate.ID,
					Feedback:    	int(feedback),
					Positive: domain.FEEDBACK_TYPE_NEGATIVE,
				})
			}
		}

		if len(feedbacks) > 0{
			if err := tx.Create(feedbacks).Error; err != nil {
				logger.Errorf("CreateCourierRating cannot perform update", err)
				return proto.Internal.Error(err)
			}
		}

		// update rate stat
		currentRate := domain.Rate{}
		s.db.Model(&currentRate).Where("rated = ?", rideInfo.CourierId).Find(&currentRate)
		updatedRate := domain.Rate{
			Rated: rideInfo.CourierId,
		}
		updatedRate.RateTotal = int64(in.Rating.Value) + currentRate.RateTotal
		updatedRate.RateCount = currentRate.RateCount + 1

		if currentRate.Rated == ""{
			if err := tx.Model(&updatedRate).Create(&updatedRate).Error; err != nil {
				return err
			}
		} else {
			if err := tx.Model(&updatedRate).Updates(&updatedRate).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		logger.Errorf("CreateCourierRating error in creating rating", err)
		return nil, proto.Internal.Error(err)
	}

	return &pb.CreateCourierRatingResponse{

	}, nil
}

func (s *Service) CreateClientRating(ctx context.Context, in *pb.CreateClientRatingRequest) (*pb.CreateClientRatingResponse, error){
	tokenUser := ctx.Value("user").(security.User)

	logger.Infof("CreateClientRating clientId = %v, ride_id = %v, rating = %v ",
		tokenUser.Id, in.GetRideId(), in.GetRating())

	rideInfo,err := s.getRideDetails(in.GetRideId(), in.GetAccessToken())

	if err != nil{
		logger.Debugf("CreateClientRating error in retrieving ride %v", err)
		if e, ok := status.FromError(err); ok {
			if e.Code() == 2{
				return nil, proto.NotFound.ErrorMsg("ride not found")
			}
		}
		return nil, err
	}

	if rideInfo.CourierId != tokenUser.Id{
		logger.Debugf("CreateClientRating courier %v does not match to the token, courier id of the ride is %v",
			tokenUser.Id, rideInfo.CourierId)
		return nil, proto.InvalidArgument.ErrorMsg("courier does not match to the token")
	}

	rate := &domain.RateItem{
		Rater:   		tokenUser.Id,
		Rated:    		rideInfo.ClientId,
		Ride:     		in.RideId,
		RaterType: 		domain.RATEE_TYPE_COURIER,
		RatedType:      domain.RATEE_TYPE_CLIENT,
		RateValue:   	int32(in.Rating.Value),
		Message:     	sql.NullString{String: in.Rating.Message, Valid: in.Rating.Message != ""},
	}

	err = s.db.Transaction(func(tx *gorm.DB) error {
		// insert rate item
		if err := tx.Create(rate).Error; err != nil {
			logger.Errorf("CreateClientRating cannot perform update", err)
			return proto.Internal.Error(err)
		}

		// insert feedbacks
		var feedbacks []*domain.RateItemFeedback

		if len(in.GetRating().GetPositiveFeedbacks()) > 0{
			for _,feedback := range in.GetRating().GetPositiveFeedbacks(){
				feedbacks = append(feedbacks, &domain.RateItemFeedback{
					RateId:   		rate.ID,
					Feedback:    	int(feedback),
					Positive: domain.FEEDBACK_TYPE_POSITIVE,
				})
			}
		}

		if len(in.GetRating().GetNegativeFeedbacks()) > 0{
			for _,feedback := range in.GetRating().GetNegativeFeedbacks(){
				feedbacks = append(feedbacks, &domain.RateItemFeedback{
					RateId:   		rate.ID,
					Feedback:    	int(feedback),
					Positive: domain.FEEDBACK_TYPE_NEGATIVE,
				})
			}
		}

		if len(feedbacks) > 0{
			if err := tx.Create(feedbacks).Error; err != nil {
				logger.Errorf("CreateClientRating cannot perform update", err)
				return proto.Internal.Error(err)
			}
		}

		// update rate stat
		currentRate := domain.Rate{}
		s.db.Model(&currentRate).Where("rated = ?", rideInfo.ClientId).Find(&currentRate)
		updatedRate := domain.Rate{
			Rated: rideInfo.ClientId,
		}
		updatedRate.RateTotal = int64(in.Rating.Value) + currentRate.RateTotal
		updatedRate.RateCount = currentRate.RateCount + 1

		if currentRate.Rated == ""{
			if err := tx.Model(&updatedRate).Create(&updatedRate).Error; err != nil {
				return err
			}
		} else {
			if err := tx.Model(&updatedRate).Updates(&updatedRate).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		logger.Errorf("CreateClientRating error in creating rating", err)
		return nil, proto.Internal.Error(err)
	}

	return &pb.CreateClientRatingResponse{

	}, nil
}

func (s *Service) GetCourierRating(ctx context.Context, in *pb.GetCourierRatingRequest) (*pb.GetCourierRatingResponse, error){
	logger.Infof("GetCourierRating client_id = %v", in.GetCourierId())

	var rates []domain.RateItem

	err := s.db.Model(&domain.RateItem{}).
		Select("id, rater, ride, rate_value, message").
		Where("rated=? AND rated_type=?", in.GetCourierId(), domain.RATEE_TYPE_COURIER).
		Scan(&rates).Error

	if err != nil {
		logger.Debugf("GetCourierRating rate error client_id = %v, err = %v", in.GetCourierId(), err)
		return nil, proto.Internal.Error(err)
	}

	if len(rates) == 0{
		logger.Debugf("GetCourierRating rate not found %v", in.GetCourierId())
		return nil, proto.NotFound.ErrorMsg("no rate found")
	}

	var result = make([]*pb.CourierRated, len(rates))

	for i,rate := range rates{
		var feedbacks0 []*domain.RateItemFeedback
		if err := s.db.Model(&domain.RateItemFeedback{}).
			Select("feedback", "positive").
			Where("rate_id=?", rate.ID).
			Scan(&feedbacks0).Error; err != nil{
			logger.Debugf("GetCourierRating rate error courier_id = %v, err = %v", in.GetCourierId(), err)
			return nil, proto.Internal.Error(err)
		}

		var feedbacksPositive = make([]pb.CourierRatingPositiveFeedback, 0)
		var feedbacksNegative = make([]pb.CourierRatingNegativeFeedback, 0)

		for _, feedback := range feedbacks0{
			if feedback.Positive == domain.FEEDBACK_TYPE_POSITIVE{
				feedbacksPositive = append(feedbacksPositive,
					pb.CourierRatingPositiveFeedback(feedback.Feedback))
			} else if feedback.Positive == domain.FEEDBACK_TYPE_NEGATIVE{
				feedbacksNegative = append(feedbacksNegative,
					pb.CourierRatingNegativeFeedback(feedback.Feedback))
			}
		}

		result[i] = &pb.CourierRated{
			ClientId: rate.Rater,
			Ride: rate.Ride,
			Rating: &pb.CourierRating{
				Value: pb.RateValue(rate.RateValue),
				Message: rate.Message.String,
				PositiveFeedbacks: feedbacksPositive,
				NegativeFeedbacks: feedbacksNegative,
			},
		}
	}

	logger.Debugf("GetCourierRating result = %v", result)

	return &pb.GetCourierRatingResponse{
		Rates: result,
	}, nil
}

func (s *Service) GetClientRating(ctx context.Context, in *pb.GetClientRatingRequest) (*pb.GetClientRatingResponse, error){
	logger.Infof("GetClientRating client_id = %v", in.GetClientId())

	var rates []domain.RateItem

	err := s.db.Model(&domain.RateItem{}).
		Select("id, rater, ride, rate_value, message").
		Where("rated=? AND rated_type=?", in.ClientId, domain.RATEE_TYPE_CLIENT).
		Scan(&rates).Error

	if err != nil {
		logger.Debugf("GetClientRating rate error client_id = %v, err = %v", in.ClientId, err)
		return nil, proto.Internal.Error(err)
	}

	if len(rates) == 0{
		logger.Debugf("GetClientRating rate not found %v", in.ClientId)
		return nil, proto.NotFound.ErrorMsg("no rate found")
	}

	var result = make([]*pb.ClientRated, len(rates))

	for i,rate := range rates{
		var feedbacks0 []*domain.RateItemFeedback
		if err := s.db.Model(&domain.RateItemFeedback{}).
			Select("feedback", "positive").
			Where("rate_id=?", rate.ID).
			Scan(&feedbacks0).Error; err != nil{
			logger.Debugf("GetClientRating rate error client_id = %v, err = %v", in.GetClientId(), err)
			return nil, proto.Internal.Error(err)
		}

		var feedbacksPositive = make([]pb.ClientRatingPositiveFeedback, 0)
		var feedbacksNegative = make([]pb.ClientRatingNegativeFeedback, 0)

		for _, feedback := range feedbacks0{
			if feedback.Positive == domain.FEEDBACK_TYPE_POSITIVE{
				feedbacksPositive = append(feedbacksPositive,
					pb.ClientRatingPositiveFeedback(feedback.Feedback))
			} else if feedback.Positive == domain.FEEDBACK_TYPE_NEGATIVE{
				feedbacksNegative = append(feedbacksNegative,
					pb.ClientRatingNegativeFeedback(feedback.Feedback))
			}
		}

		result[i] = &pb.ClientRated{
			CourierId: rate.Rater,
			Ride: rate.Ride,
			Rating: &pb.ClientRating{
				Value: pb.RateValue(rate.RateValue),
				Message: rate.Message.String,
				PositiveFeedbacks: feedbacksPositive,
				NegativeFeedbacks: feedbacksNegative,
			},
		}
	}

	logger.Debugf("GetClientRating result = %v", result)

	return &pb.GetClientRatingResponse{
		Rates: result,
	}, nil
}

func (s *Service) GetCourierRatingStat(ctx context.Context, in *pb.GetCourierRatingStatRequest) (*pb.GetCourierRatingStatResponse, error){
	logger.Infof("GetCourierRatingStat courier_id = %v", in.GetCourierId())

	rate := domain.Rate{}
	err := s.db.Model(&rate).Where("rated = ?", in.CourierId).Find(&rate).Error

	if err != nil {
		logger.Debugf("GetCourierRatingStat stat error, courier_id = %v, error = %v", in.GetCourierId(), err)
		return nil, proto.Internal.Error(err)
	}

	var numberOfRates int64
	var scoreAvg float64

	if rate.RateCount > 0{
		numberOfRates = rate.RateCount
		scoreAvg = score(rate.RateTotal, rate.RateCount)
	}

	return &pb.GetCourierRatingStatResponse{
		Score: &pb.RateScore{
			NumberOfRates: numberOfRates,
			ScoreAvg: scoreAvg,
		},
	}, nil

}

func (s *Service) GetClientRatingStat(ctx context.Context, in *pb.GetClientRatingStatRequest) (*pb.GetClientRatingStatResponse, error){
	logger.Infof("GetCourierRatingStat client_id = %v", in.GetClientId())

	rate := domain.Rate{}
	err := s.db.Model(&rate).Where("rated = ?", in.ClientId).Find(&rate).Error

	if err != nil {
		logger.Debugf("GetClientRatingStat stat error, client_id = %v, error = %v", in.GetClientId(), err)
		return nil, proto.Internal.Error(err)
	}

	var numberOfRates int64
	var scoreAvg float64

	if rate.RateCount > 0{
		numberOfRates = rate.RateCount
		scoreAvg = score(rate.RateTotal, rate.RateCount)
	}

	return &pb.GetClientRatingStatResponse{
		Score: &pb.RateScore{
			NumberOfRates: numberOfRates,
			ScoreAvg: scoreAvg,
		},
	}, nil
}

func (s *Service) GetCourierRatingStatByToken(ctx context.Context, in *pb.GetCourierRatingStatByTokenRequest) (*pb.GetCourierRatingStatByTokenResponse, error){
	logger.Infof("GetCourierRatingStatByToken")

	tokenUser := ctx.Value("user").(security.User)

	response,err := s.GetCourierRatingStat(ctx, &pb.GetCourierRatingStatRequest{
		CourierId: tokenUser.Id,
	})

	if err != nil{
		return nil, err
	}

	return &pb.GetCourierRatingStatByTokenResponse{
		CourierId: tokenUser.Id,
		Score: response.Score,
	}, nil
}

func (s *Service) GetClientRatingStatByToken(ctx context.Context, in *pb.GetClientRatingStatByTokenRequest) (*pb.GetClientRatingStatByTokenResponse, error){
	logger.Infof("GetClientRatingStatByToken")

	tokenUser := ctx.Value("user").(security.User)

	response,err := s.GetClientRatingStat(ctx, &pb.GetClientRatingStatRequest{
		ClientId: tokenUser.Id,
	})

	if err != nil{
		return nil, err
	}

	return &pb.GetClientRatingStatByTokenResponse{
		ClientId: tokenUser.Id,
		Score: response.Score,
	}, nil
}

func score(RateTotal int64, RateCount int64) float64{
	return math.Round((float64(RateTotal) / float64(RateCount)) * 100) / 100
}

type RideInfo struct {
	CourierId string
	ClientId string
}

func (s *Service) getRideDetails(rideId string, accessToken string) (*RideInfo,error) {
	request := &delivery.GetRequestsRequest{
		Filter: &delivery.GetRequestsRequest_Id{
			Id: rideId,
		},
	}

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	addr := fmt.Sprintf("%v:%v", s.config.RideService.Host, s.config.RideService.Port)

	opts = append(opts, grpc.WithBlock())

	conn, err := grpc.Dial(addr, opts...)

	if err != nil {
		return nil, err
	}

	defer conn.Close()

	clientDeadline := time.Now().Add(time.Duration(6000) * time.Millisecond)

	var header = metadata.Pairs("access_token", accessToken)

	ctx := metadata.NewOutgoingContext(context.Background(), header)

	ctx, cancel := context.WithDeadline(ctx, clientDeadline)

	defer cancel()

	c := delivery.NewDeliveryClient(conn)

	out, err := c.GetRequests(ctx, request)

	if err != nil{
		return nil, err
	}

	if out.GetRequests() == nil || len(out.GetRequests()) == 0 {
		return nil, proto.NotFound.ErrorMsg("ride not found")
	}

	r := out.GetRequests()[0]

	return &RideInfo{
		CourierId: r.CourierId,
		ClientId: r.CustomerId,
	}, nil
}

func NewService(config config.RatingData, jwtConfig config.JwtData) *Service {
	dbInstance,err := db.NewOrm(config.Database)
	if err != nil{
		logger.Fatalf("cannot connect to database", err)
	}
	jwtUtils, err := security.NewJWTUtils(jwtConfig)
	if err != nil{
		logger.Fatalf("cannot create jwtutils ", err)
	}
	return &Service{
		config: config,
		db: dbInstance,
		jwtUtils: &jwtUtils,
	}
}