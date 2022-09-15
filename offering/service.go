package offering

import (
	"context"
	"gitlab.artin.ai/backend/courier-management/common/http"
	"gitlab.artin.ai/backend/courier-management/common/service"
	"gitlab.artin.ai/backend/courier-management/offering/api"
	"gitlab.artin.ai/backend/courier-management/offering/db"
	"gitlab.artin.ai/backend/courier-management/offering/healthcheck"
	"gitlab.artin.ai/backend/courier-management/offering/maps"
	"gitlab.artin.ai/backend/courier-management/offering/messaging"
	"gitlab.artin.ai/backend/courier-management/offering/prometheus"
	"gitlab.artin.ai/backend/courier-management/offering/services"
)

func init() {
	service.RegisterService(offeringService{})
}

type offeringService struct {
}

func (s offeringService) Start(ctx context.Context) {
	// setup redis client
	db.SetupRedisClient()
	// setup tile38 client
	db.SetupTile38Client()
	db.SetupMariaDbClient()
	db.MigrateGormIntoMariaDb()
	maps.SetupGoogleClient()

	// set up message queue subscription
	messaging.StartSubscriptions()
	// start internal http server (add prometheus and health-check endpoints)
	go http.CreateServer().
		AddPrometheus(ctx, prometheus.Setup).
		AddHealthCheck(healthcheck.Readyz).
		ExposeEndpoint()
	// start offering gRPC server
	api.CreateGrpcServer()
}

func (s offeringService) Stop() error {
	messaging.StopSubscriptions()
	api.StopGrpcServer()
	services.CloseAllConnections()
	return nil
}

func (s offeringService) Name() string {
	return service.Offering
}

func (s offeringService) RelDir() string {
	return service.Offering
}
