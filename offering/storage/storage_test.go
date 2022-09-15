package storage

import (
	"gitlab.artin.ai/backend/courier-management/common/config"
	"gitlab.artin.ai/backend/courier-management/common/logger"
)

func init() {
	config.InitTestConfig()
	logger.InitLogger()
}
