package api

import (
	"net"

	"github.com/kkjhamb01/courier-management/common/config"
	"github.com/kkjhamb01/courier-management/common/logger"
	"github.com/kkjhamb01/courier-management/promotion/business"
	pb "github.com/kkjhamb01/courier-management/promotion/proto"
	"github.com/kkjhamb01/courier-management/uaa/security"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	// UnimplementedAppCallbackServer must be embedded to have forward compatible implementations.
	pb.UnimplementedPromotionServiceServer
	pb.UnimplementedPromotionAdminServiceServer
	pb.UnimplementedPromotionUserServiceServer
	server  *grpc.Server
	service *business.Service
}

var server = grpcServer{}

func CreateApiServer() {
	createGrpcServer()
}

func createGrpcServer() {
	// create listener
	lis, err := net.Listen("tcp", config.Promotion().Server.Address)
	if err != nil {
		logger.Fatalf("failed to listen: %v", err)
	}

	authInterceptor := NewAuthInterceptor(config.Promotion(), config.Jwt(), rolesMap())
	validationInterceptor := NewValidationInterceptor()

	opts := []grpc.ServerOption{
		// chain multiple interceptors
		grpc.UnaryInterceptor(ChainInterceptors(
			validationInterceptor.Unary(),
			authInterceptor.Unary(),
		)),
		//grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
	}

	// create grpc server
	server.server = grpc.NewServer(opts...)
	server.service = business.NewService(config.Promotion(), config.Jwt())

	reflection.Register(server.server)
	pb.RegisterPromotionServiceServer(server.server, &server)
	pb.RegisterPromotionAdminServiceServer(server.server, &server)
	pb.RegisterPromotionUserServiceServer(server.server, &server)

	logger.Infof("starting promotion server on:%v ...", config.Promotion().Server.Address)

	// and start...
	if err := server.server.Serve(lis); err != nil {
		logger.Fatalf("failed to serve: %v", err)
	}
}

// this function creates a map to set which roles have access to which requests
func rolesMap() map[string][]string {
	const PREFIX_ADMIN_SERVICE = "/promotion.PromotionAdminService/"
	const PREFIX_SERVICE_SERVICE = "/promotion.PromotionService/"
	const PREFIX_USER_SERVICE = "/promotion.PromotionUserService/"

	var ADMIN = security.Role_name[int32(security.Role_ADMIN)]
	var CLIENT = security.Role_name[int32(security.Role_CLIENT)]

	var ResourceAccessClient = []string{CLIENT}
	var ResourceAccessAdmin = []string{ADMIN}

	// nil values are open to everyone (even without token)
	return map[string][]string{
		PREFIX_ADMIN_SERVICE + "CreatePromotion":       ResourceAccessAdmin,
		PREFIX_ADMIN_SERVICE + "AssignPromotionToUser": ResourceAccessAdmin,
		PREFIX_ADMIN_SERVICE + "GetPromotions":         ResourceAccessAdmin,
		PREFIX_ADMIN_SERVICE + "GetUsers":              ResourceAccessAdmin,
		PREFIX_SERVICE_SERVICE + "AssignUserReferral":  nil,
		PREFIX_SERVICE_SERVICE + "ApplyPromotion":      nil,
		PREFIX_USER_SERVICE + "GetPromotionsOfUser":    ResourceAccessClient,
	}
}

func StopGrpcServer() {
	server.server.Stop()
	server.server = nil
}
