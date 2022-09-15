package api

import (
	"net"

	"github.com/kkjhamb01/courier-management/common/config"
	"github.com/kkjhamb01/courier-management/common/logger"
	"github.com/kkjhamb01/courier-management/rating/business"
	pb "github.com/kkjhamb01/courier-management/rating/proto"
	"github.com/kkjhamb01/courier-management/uaa/security"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	// UnimplementedAppCallbackServer must be embedded to have forward compatible implementations.
	pb.UnimplementedRatingServiceServer
	server  *grpc.Server
	service *business.Service
}

var server = grpcServer{}

func CreateApiServer() {
	createGrpcServer()
}

func createGrpcServer() {
	// create listener
	lis, err := net.Listen("tcp", config.Rating().Server.Address)
	if err != nil {
		logger.Fatalf("failed to listen: %v", err)
	}

	authInterceptor := NewAuthInterceptor(config.Rating(), config.Jwt(), rolesMap())
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
	server.service = business.NewService(config.Rating(), config.Jwt())

	reflection.Register(server.server)
	pb.RegisterRatingServiceServer(server.server, &server)

	logger.Infof("starting rating server on:%v ...", config.Rating().Server.Address)

	// and start...
	if err := server.server.Serve(lis); err != nil {
		logger.Fatalf("failed to serve: %v", err)
	}
}

// this function creates a map to set which roles have access to which requests
func rolesMap() map[string][]string {
	const PREFIX_RATING_SERVICE = "/rating.RatingService/"

	var ADMIN = security.Role_name[int32(security.Role_ADMIN)]
	var COURIER = security.Role_name[int32(security.Role_COURIER)]
	var CLIENT = security.Role_name[int32(security.Role_CLIENT)]

	var ResourceAccessCourier = []string{COURIER}
	var ResourceAccessClient = []string{CLIENT}
	var ResourceAccessAll = []string{ADMIN, COURIER, CLIENT}
	//var ResourceAccessAdmin = []string{ADMIN}

	// nil values are open to everyone (even without token)
	return map[string][]string{
		PREFIX_RATING_SERVICE + "CreateCourierRating":  ResourceAccessClient,
		PREFIX_RATING_SERVICE + "CreateClientRating":   ResourceAccessCourier,
		PREFIX_RATING_SERVICE + "GetCourierRating":     ResourceAccessAll,
		PREFIX_RATING_SERVICE + "GetClientRating":      ResourceAccessAll,
		PREFIX_RATING_SERVICE + "GetCourierRatingStat": nil,
		PREFIX_RATING_SERVICE + "GetClientRatingStat":  nil,
	}
}

func StopGrpcServer() {
	server.server.Stop()
	server.server = nil
}
