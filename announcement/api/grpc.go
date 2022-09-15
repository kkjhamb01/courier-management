package api

import (
	"net"

	"github.com/kkjhamb01/courier-management/announcement/business"
	pb "github.com/kkjhamb01/courier-management/announcement/proto"
	"github.com/kkjhamb01/courier-management/common/config"
	"github.com/kkjhamb01/courier-management/common/logger"
	"github.com/kkjhamb01/courier-management/uaa/security"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	// UnimplementedAppCallbackServer must be embedded to have forward compatible implementations.
	pb.UnimplementedAnnouncementUserServiceServer
	pb.UnimplementedAnnouncementAdminServiceServer
	server  *grpc.Server
	service *business.Service
}

var server = grpcServer{}

func CreateApiServer() {
	createGrpcServer()
}

func createGrpcServer() {
	// create listener
	lis, err := net.Listen("tcp", config.Announcement().Server.Address)
	if err != nil {
		logger.Fatalf("failed to listen: %v", err)
	}

	authInterceptor := NewAuthInterceptor(config.Announcement(), config.Jwt(), rolesMap())
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
	server.service = business.NewService(config.GetData())

	reflection.Register(server.server)
	pb.RegisterAnnouncementUserServiceServer(server.server, &server)
	pb.RegisterAnnouncementAdminServiceServer(server.server, &server)

	logger.Infof("starting announcement server on:%v ...", config.Announcement().Server.Address)

	// and start...
	if err := server.server.Serve(lis); err != nil {
		logger.Fatalf("failed to serve: %v", err)
	}
}

// this function creates a map to set which roles have access to which requests
func rolesMap() map[string][]string {
	const PREFIX_ADMIN_SERVICE = "/announcement.AnnouncementAdminService/"
	const PREFIX_USER_SERVICE = "/announcement.AnnouncementUserService/"

	var ADMIN = security.Role_name[int32(security.Role_ADMIN)]
	var CLIENT = security.Role_name[int32(security.Role_CLIENT)]
	var COURIER = security.Role_name[int32(security.Role_COURIER)]

	var ResourceAccessUser = []string{CLIENT, COURIER}
	var ResourceAccessAdmin = []string{ADMIN}

	// nil values are open to everyone (even without token)
	return map[string][]string{
		PREFIX_ADMIN_SERVICE + "CreateAnnouncement":       ResourceAccessAdmin,
		PREFIX_ADMIN_SERVICE + "AssignAnnouncementToUser": ResourceAccessAdmin,
		PREFIX_ADMIN_SERVICE + "GetAnnouncements":         ResourceAccessAdmin,
		PREFIX_ADMIN_SERVICE + "GetUsers":                 ResourceAccessAdmin,
		PREFIX_USER_SERVICE + "GetAnnouncementsOfUser":    ResourceAccessUser,
	}
}

func StopGrpcServer() {
	server.server.Stop()
	server.server = nil
}
