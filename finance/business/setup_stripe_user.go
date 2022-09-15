package business

import (
	"context"
	"errors"
	"fmt"
	"gitlab.artin.ai/backend/courier-management/common/config"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	commonPb "gitlab.artin.ai/backend/courier-management/grpc/common/go"
	financePb "gitlab.artin.ai/backend/courier-management/grpc/finance/go"
)

var (
	CreateStripeCustomerMaxRetriesErr = errors.New("create stripe customer/account exceeded the max limit")
)

func SetUpStripeUser(ctx context.Context, userType commonPb.UserType, userId string, accessToken string, createRetries int) (financePb.SetUpStripeUserResponse, error) {

	logger.Infof("Business/SetUpStripeUser userType = %v, userId = %v, createRetries = %v", userType, userId, createRetries)

	if createRetries > config.Finance().CreateStripeAccountMaxRetries {
		return financePb.SetUpStripeUserResponse{}, CreateStripeCustomerMaxRetriesErr
	}

	var email, name, phone, firstName, lastName string
	var isCourier bool
	switch userType {
	case commonPb.UserType_CUSTOMER:
		isCourier = false
		userProfile, err := getCustomerInfo(ctx, accessToken)
		logger.Infof("Business/SetUpStripeUser userProfile = %+v", userProfile)
		if err != nil {
			logger.Errorf("Business/SetUpStripeUser failed to fetch customer info %v", err)
			return financePb.SetUpStripeUserResponse{}, err
		}
		email = userProfile.Email
		name = userProfile.FirstName + " " + userProfile.LastName
		firstName = userProfile.FirstName
		lastName = userProfile.LastName
		phone = userProfile.PhoneNumber
	case commonPb.UserType_COURIER:
		isCourier = true
		courierProfile, err := getCourierInfo(ctx, accessToken)
		logger.Infof("Business/SetUpStripeUser courierProfile = %+v", courierProfile)
		if err != nil {
			logger.Errorf("Business/SetUpStripeUser failed to fetch courier info %v", err)
			return financePb.SetUpStripeUserResponse{}, err
		}
		email = courierProfile.Email
		name = courierProfile.FirstName + " " + courierProfile.LastName
		phone = courierProfile.PhoneNumber
	default:
		err := fmt.Errorf("the user type: %v is not known", userType)
		logger.Errorf("Business/SetUpStripeUser failed to fetch email, name and phone number of the user %v", err)
		return financePb.SetUpStripeUserResponse{}, err
	}

	var response financePb.SetUpStripeUserResponse
	if createRetries == 0 {
		err := CreateStripeCustomer(ctx, userId, email, name, phone)
		if err != nil {
			logger.Errorf("Business/SetUpStripeUser failed to create stripe customer %v", err)
			return financePb.SetUpStripeUserResponse{}, err
		}
	}

	if isCourier {
		onboardingUrl, err := CreateStripeAccount(ctx, userId, accessToken, firstName, lastName, createRetries == 0)
		logger.Infof("Business/SetUpStripeUser onboardingUrl = %+v", onboardingUrl)
		if err != nil {
			logger.Errorf("Business/SetUpStripeUser failed to create stripe account %v", err)
			return response, err
		}
		response.Event = &financePb.SetUpStripeUserResponse_OnboardingUrl{
			OnboardingUrl: &onboardingUrl,
		}
	} else {
		err := CreateAccount(ctx, userId, accessToken)
		if err != nil {
			logger.Errorf("Business/SetUpStripeUser failed to create customer account %v", err)
			return response, err
		}

		response.Event = &financePb.SetUpStripeUserResponse_OnboardingResult{
			OnboardingResult: &financePb.OnboardingResult{
				Successful: true,
				Desc:       "Stripe customer has been created for the customer",
			},
		}
	}

	return response, nil
}
