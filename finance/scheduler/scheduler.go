package scheduler

import (
	"github.com/kkjhamb01/courier-management/common/logger"
	"github.com/robfig/cron/v3"
)

func startSchedule(duration string, job func()) (*cron.Cron, error) {

	logger.Infof("startSchedule duration = %v", duration)

	c := cron.New()
	_, err := c.AddFunc(duration, job)
	if err != nil {
		logger.Error("failed to add function to cron", err)
		return nil, err
	}
	c.Start()

	return c, nil
}

func stopSchedule(c *cron.Cron) {
	c.Stop()
}
