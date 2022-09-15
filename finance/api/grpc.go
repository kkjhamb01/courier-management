package api

import (
	"gitlab.artin.ai/backend/courier-management/common/config"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	financePb "gitlab.artin.ai/backend/courier-management/grpc/finance/go"
	"google.golang.org/grpc"
	"net"
)

type serverImpl struct {
	// UnimplementedFinanceServer must be embedded to have forward compatible implementations.
	financePb.UnimplementedFinanceServer
	grpcServer *grpc.Server
}

var server serverImpl

func CreateGrpcServer() {
	// create listener
	port := ":" + config.Finance().GrpcPort
	lis, err := net.Listen("tcp", port)
	if err != nil {
		logger.Fatalf("failed to listen: %v", err)
	}

	// create grpc server
	authInterceptor := NewAuthInterceptor(config.Party(), config.Jwt())

	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(ChainInterceptors(
			authInterceptor.Unary(),
		)),
		grpc.StreamInterceptor(ChainStreamInterceptors(
			authInterceptor.Stream(),
		)),
	}
	server.grpcServer = grpc.NewServer(opts...)
	financePb.RegisterFinanceServer(server.grpcServer, &server)

	logger.Infof("starting finance grpc server on localhost:%v ...", port)

	// and start...
	if err := server.grpcServer.Serve(lis); err != nil {
		logger.Fatalf("failed to serve: %v", err)
	}
}

func StopGrpcServer() {
	server.grpcServer.Stop()
	server.grpcServer = nil
}
