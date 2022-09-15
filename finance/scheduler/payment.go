package scheduler

import (
	"github.com/robfig/cron/v3"
	"gitlab.artin.ai/backend/courier-management/common/config"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	"gitlab.artin.ai/backend/courier-management/common/logger/tag"
	"gitlab.artin.ai/backend/courier-management/finance/business"
)

var createPaymentSchedule *cron.Cron

func StartCreatePayment() {

	logger.Infof("StartCreatePayment")

	duration := config.Finance().CronJobSchedule
	job := func() {
		//logger.Info("Payment Schedule Started")
		business.MakeSettlementPayments()
	}

	var err error
	createPaymentSchedule, err = startSchedule(duration, job)
	if err != nil {
		logger.Fatal("failed to start create payment schedule", tag.Err("err", err))
		panic(err)
	}
}

func StopCreateRequest() {
	stopSchedule(createPaymentSchedule)
}
