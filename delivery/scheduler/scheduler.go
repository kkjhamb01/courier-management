package scheduler

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	"time"
)

func startSchedule(duration time.Duration, job func()) (*cron.Cron, error) {
	functionParam := fmt.Sprintf("@every %vs", duration.Seconds())

	c := cron.New()
	_, err := c.AddFunc(functionParam, job)
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
