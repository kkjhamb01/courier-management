package api

import (
	"gitlab.artin.ai/backend/courier-management/common/config"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	deliveryPb "gitlab.artin.ai/backend/courier-management/grpc/delivery/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type serverImpl struct {
	// UnimplementedDeliveryServer must be embedded to have forward compatible implementations.
	deliveryPb.UnimplementedDeliveryServer
	grpcServer *grpc.Server
}

var server serverImpl

func CreateGrpcServer() {
	// create grpc listener
	port := config.Delivery().GrpcPort
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		logger.Fatalf("failed to listen: %v", err)
	}

	// create grpc server
	authInterceptor := NewAuthInterceptor(config.Party(), config.Jwt())

	//tracingInterceptor := tracing.NewTracingInterceptor()

	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(ChainInterceptors(
			//tracingInterceptor.Unary(),
			authInterceptor.Unary(),
		)),
	}
	server.grpcServer = grpc.NewServer(opts...)
	deliveryPb.RegisterDeliveryServer(server.grpcServer, &server)

	reflection.Register(server.grpcServer)

	logger.Infof("starting location grpc server on localhost:%v ...", port)

	// and start...
	if err := server.grpcServer.Serve(lis); err != nil {
		logger.Fatalf("failed to serve: %v", err)
	}
}

func StopGrpcServer() {
	server.grpcServer.Stop()
	server.grpcServer = nil
}
