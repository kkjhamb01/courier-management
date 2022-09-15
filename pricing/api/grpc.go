package api

import (
	"gitlab.artin.ai/backend/courier-management/common/config"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	pricingPb "gitlab.artin.ai/backend/courier-management/grpc/pricing/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type serverImpl struct {
	// UnimplementedPricingServer must be embedded to have forward compatible implementations.
	pricingPb.UnimplementedPricingServer
	grpcServer *grpc.Server
}

var server serverImpl

func CreateGrpcServer() {
	// create listener
	addr := ":" + config.Pricing().GrpcPort
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		logger.Fatalf("failed to listen: %v", err)
	}

	// create grpc server
	var opt []grpc.ServerOption
	server.grpcServer = grpc.NewServer(opt...)
	pricingPb.RegisterPricingServer(server.grpcServer, &server)

	reflection.Register(server.grpcServer)

	logger.Infof("starting pricing grpc server on localhost:%v ...", addr)

	// and start...
	if err := server.grpcServer.Serve(lis); err != nil {
		logger.Fatalf("failed to serve: %v", err)
	}
}

func StopGrpcServer() {
	server.grpcServer.Stop()
	server.grpcServer = nil
}
