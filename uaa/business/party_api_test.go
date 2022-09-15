package business

import (
	"encoding/base64"
	"fmt"
	"strings"
	"testing"

	gproto "github.com/golang/protobuf/proto"
	"github.com/kkjhamb01/courier-management/common/config"
	"github.com/kkjhamb01/courier-management/common/logger"
	"github.com/kkjhamb01/courier-management/party/proto"
	uaapb "github.com/kkjhamb01/courier-management/uaa/proto"
	"github.com/kkjhamb01/courier-management/uaa/security"
)

func init() {
	config.InitTestConfig()
	logger.InitLogger()
}

func TestParty1(t *testing.T) {
	t.Skip()
	partyApi := NewPartyAPI(config.Uaa())
	jwt, _ := security.NewJWTUtils(config.Jwt())
	user, _ := partyApi.FindUserByPhoneNumber("989126031724", proto.UserType_USER_TYPE_CURIOUR)
	fmt.Print(user)
	token, _ := jwt.GenerateToken(*user)
	fmt.Print(token)
}

func TestEncoding(t *testing.T) {
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
