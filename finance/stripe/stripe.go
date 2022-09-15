package stripe

import (
	"github.com/kkjhamb01/courier-management/common/config"
	"github.com/stripe/stripe-go/v72"
)

func Setup() {
	stripe.Key = config.Finance().StripeSecretKey
}
