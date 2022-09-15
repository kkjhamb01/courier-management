package api

import (
	"net"

	"github.com/kkjhamb01/courier-management/common/config"
	"github.com/kkjhamb01/courier-management/common/logger"
	"github.com/kkjhamb01/courier-management/party/business"
	pb "github.com/kkjhamb01/courier-management/party/proto"
	"github.com/kkjhamb01/courier-management/uaa/security"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	// UnimplementedAppCallbackServer must be embedded to have forward compatible implementations.
	pb.UnimplementedCourierAccountServiceServer
	pb.UnimplementedDocumentServiceServer
	pb.UnimplementedUaaServiceServer
	pb.UnimplementedUserAccountServiceServer
	pb.UnimplementedUserStatusServiceServer
	pb.UnimplementedPartyAdminServiceServer
	pb.UnimplementedUserStatusByTokenServiceServer
	pb.UnimplementedInterServiceServer
	server  *grpc.Server
	service *business.Service
}

var server = grpcServer{}

func CreateApiServer() {
	createGrpcServer()
}

func createGrpcServer() {
	// create listener
	lis, err := net.Listen("tcp", config.Party().Server.Address)
	if err != nil {
		logger.Fatalf("failed to listen: %v", err)
	}

	authInterceptor := NewAuthInterceptor(config.Party(), config.Jwt(), rolesMap())
	validationInterceptor := NewValidationInterceptor()
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
	server.service = business.NewService(config.GetData(), config.Jwt())

	reflection.Register(server.server)
	pb.RegisterCourierAccountServiceServer(server.server, &server)
	pb.RegisterDocumentServiceServer(server.server, &server)
	pb.RegisterUaaServiceServer(server.server, &server)
	pb.RegisterUserAccountServiceServer(server.server, &server)
	pb.RegisterUserStatusServiceServer(server.server, &server)
	pb.RegisterPartyAdminServiceServer(server.server, &server)
	pb.RegisterUserStatusByTokenServiceServer(server.server, &server)
	pb.RegisterInterServiceServer(server.server, &server)

	logger.Infof("starting party server on:%v ...", config.Party().Server.Address)

	// and start...
	if err := server.server.Serve(lis); err != nil {
		logger.Fatalf("failed to serve: %v", err)
	}
}

