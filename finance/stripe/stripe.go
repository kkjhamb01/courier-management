package stripe

import (
	"github.com/stripe/stripe-go/v72"
	"gitlab.artin.ai/backend/courier-management/common/config"
)

func Setup() {
	stripe.Key = config.Finance().StripeSecretKey
}
