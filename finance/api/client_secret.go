package api

import (
	"context"
	"errors"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	"gitlab.artin.ai/backend/courier-management/common/logger/tag"
	"gitlab.artin.ai/backend/courier-management/finance/business"
	financePb "gitlab.artin.ai/backend/courier-management/grpc/finance/go"
	"gitlab.artin.ai/backend/courier-management/uaa/security"
)

func (s serverImpl) GetClientSecret(ctx context.Context, req *financePb.GetClientSecretRequest) (*financePb.GetClientSecretResponse, error) {

	logger.Infof("GetClientSecret request = %+v", req)

	if err := req.Validate(); err != nil {
		logger.Error("GetClientSecret request validation failed", err, tag.Obj("request", req))
		return nil, err
	}

	if ctx.Value(CtxKeyUser) == nil {
		err := errors.New("no user found in context")
		return nil, err
	}

	tokenUser := ctx.Value(CtxKeyUser).(security.User)

	ClientSecret, err := business.GetClientSecret(ctx, tokenUser.Id)
	if err != nil {
		logger.Error("failed to get client secret", err)
		return nil, err
	}

	logger.Info("GetClientSecret successfully called", tag.Obj("request", req))
	return &financePb.GetClientSecretResponse{
		ClientSecret: ClientSecret,
	}, nil
}
