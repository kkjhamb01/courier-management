package business

import (
	"fmt"
	"gitlab.artin.ai/backend/courier-management/common/config"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	"testing"
)

func init(){
	config.InitTestConfig()
	logger.InitLogger()
}

func TestId(t *testing.T) {
	for i:=1; i<11; i=i+1 {
		id := generateId(true)
		fmt.Printf("courier id = %v\n", id)
	}
	for i:=1; i<11; i=i+1 {
		id := generateId(false)
		fmt.Printf("client id = %v\n", id)
	}
}
