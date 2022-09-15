package delivery

import (
	"context"

	"github.com/kkjhamb01/courier-management/common/http"
	"github.com/kkjhamb01/courier-management/common/service"
	"github.com/kkjhamb01/courier-management/delivery/api"
	"github.com/kkjhamb01/courier-management/delivery/db"
	"github.com/kkjhamb01/courier-management/delivery/healthcheck"
	"github.com/kkjhamb01/courier-management/delivery/maps"
	"github.com/kkjhamb01/courier-management/delivery/messaging"
	"github.com/kkjhamb01/courier-management/delivery/scheduler"
	"github.com/kkjhamb01/courier-management/delivery/services"
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
