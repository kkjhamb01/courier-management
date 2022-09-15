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

func ChainUnaryInterceptors(interceptors ...grpc.UnaryServerInterceptor) grpc.UnaryServerInterceptor {
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

func ChainStreamInterceptors(interceptors ...grpc.StreamServerInterceptor) grpc.StreamServerInterceptor {
	n := len(interceptors)

	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		chainer := func(currentInter grpc.StreamServerInterceptor, currentHandler grpc.StreamHandler) grpc.StreamHandler {
			return func(srv interface{}, stream grpc.ServerStream) error {
				return currentInter(srv, stream, info, currentHandler)
			}
		}

		chainedHandler := handler
		for i := n - 1; i >= 0; i-- {
			chainedHandler = chainer(interceptors[i], chainedHandler)
		}

		return chainedHandler(srv, stream)
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
		ctx = context.WithValue(ctx, "access_token", tokens[0])

		user, err := interceptor.jwtUtils.ValidateUnsigned(tokens[0], true)
		if err != nil {
			return nil, err
		}
		if user == nil || user.Id == "" {
			return nil, errors.New("failed to get user id")
		}
		ctx = context.WithValue(ctx, "user", *user)
		logger.Infof("user %v", user.Id)
		return handler(ctx, req)
	}
}

type ServerStreamWrapper struct {
	accessToken string
	user        security.User
	grpc.ServerStream
}

func (s ServerStreamWrapper) Context() context.Context {
	ctx := s.ServerStream.Context()
	ctx = context.WithValue(ctx, "access_token", s.accessToken)
	ctx = context.WithValue(ctx, "user", s.user)
	return ctx
}

func (interceptor *UserExtractor) Stream() grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream,
		info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		logger.Infof("--> stream interceptor: %v", info.FullMethod)

		streamWrapper := ServerStreamWrapper{
			ServerStream: stream,
		}

		// only extract user, without authentication
		headers, ok := metadata.FromIncomingContext(stream.Context())
		if !ok {
			return errors.New("no metadata is supplied")
		}

		tokens := headers.Get("access_token")
		if len(tokens) == 0 {
			return errors.New("access_token is not supplied")
		}
		streamWrapper.accessToken = tokens[0]

		user, err := interceptor.jwtUtils.ValidateUnsigned(tokens[0], true)
		if err != nil {
			return err
		}
		if user == nil || user.Id == "" {
			return errors.New("failed to get user id")
		}
		streamWrapper.user = *user

		return handler(srv, streamWrapper)
	}
}

func NewAuthInterceptor(config config.PartyData, jwtConfig config.JwtData) *UserExtractor {
	jwtUtils, err := security.NewJWTUtils(jwtConfig)
	if err != nil {
		logger.Fatalf("cannot create jwtutils ", err)
	}
	return &UserExtractor{&jwtUtils}
}
