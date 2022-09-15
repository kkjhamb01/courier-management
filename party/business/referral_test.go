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

func TestReferral(t *testing.T) {
	service := NewService(config.GetData(), config.Jwt())

	var idList = []string{
		"eb63b23f-ab55-13f0-c23f-ea309bd25aac",
		"bb63b23f-ab55-13f0-c23f-ea309bd25aac",
		"cb63b23f-ab55-13f0-c23f-ea309bd25aac",
		"db63b23f-ab55-13f0-c23f-ea309bd25aac",
		"ab63b23f-ab55-13f0-c23f-ea309bd25aac",
		"fb63b23f-ab55-13f0-c23f-ea309bd25aac",
		"ee63b23f-ab55-13f0-c23f-ea309bd25aac",
		"e363b23f-ab55-13f0-c23f-ea309bd25aac",
		"e963b23f-ab55-13f0-c23f-ea309bd25aac",
		"e063b23f-ab55-13f0-c23f-ea309bd25aac",
	}
	for _,id := range idList{
		referral := service.CalculateReferral(id)
		fmt.Printf("id = %v, referral = %v\n", id, referral)
	}
}