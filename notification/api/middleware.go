package api

import (
	"context"

	"github.com/kkjhamb01/courier-management/common/config"
	"github.com/kkjhamb01/courier-management/common/logger"
	uaaproto "github.com/kkjhamb01/courier-management/uaa/proto"
	"github.com/kkjhamb01/courier-management/uaa/security"
	"google.golang.org/grpc"
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
			return nil, uaaproto.InvalidArgument.Error(err)
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
	jwtUtils *security.JWTUtils
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
		logger.Infof("--> unary interceptor: ", info.FullMethod)

		// only extract user, without authentication
		user := interceptor.extractUser(ctx, req)

		if user != nil && user.Id != "" {
			ctx = context.WithValue(context.Background(), "user", *user)
		} else {
			return nil, uaaproto.Unauthenticated.ErrorMsg("invalid token")
		}

		return handler(ctx, req)
	}
}

func (interceptor *AuthInterceptor) extractUser(ctx context.Context, req interface{}) *security.User {
	var user *security.User
	if a, ok := req.(authenticatedRequest); ok {
		token := a.GetAccessToken()
		user, _ = interceptor.jwtUtils.ValidateUnsigned(token, true)
	}
	logger.Debugf("extract user from request ", user)
	return user
}

func NewAuthInterceptor(config config.NotificationData, jwtConfig config.JwtData) *AuthInterceptor {
	jwtUtils, err := security.NewJWTUtils(jwtConfig)
	if err != nil {
		logger.Fatalf("cannot create jwtutils ", err)
	}
	return &AuthInterceptor{&jwtUtils}
}
