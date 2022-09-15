package api

import (
	"net"

	"github.com/kkjhamb01/courier-management/common/config"
	"github.com/kkjhamb01/courier-management/common/logger"
	"github.com/kkjhamb01/courier-management/uaa/business"
	pb "github.com/kkjhamb01/courier-management/uaa/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	// UnimplementedAppCallbackServer must be embedded to have forward compatible implementations.
	pb.UnimplementedCourierRegisterServiceServer
	pb.UnimplementedTokenServiceServer
	pb.UnimplementedUserRegisterServiceServer
	pb.UnimplementedAdminServiceServer
	server              *grpc.Server
	registrationService *business.RegistrationService
	tokenService        *business.TokenService
	adminService        *business.AdminService
}

var server = grpcServer{}

func CreateApiServer() {
	createGrpcServer()
}

func createGrpcServer() {
	// create listener
	lis, err := net.Listen("tcp", config.Uaa().Server.Address)
	if err != nil {
		logger.Fatalf("failed to listen: %v", err)
	}

	validationInterceptor := NewValidationInterceptor()
	authInterceptor := NewAuthInterceptor(config.Jwt())
	//tracingInterceptor := tracing.NewTracingInterceptor()

	opts := []grpc.ServerOption{
		// chain multiple interceptors
		grpc.UnaryInterceptor(ChainInterceptors(
			//tracingInterceptor.Unary(),
			validationInterceptor.Unary(),
			authInterceptor.Unary(),
		)),
		//grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
	}

	// create grpc server
	server.server = grpc.NewServer(opts...)
	partyAPI := business.NewPartyAPI(config.Uaa())
	server.registrationService = business.NewRegistrationService(config.Uaa(), config.Jwt(), partyAPI)
	server.tokenService = business.NewTokenService(config.Uaa(), config.Jwt(), partyAPI)
	server.adminService = business.NewAdminService(config.Uaa(), config.Jwt(), partyAPI)

	reflection.Register(server.server)
	pb.RegisterCourierRegisterServiceServer(server.server, &server)
	pb.RegisterTokenServiceServer(server.server, &server)
	pb.RegisterUserRegisterServiceServer(server.server, &server)
	pb.RegisterAdminServiceServer(server.server, &server)

	logger.Infof("starting uaa server on:%v ...", config.Uaa().Server.Address)

	// and start...
	if err := server.server.Serve(lis); err != nil {
		logger.Fatalf("failed to serve: %v", err)
	}
}

func StopGrpcServer() {
	server.server.Stop()
	server.server = nil
}
