package api

import (
	"net"

	"github.com/kkjhamb01/courier-management/common/config"
	"github.com/kkjhamb01/courier-management/common/logger"
	offeringPb "github.com/kkjhamb01/courier-management/grpc/offering/go"
	"google.golang.org/grpc"
)

type serverImpl struct {
	// UnimplementedOfferingStreamServer must be embedded to have forward compatible implementations.
	offeringPb.UnimplementedOfferingServer
	grpcServer *grpc.Server
}

var server serverImpl

func CreateGrpcServer() {
	// create listener
	lis, err := net.Listen("tcp", ":"+config.Offering().GrpcPort)
	if err != nil {
		logger.Fatalf("failed to listen: %v", err)
	}

	// create grpc server
	authInterceptor := NewAuthInterceptor(config.Party(), config.Jwt())

	//tracingInterceptor := tracing.NewTracingInterceptor()

	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(ChainUnaryInterceptors(
			//tracingInterceptor.Unary(),
			authInterceptor.Unary(),
		)),
		grpc.StreamInterceptor(ChainStreamInterceptors(
			//tracingInterceptor.Stream(),
			authInterceptor.Stream(),
		)),
	}
	server.grpcServer = grpc.NewServer(opts...)
	offeringPb.RegisterOfferingServer(server.grpcServer, &server)

	logger.Infof("starting offering grpc server on localhost:%v ...", config.Offering().GrpcPort)

	// and start...
	if err := server.grpcServer.Serve(lis); err != nil {
		logger.Fatalf("failed to serve: %v", err)
	}
}

func StopGrpcServer() {
	server.grpcServer.Stop()
	server.grpcServer = nil
}