// this function creates a map to set which roles have access to which requests
func rolesMap() map[string][]string {
	const PREFIX_COURIER_ACCOUNT = "/party.CourierAccountService/"
	const PREFIX_DOCUMENT = "/party.DocumentService/"
	//const PREFIX_ADMIN = "/party.PartyAdminService/"
	const PREFIX_USER_ACCOUNT = "/party.UserAccountService/"
	const PREFIX_USER_STATUS_BY_TOKEN = "/party.UserStatusByTokenService/"

	var ADMIN = security.Role_name[int32(security.Role_ADMIN)]
	var COURIER = security.Role_name[int32(security.Role_COURIER)]
	var CLIENT = security.Role_name[int32(security.Role_CLIENT)]

	var RESOURCE_ACCESS_COURIER = []string{COURIER}
	var RESOURCE_ACCESS_CLIENT = []string{CLIENT}
	var RESOURCE_ACCESS_ALL = []string{ADMIN, COURIER, CLIENT}
	//var RESOURCE_ACCESS_ADMIN = []string{ADMIN}

	// nil values are open to everyone (even without token)
	return map[string][]string{
		PREFIX_COURIER_ACCOUNT + "CreateCourierAccount":              RESOURCE_ACCESS_COURIER,
		PREFIX_COURIER_ACCOUNT + "GetCourierAccount":                 RESOURCE_ACCESS_ALL,
		PREFIX_COURIER_ACCOUNT + "FindCourierAccounts":               RESOURCE_ACCESS_ALL,
		PREFIX_COURIER_ACCOUNT + "UpdateCourierAccount":              RESOURCE_ACCESS_COURIER,
		PREFIX_COURIER_ACCOUNT + "UpdateProfileAdditionalInfo":       RESOURCE_ACCESS_COURIER,
		PREFIX_COURIER_ACCOUNT + "DeleteProfileAdditionalInfo":       RESOURCE_ACCESS_COURIER,
		PREFIX_COURIER_ACCOUNT + "GetProfileAdditionalInfo":          RESOURCE_ACCESS_ALL,
		PREFIX_COURIER_ACCOUNT + "GetProfileAdditionalInfoStatus":    RESOURCE_ACCESS_ALL,
		PREFIX_COURIER_ACCOUNT + "UpdateProfileAdditionalInfoStatus": RESOURCE_ACCESS_COURIER,
		PREFIX_COURIER_ACCOUNT + "GetProfileStatus":                  RESOURCE_ACCESS_ALL,
		PREFIX_COURIER_ACCOUNT + "SearchMot":                         RESOURCE_ACCESS_ALL,
		PREFIX_COURIER_ACCOUNT + "UpdateCourierPhoneNumber":          RESOURCE_ACCESS_COURIER,

		PREFIX_DOCUMENT + "Upload":             RESOURCE_ACCESS_ALL,
		PREFIX_DOCUMENT + "GetDocumentsOfUser": RESOURCE_ACCESS_ALL,
		PREFIX_DOCUMENT + "GetDocument":        RESOURCE_ACCESS_ALL,
		PREFIX_DOCUMENT + "Download":           RESOURCE_ACCESS_ALL,
		PREFIX_DOCUMENT + "DirectDownload":     RESOURCE_ACCESS_ALL,

		PREFIX_USER_ACCOUNT + "CreateUserAccount":     RESOURCE_ACCESS_ALL,
		PREFIX_USER_ACCOUNT + "GetUserAccount":        RESOURCE_ACCESS_ALL,
		PREFIX_USER_ACCOUNT + "FindUserAccounts":      RESOURCE_ACCESS_ALL,
		PREFIX_USER_ACCOUNT + "UpdateUserAccount":     RESOURCE_ACCESS_ALL,
		PREFIX_USER_ACCOUNT + "UpdateUserCard":        RESOURCE_ACCESS_ALL,
		PREFIX_USER_ACCOUNT + "DeleteUserCard":        RESOURCE_ACCESS_ALL,
		PREFIX_USER_ACCOUNT + "GetUserCard":           RESOURCE_ACCESS_ALL,
		PREFIX_USER_ACCOUNT + "UpdateUserPhoneNumber": RESOURCE_ACCESS_ALL,
		PREFIX_USER_ACCOUNT + "GetUserAddress":        RESOURCE_ACCESS_ALL,
		PREFIX_USER_ACCOUNT + "UpdateUserAddress":     RESOURCE_ACCESS_ALL,

		PREFIX_USER_STATUS_BY_TOKEN + "GetCourierUserStatusByToken":    RESOURCE_ACCESS_COURIER,
		PREFIX_USER_STATUS_BY_TOKEN + "UpdateCourierUserStatusByToken": RESOURCE_ACCESS_COURIER,
		PREFIX_USER_STATUS_BY_TOKEN + "GetClientUserStatusByToken":     RESOURCE_ACCESS_CLIENT,
		PREFIX_USER_STATUS_BY_TOKEN + "UpdateClientUserStatusByToken":  RESOURCE_ACCESS_CLIENT,

		/*PREFIX_ADMIN + "ServiceGetCourierRegistrationStat": RESOURCE_ACCESS_ADMIN,
		PREFIX_ADMIN + "ServiceGetClientRegistrationStat":  RESOURCE_ACCESS_ADMIN,
		PREFIX_ADMIN + "ServiceGetCourierRegistration":     RESOURCE_ACCESS_ADMIN,
		PREFIX_ADMIN + "ServiceGetClientRegistration":      RESOURCE_ACCESS_ADMIN,
		PREFIX_ADMIN + "ServiceGetDocumentsOfUser":         RESOURCE_ACCESS_ADMIN,
		PREFIX_ADMIN + "ServiceGetProfileAdditionalInfo":   RESOURCE_ACCESS_ADMIN,
		PREFIX_ADMIN + "ServiceUpdateProfileStatus":        RESOURCE_ACCESS_ADMIN,*/
	}
}

func StopGrpcServer() {
	server.server.Stop()
	server.server = nil
}
