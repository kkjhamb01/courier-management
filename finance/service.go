package finance

import (
	"context"
	"gitlab.artin.ai/backend/courier-management/common/service"
	"gitlab.artin.ai/backend/courier-management/finance/api"
	"gitlab.artin.ai/backend/courier-management/finance/db"
	"gitlab.artin.ai/backend/courier-management/finance/scheduler"
	"gitlab.artin.ai/backend/courier-management/finance/services"
	"gitlab.artin.ai/backend/courier-management/finance/stripe"
)

func init() {
	service.RegisterService(financeService{})
}

type financeService struct {
}

func (s financeService) Start(ctx context.Context) {
	// TODO start internal http server (add prometheus and health-check endpoints)

	// setup finance service
	db.SetupRedisClient()
	// setup stripe
	stripe.Setup()
	go stripe.StartWebhook()
	// setup mysql
	db.SetupMariaDbClient()
	db.MigrateGormIntoMariaDb()
	err := db.InitSeedData()
	if err != nil {
		panic(err)
	}
	scheduler.StartCreatePayment()
	// start finance gRPC server
	api.CreateGrpcServer()
}

func (s financeService) Stop() error {
	api.StopGrpcServer()
	services.CloseAllConnections()
	stripe.StopWebhook()
	return nil
}

func (s financeService) Name() string {
	return service.Finance
}

func (s financeService) RelDir() string {
	return service.Finance
}
