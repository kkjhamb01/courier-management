package scheduler

import (
	"github.com/robfig/cron/v3"
	"gitlab.artin.ai/backend/courier-management/common/config"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	"gitlab.artin.ai/backend/courier-management/common/logger/tag"
	"gitlab.artin.ai/backend/courier-management/delivery/business"
)

var createRequestSchedule *cron.Cron

func StartCreateRequest() {

	logger.Infof("StartCreateRequest")

	duration := config.Delivery().CreateRequestScheduleDuration
	job := business.CreateRequestSchedulerJob

	var err error
	createRequestSchedule, err = startSchedule(duration, job)
	if err != nil {
		logger.Fatal("failed to start create request schedule", tag.Err("err", err))
		panic(err)
	}
}

func StopCreateRequest() {
	stopSchedule(createRequestSchedule)
}
