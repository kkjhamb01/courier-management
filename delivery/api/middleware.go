package api

import (
	"context"
	"errors"

	"github.com/kkjhamb01/courier-management/common/config"
	"github.com/kkjhamb01/courier-management/common/logger"
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

type UserExtractor struct {
	jwtUtils *security.JWTUtils
}

func (interceptor *UserExtractor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		logger.Infof("unary API call: %v, request: %v", info.FullMethod, req)
		// only extract user, without authentication
		headers, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, errors.New("no metadata is supplied")
		}

		tokens := headers.Get("access_token")
		if len(tokens) == 0 {
			return nil, errors.New("access_token is not supplied")
		}

		user, err := interceptor.jwtUtils.ValidateUnsigned(tokens[0], true)
		if err != nil {
			return nil, err
		}
		if user == nil || user.Id == "" {
			return nil, errors.New("failed to get user id")
		}
		ctx = context.WithValue(ctx, "user", *user)
		return handler(ctx, req)
	}
}

func NewAuthInterceptor(config config.PartyData, jwtConfig config.JwtData) *UserExtractor {
	jwtUtils, err := security.NewJWTUtils(jwtConfig)
	if err != nil {
		logger.Fatalf("cannot create jwtutils ", err)
	}
	return &UserExtractor{&jwtUtils}
}
