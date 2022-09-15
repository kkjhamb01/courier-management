package business

import (
	"encoding/base64"
	"fmt"
	gproto "github.com/golang/protobuf/proto"
	"gitlab.artin.ai/backend/courier-management/common/config"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	"gitlab.artin.ai/backend/courier-management/party/proto"
	uaapb "gitlab.artin.ai/backend/courier-management/uaa/proto"
	"gitlab.artin.ai/backend/courier-management/uaa/security"
	"strings"
	"testing"
)

func init(){
	config.InitTestConfig()
	logger.InitLogger()
}

func TestParty1(t *testing.T) {
	t.Skip()
	partyApi := NewPartyAPI(config.Uaa())
	jwt,_ := security.NewJWTUtils(config.Jwt())
	user,_ := partyApi.FindUserByPhoneNumber("989126031724", proto.UserType_USER_TYPE_CURIOUR)
	fmt.Print(user)
	token,_ := jwt.GenerateToken(*user)
	fmt.Print(token)
}

func TestEncoding(t *testing.T)  {
	var decoded []byte
	var b = "AAAAACYKGWJlaG5hbS5uaWtiYWtodEBnbWFpbC5jb20SCUJlaG5hbTIwMQ=="
	//base64.StdEncoding.Decode(*decoded, []byte(b))
	decoded, _ = base64.StdEncoding.DecodeString(b)
	sdecoded := string(decoded)
	decoded = []byte(sdecoded[strings.Index(sdecoded, "&")+1:])
	m := uaapb.AdminLoginRequest{}
	gproto.Unmarshal(decoded, &m)
	fmt.Printf("message = %v\n", m)
}