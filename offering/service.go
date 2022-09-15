package offering

import (
	"context"

	"github.com/kkjhamb01/courier-management/common/http"
	"github.com/kkjhamb01/courier-management/common/service"
	"github.com/kkjhamb01/courier-management/offering/api"
	"github.com/kkjhamb01/courier-management/offering/db"
	"github.com/kkjhamb01/courier-management/offering/healthcheck"
	"github.com/kkjhamb01/courier-management/offering/maps"
	"github.com/kkjhamb01/courier-management/offering/messaging"
	"github.com/kkjhamb01/courier-management/offering/prometheus"
	"github.com/kkjhamb01/courier-management/offering/services"
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
