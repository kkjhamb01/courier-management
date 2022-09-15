package api

import (
	"context"
	"errors"
	"io"

	"github.com/kkjhamb01/courier-management/common/logger"
	"github.com/kkjhamb01/courier-management/common/logger/tag"
	"github.com/kkjhamb01/courier-management/finance/business"
	"github.com/kkjhamb01/courier-management/finance/pubsub"
	financePb "github.com/kkjhamb01/courier-management/grpc/finance/go"
	"github.com/kkjhamb01/courier-management/uaa/security"
)

func (s serverImpl) SetUpStripeUser(stream financePb.Finance_SetUpStripeUserServer) error {
	for createRetries := 0; ; createRetries++ {
		req, err := stream.Recv()

		logger.Infof("SetUpStripeUser createRetries = %v", createRetries)

		if err == io.EOF {
			logger.Info("SetUpStripeUser CreateCustomer done")
			return nil
		}

		ctx := stream.Context()

		if err = req.Validate(); err != nil {
			logger.Errorf("SetUpStripeUser request validation failed %v", err)
			return err
		}

		if ctx.Value(CtxKeyUser) == nil {
			err := errors.New("SetUpStripeUser no user found in context")
			return err
		}

		if ctx.Value(CtxKeyAccessToken) == nil {
			err := errors.New("SetUpStripeUser no access token found in context")
			return err
		}

		tokenUser := ctx.Value(CtxKeyUser).(security.User)
		accessToken := ctx.Value(CtxKeyAccessToken).(string)

		// start subscribing only for the first time
		if createRetries == 0 {
			newOfferCh := pubsub.SubscribeById(pubsub.TopicOnboardingResult, tokenUser.Id)
			go listenToOnboardingEvent(ctx, pubsub.TopicOnboardingResult, newOfferCh, tokenUser.Id, stream, onboardingResultEventHandler)
		}

		setupResponse, err := business.SetUpStripeUser(ctx, req.UserType, tokenUser.Id, accessToken, createRetries)
		if err != nil {
			logger.Error("SetUpStripeUser failed to set up stripe user", err)
			return err
		}

		logger.Info("SetUpStripeUser successfully called", tag.Obj("request", req))
		err = stream.Send(&setupResponse)
		if err != nil {
			logger.Error("SetUpStripeUser failed to send the response to client", err)
			return err
		}
	}
}

func onboardingResultEventHandler(event pubsub.Event, stream financePb.Finance_SetUpStripeUserServer) {

	// get the event type
	onboardingResult := event.OnboardingResultData()

	logger.Infof("onboardingResultEventHandler event = %+v", onboardingResult)

	// create the response
	response := &financePb.SetUpStripeUserResponse{
		Event: &financePb.SetUpStripeUserResponse_OnboardingResult{
			OnboardingResult: &onboardingResult,
		},
	}

	//send the stream
	err := stream.Send(response)
	if err != nil {
		// ignore error to tolerate failure
		logger.Errorf("onboardingResultEventHandler failed to send event %v, %+v", err, event)
	}
}

func listenToOnboardingEvent(ctx context.Context, topic string, eventCha chan pubsub.Event,
	courierId string, stream financePb.Finance_SetUpStripeUserServer,
	eventHandler func(pubsub.Event, financePb.Finance_SetUpStripeUserServer)) {

	logger.Infof("listenToOnboardingEvent topic = %v, courierId = %v", topic, courierId)

	for {
		select {
		case event := <-eventCha:
			eventHandler(event, stream)
		case <-ctx.Done():
			pubsub.UnsubscribeById(topic, courierId)
			return
		}
	}
}
