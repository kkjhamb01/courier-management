package business

import (
	"fmt"
	"testing"

	"github.com/kkjhamb01/courier-management/common/config"
	"github.com/kkjhamb01/courier-management/common/logger"
)

func init() {
	config.InitTestConfig()
	logger.InitLogger()
}

func TestId(t *testing.T) {
	for i := 1; i < 11; i = i + 1 {
		id := generateId(true)
		fmt.Printf("courier id = %v\n", id)
	}
	for i := 1; i < 11; i = i + 1 {
		id := generateId(false)
		fmt.Printf("client id = %v\n", id)
	}
}
