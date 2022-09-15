package scheduler

import (
	"github.com/robfig/cron/v3"
	"gitlab.artin.ai/backend/courier-management/common/logger"
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
