package storage

import (
	"github.com/kkjhamb01/courier-management/common/config"
	"github.com/kkjhamb01/courier-management/common/logger"
)

func init() {
	config.InitTestConfig()
	logger.InitLogger()
}
