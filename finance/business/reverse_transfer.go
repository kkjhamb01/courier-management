package business

import (
	"errors"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/reversal"
	"gitlab.artin.ai/backend/courier-management/common/logger"
)

func ReveresTransfer(tr *stripe.Transfer) error {
	if tr == nil {
		return errors.New("stripe.Transfer is nil")
	}

	params := &stripe.ReversalParams{
		Transfer:    stripe.String(tr.ID),
		Description: stripe.String("finance service failure"),
	}
	_, err := reversal.New(params)
	if err != nil {
		logger.Error("failed to reveres transfer", err)
		return err
	}

	return nil
}
