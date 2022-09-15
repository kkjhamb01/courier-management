package business

import (
	"context"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/account"
	"github.com/stripe/stripe-go/v72/accountlink"
	"gitlab.artin.ai/backend/courier-management/common/config"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	"gitlab.artin.ai/backend/courier-management/finance/storage"
	financePb "gitlab.artin.ai/backend/courier-management/grpc/finance/go"
	"google.golang.org/protobuf/types/known/timestamppb"
	"os"
	"time"
)

func CreateStripeAccount(ctx context.Context, userId, accessToken, firstName, lastName string, firstTime bool) (financePb.OnboardingUrl, error) {

	logger.Infof("CreateStripeAccount userId = %v, firstName = %v, lastName = %v, firstTime = %v", userId, firstName, lastName, firstTime)

	var accId string
	if firstTime {
		err := CreateAccount(ctx, userId, accessToken)
		if err != nil {
			logger.Errorf("failed to create an account for the user %v", err, userId)
			return financePb.OnboardingUrl{}, err
		}

		accParams := &stripe.AccountParams{
			Type:    stripe.String(string(stripe.AccountTypeExpress)),
			Country: stripe.String("US"),
			Individual: &stripe.PersonParams{
				FirstName: stripe.String(firstName),
				LastName:  stripe.String(lastName),
			},
			Capabilities: &stripe.AccountCapabilitiesParams{
				Transfers: &stripe.AccountCapabilitiesTransfersParams{
					Requested: stripe.Bool(true),
				},
			},
			BusinessType: stripe.String("individual"),
			BusinessProfile: &stripe.AccountBusinessProfileParams{
				Name: stripe.String("Artin"),
				URL:  stripe.String("artin.co.uk"),
			},
		}
		acc, err := account.New(accParams)
		if err != nil {
			logger.Error("failed to create account", err)
			return financePb.OnboardingUrl{}, err
		}
		accId = acc.ID
		logger.Info("stripe id " + acc.ID)

		err = storage.SetStripeAccId(ctx, userId, accId)
		if err != nil {
			logger.Error("failed to save stripe account id in storage", err)
			return financePb.OnboardingUrl{}, err
		}

		err = storage.SetAccUserId(ctx, accId, userId)
		if err != nil {
			logger.Error("failed to save stripe account id in storage", err)
			return financePb.OnboardingUrl{}, err
		}
	} else {
		var err error
		accId, err = storage.GetStripeAccId(ctx, userId)
		if err != nil {
			logger.Error("failed to get stripe account id from storage", err)
			return financePb.OnboardingUrl{}, err
		}
	}

	financeHost := os.Getenv("FINANCE_SERVICE_SERVICE_HOST")
	financePort := config.Finance().WebhookPort
	financeHttpAddress := financeHost + ":" + financePort

	refreshUrl := "https://" + financeHttpAddress + "/refresh"
	returnUrl := "https://" + financeHttpAddress + "/return"

	params := &stripe.AccountLinkParams{
		Account:    stripe.String(accId),
		RefreshURL: stripe.String(refreshUrl),
		ReturnURL:  stripe.String(returnUrl),
		Type:       stripe.String("account_onboarding"),
	}
	accRef, err := accountlink.New(params)
	if err != nil {
		logger.Error("failed to create account onboarding page", err)
		return financePb.OnboardingUrl{}, err
	}

	return financePb.OnboardingUrl{
		Url:        accRef.URL,
		CreatedAt:  timestamppb.New(time.Unix(0, accRef.Created)),
		ExpiresAt:  timestamppb.New(time.Unix(0, accRef.ExpiresAt)),
		RefreshUrl: refreshUrl,
		ReturnUrl:  returnUrl,
	}, nil
}
