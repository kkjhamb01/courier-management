package delivery

import (
	"context"
	"gitlab.artin.ai/backend/courier-management/common/http"
	"gitlab.artin.ai/backend/courier-management/common/service"
	"gitlab.artin.ai/backend/courier-management/delivery/api"
	"gitlab.artin.ai/backend/courier-management/delivery/db"
	"gitlab.artin.ai/backend/courier-management/delivery/healthcheck"
	"gitlab.artin.ai/backend/courier-management/delivery/maps"
	"gitlab.artin.ai/backend/courier-management/delivery/messaging"
	"gitlab.artin.ai/backend/courier-management/delivery/scheduler"
	"gitlab.artin.ai/backend/courier-management/delivery/services"
)

func init() {
	service.RegisterService(deliveryService{})
}

type deliveryService struct {
}

func (s deliveryService) Start(ctx context.Context) {
	db.SetupMariaDbClient()
	db.MigrateGormIntoMariaDb()
	db.SetupRedisClient()
	messaging.StartSubscriptions()

	maps.SetupGoogleClient()
	scheduler.StartCreateRequest()
	// start internal http server (add prometheus and health-check endpoints)
	go http.CreateServer().
		// TODO add location service prometheus
		//AddPrometheus(ctx, prometheus.Setup).
		AddHealthCheck(healthcheck.Readyz).
		ExposeEndpoint()
	// start location gRPC server
	api.CreateGrpcServer()
}

func (s deliveryService) Stop() error {
	//messaging.StopSubscriptions()
	api.StopGrpcServer()
	services.CloseAllConnections()
	messaging.StopSubscriptions()
	scheduler.StopCreateRequest()
	return nil
}

func (s deliveryService) Name() string {
	return service.Delivery
}

func (s deliveryService) RelDir() string {
	return service.Delivery
}
