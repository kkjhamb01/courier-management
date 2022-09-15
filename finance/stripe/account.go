package stripe

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/stripe/stripe-go/v72"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	"gitlab.artin.ai/backend/courier-management/delivery/services"
	"gitlab.artin.ai/backend/courier-management/finance/pubsub"
	"gitlab.artin.ai/backend/courier-management/finance/storage"
	financePb "gitlab.artin.ai/backend/courier-management/grpc/finance/go"
	"gitlab.artin.ai/backend/courier-management/party/proto"
)

var (
	NoExternalAccount = errors.New("no external account is associated")
)

func onAccountUpdate(ctx context.Context, event stripe.Event) error {

	logger.Infof("onAccountUpdate event = %+v", event)

	var account stripe.Account
	err := json.Unmarshal(event.Data.Raw, &account)
	if err != nil {
		return err
	}

	logger.Infof("onAccountUpdate account to be created: ID = %v, Email = %v, Individual = %v", account.ID, account.Email, account.Individual)

	if !account.DetailsSubmitted {
		logger.Info("account creation is not finalized, skipping the stripe notification")
		return nil
	}

	userId, err := storage.GetAccUserId(ctx, account.ID)
	if err != nil {
		logger.Error("failed get account user if from storage", err)
		return err
	}
	// TODO check account's validity
	//account.Individual.FirstName
	//account.Individual.LastName
	//account.ExternalAccounts.Data[0].BankAccount.ID

	// TODO this bool indicates if the account first/last names matches our platforms data
	var partyCheckUserResult = true

	if account.ExternalAccounts == nil || len(account.ExternalAccounts.Data) < 1 {
		logger.Error("data returned from the Stripe is not expected", NoExternalAccount)
		return NoExternalAccount
	}

	conn, err := services.ConnectToParty()
	if err != nil {
		logger.Error("failed to connect to the party service", err)
		return err
	}
	partyC := proto.NewInterServiceClient(conn)
	_, err = partyC.InterServiceUpdateProfileAdditionalInfo(ctx, &proto.InterServiceUpdateProfileAdditionalInfoRequest{
		UserId: userId,
		Info: &proto.InterServiceUpdateProfileAdditionalInfoRequest_BankAccount{
			BankAccount: &proto.BankAccount{
				// TODO: check this part later: assuming only one account would be associated
				BankName:          account.ExternalAccounts.Data[0].BankAccount.BankName,
				AccountNumber:     account.ExternalAccounts.Data[0].BankAccount.ID,
				AccountHolderName: account.ExternalAccounts.Data[0].BankAccount.AccountHolderName,
				// TODO: make sure that routing number is the same as sort code
				SortCode: account.ExternalAccounts.Data[0].BankAccount.RoutingNumber,
				// TODO: check whether it is OK to leave document ids empty here
				//DocumentIds:       nil,
			},
		},
	})
	if err != nil {
		logger.Error("failed to update bank account in the party service", err)
		return err
	}

	pubsub.PublishById(pubsub.OnboardingResult(financePb.OnboardingResult{
		Successful: partyCheckUserResult,
		Desc:       "",
	}), userId)

	return nil
}
