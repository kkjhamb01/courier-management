package business

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/kkjhamb01/courier-management/common/config"
	"github.com/kkjhamb01/courier-management/common/logger"
	_ "github.com/kkjhamb01/courier-management/party/proto"
	pb "github.com/kkjhamb01/courier-management/party/proto"
)

func init() {
	config.InitTestConfig()
	logger.InitLogger()
}

func TestMotError(t *testing.T) {
	service := NewService(config.GetData(), config.Jwt())
	ctx := context.Background()
	response, err := service.SearchMot(ctx, &pb.SearchMotRequest{
		AccessToken:        "1234",
		RegistrationNumber: "SP58SAM",
	})
	if err != nil {
		logger.Fatalf("cannot connect to database", err)
	}
	logger.Infof("response = %v", response)
}

func TestParseDate(t *testing.T) {
	t.Skip()
	date, err := time.Parse(time.RFC3339, fmt.Sprintf("%s-01T11:45:26.371Z", "2019-10"))
	logger.Debugf("date = %v %v", date, err)
	if time.Now().Sub(date).Hours() > 24*365 {
		logger.Debugf("month of first registration is more than one year ago")
	} else {
		logger.Debugf("month of first registration is less than one year ago")
	}
}
