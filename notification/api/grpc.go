package api

import (
	"net"

	"github.com/kkjhamb01/courier-management/common/config"
	"github.com/kkjhamb01/courier-management/common/logger"
	"github.com/kkjhamb01/courier-management/notification/business"
	pb "github.com/kkjhamb01/courier-management/notification/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	// UnimplementedAppCallbackServer must be embedded to have forward compatible implementations.
	pb.UnimplementedNotificationServiceServer
	server  *grpc.Server
	service *business.Service
}

var server = grpcServer{}

func CreateApiServer() {
	createGrpcServer()
}

func createGrpcServer() {
	// create listener
	lis, err := net.Listen("tcp", config.Notification().Server.Address)
	if err != nil {
		logger.Fatalf("failed to listen: %v", err)
	}

	authInterceptor := NewAuthInterceptor(config.Notification(), config.Jwt())
	validationInterceptor := NewValidationInterceptor()

	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(ChainInterceptors(
			validationInterceptor.Unary(),
			authInterceptor.Unary(),
		)),
	}

	server.server = grpc.NewServer(opts...)
	server.service = business.NewService(config.Notification())

	reflection.Register(server.server)
	pb.RegisterNotificationServiceServer(server.server, &server)

	logger.Infof("starting notification registration server on:%v ...", config.Notification().Server.Address)

	// and start...
	if err := server.server.Serve(lis); err != nil {
		logger.Fatalf("failed to serve: %v", err)
	}
}

func StopGrpcServer() {
	server.server.Stop()
	server.server = nil
}
