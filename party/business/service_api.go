package business

import (
	"context"
	"gitlab.artin.ai/backend/courier-management/common/config"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	"gitlab.artin.ai/backend/courier-management/promotion/proto"
	"google.golang.org/grpc"
	"time"
)

type PromotionAPI struct{
	config    config.PromotionData
}

func (api PromotionAPI) AssignUserReferral(userId string, referral string, referredId string) error {
	logger.Debugf("AssignUserReferral userId %v referral %v", userId, referral)
	query := &proto.AssignUserReferralRequest{
		UserId: userId,
		Referral: referral,
		ReferredId: referredId,
	}
	conn := api.getConn()
	defer conn.Close()
	clientDeadline := time.Now().Add(time.Duration(6000) * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), clientDeadline)
	defer cancel()

	c := proto.NewPromotionServiceClient(conn)

	_, err := c.AssignUserReferral(ctx, query)
	return err
}

func (api PromotionAPI) getConn() *grpc.ClientConn{
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	opts = append(opts, grpc.WithBlock())
	conn, err := grpc.Dial(api.config.Server.Address, opts...)
	if err != nil {
		logger.Errorf("cannot connect to promotion: %v", err)
	}
	return conn
}

func NewPromotionAPI(config config.PromotionData) PromotionAPI{
	return PromotionAPI{
		config: config,
	}
}