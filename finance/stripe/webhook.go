package stripe

import (
	"github.com/gofiber/fiber"
	limiter "github.com/shareed2k/fiber_limiter"
	"github.com/stripe/stripe-go/v72"
	"gitlab.artin.ai/backend/courier-management/common/config"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	"gitlab.artin.ai/backend/courier-management/common/logger/tag"
	"gitlab.artin.ai/backend/courier-management/finance/db"
	"time"
)

var app *fiber.App

func StartWebhook() {

	logger.Infof("StartWebhook")

	app = fiber.New()

	// 3 requests per second max
	cfg := limiter.Config{
		Rediser:   db.RedisV7Client(),
		Max:       3,
		Burst:     3,
		Period:    time.Second,
		Algorithm: limiter.GCRAAlgorithm,
		Handler:   webhookHandler,
	}
	app.Use(limiter.New(cfg))

	app.Post("/webhook", webhookHandler)
	app.Post("/return", other)
	app.Post("/refresh", other)

	err := app.Listen(":" + config.Finance().WebhookPort)
	if err != nil {
		logger.Fatal("failed to start stripe webhook server", tag.Err("err", err))
	}
}

func other(c *fiber.Ctx) {
	logger.Info("message from stripe: " + c.Body())
}

func webhookHandler(c *fiber.Ctx) {
	event := stripe.Event{}
	err := c.BodyParser(&event)

	logger.Infof("webhookHandler event = %+v", event)

	if err != nil {
		logger.Error("webhookHandler failed to parse webhook request body", err)
		c.Status(500)
		return
	}

	logger.Info("webhookHandler Stripe webhook called", tag.Obj("event", event))

	switch event.Type {
	case "payment_method.detached":
		err = onPaymentMethodDetached(c.Context(), event)
		if err != nil {
			logger.Error("webhookHandler failed to handle on payment method detached event", err)
			c.Status(500)
			return
		}
	case "payment_method.attached":
		err = onPaymentMethodAttached(c.Context(), event)
		if err != nil {
			logger.Error("webhookHandler failed to handle on payment method attached event", err)
			c.Status(500)
			return
		}
	case "account.updated":
		err = onAccountUpdate(c.Context(), event)
		if err != nil {
			logger.Error("webhookHandler failed to handle on account update event", err)
			c.Status(500)
			return
		}
	}

	return
}

func StopWebhook() {
	err := app.Shutdown()
	if err != nil {
		logger.Error("failed to start stripe webhook server", err)
	}
}
