package scheduler

import (
	"github.com/kkjhamb01/courier-management/common/config"
	"github.com/kkjhamb01/courier-management/common/logger"
	"github.com/kkjhamb01/courier-management/common/logger/tag"
	"github.com/kkjhamb01/courier-management/finance/business"
	"github.com/robfig/cron/v3"
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
