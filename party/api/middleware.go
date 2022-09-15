package api

import (
	"context"

	"github.com/kkjhamb01/courier-management/common/config"
	"github.com/kkjhamb01/courier-management/common/logger"
	"github.com/kkjhamb01/courier-management/uaa/proto"
	"github.com/kkjhamb01/courier-management/uaa/security"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func ChainInterceptors(interceptors ...grpc.UnaryServerInterceptor) grpc.UnaryServerInterceptor {
	n := len(interceptors)

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		chainer := func(currentInter grpc.UnaryServerInterceptor, currentHandler grpc.UnaryHandler) grpc.UnaryHandler {
			return func(currentCtx context.Context, currentReq interface{}) (interface{}, error) {
				return currentInter(currentCtx, currentReq, info, currentHandler)
			}
		}

		chainedHandler := handler
		for i := n - 1; i >= 0; i-- {
			chainedHandler = chainer(interceptors[i], chainedHandler)
		}

		return chainedHandler(ctx, req)
	}
}

type ValidationInterceptor struct {
}

type Validator interface {
	Validate(all bool) error
}

func (v *ValidationInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (interface{}, error) {
		if err := v.validate(req); err != nil {
			return nil, proto.InvalidArgument.ErrorMsg(err.Error())
		}
		return handler(ctx, req)
	}
}

func (v *ValidationInterceptor) validate(candidate interface{}) error {
	if validator, ok := candidate.(Validator); ok {
		return validator.Validate(true)
	}
	return nil
}

func NewValidationInterceptor() *ValidationInterceptor {
	return &ValidationInterceptor{}
}

type AuthInterceptor struct {
	jwtUtils        *security.JWTUtils
	accessibleRoles map[string][]string
}

type authenticatedRequest interface {
	GetAccessToken() string
}

func (interceptor *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {

		// only extract user, without authentication
		user, err := interceptor.extractUser(ctx, req)

		if err != nil {
			return nil, err
		}

		err = interceptor.authorize(ctx, user, info.FullMethod)
		if err != nil {
			return nil, err
		}

		if user != nil && user.Id != "" {
			ctx = context.WithValue(context.Background(), "user", *user)
		}

		return handler(ctx, req)
	}
}

func (interceptor *AuthInterceptor) extractUser(ctx context.Context, req interface{}) (*security.User, error) {
	var user *security.User
	var token string
	var err error
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		authorization := md["access_token"]
		if len(authorization) > 0 {
			token = authorization[0]
		}
	}

	if token == "" {
		if a, ok := req.(authenticatedRequest); ok {
			token = a.GetAccessToken()
		}
	}

	if token != "" {
		user, err = interceptor.jwtUtils.ValidateUnsigned(token, true)
		if err != nil {
			return nil, err
		}
	}

	if user != nil {
		logger.Debugf("request has a valid token %v", user.Id)
	} else {
		logger.Debugf("request has an invalid token")
	}

	return user, nil
}

func (interceptor *AuthInterceptor) authorize(ctx context.Context, user *security.User, method string) error {
	accessibleRoles, ok := interceptor.accessibleRoles[method]
	if !ok || accessibleRoles == nil {
		//everyone can access
		return nil
	}
	if user == nil {
		// request is not authenticated
		return proto.Unauthenticated.ErrorMsg("user is not authenticated")
	}

	/*for _, resourceRole := range accessibleRoles {
		for _, userRole := range user.Roles {
			if resourceRole ==  security.Role_name[int32(userRole)] {
				return nil
			}
		}
	}
	return proto.PermissionDenied.ErrorMsg("You don't have enought permission to aceess this resource")*/

	return nil
}

func NewAuthInterceptor(config config.PartyData, jwtConfig config.JwtData, accessibleRoles map[string][]string) *AuthInterceptor {
	jwtUtils, err := security.NewJWTUtils(jwtConfig)
	if err != nil {
		logger.Fatalf("cannot create jwtutils ", err)
	}
	return &AuthInterceptor{&jwtUtils, accessibleRoles}
}
