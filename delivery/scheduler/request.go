package scheduler

import (
	"github.com/kkjhamb01/courier-management/common/config"
	"github.com/kkjhamb01/courier-management/common/logger"
	"github.com/kkjhamb01/courier-management/common/logger/tag"
	"github.com/kkjhamb01/courier-management/delivery/business"
	"github.com/robfig/cron/v3"
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
