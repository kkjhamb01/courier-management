package security

import (
	"encoding/json"
	"fmt"
	"gitlab.artin.ai/backend/courier-management/common/config"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	"testing"
)

func init(){
	config.InitTestConfig()
	logger.InitLogger()
}

func TestGenerateToken(t *testing.T){
	jwtUtils,_ := NewJWTUtils(config.Jwt())
	token := "eyJhbGciOiJSUzI1NiIsImtpZCI6IjdmNTQ4ZjY3MDg2OTBjMjExMjBiMGFiNjY4Y2FhMDc5YWNiYzJiMmYiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJodHRwczovL2FjY291bnRzLmdvb2dsZS5jb20iLCJhenAiOiI0ODQ4NDczNzQyNjMtNmh2NjJnMWwzMGRoM2RmODFlb2EwZGduYzNnamE4ZjYuYXBwcy5nb29nbGV1c2VyY29udGVudC5jb20iLCJhdWQiOiI0ODQ4NDczNzQyNjMtNmh2NjJnMWwzMGRoM2RmODFlb2EwZGduYzNnamE4ZjYuYXBwcy5nb29nbGV1c2VyY29udGVudC5jb20iLCJzdWIiOiIxMTI0Mjc1ODE3MDc1MTUwODM2NzMiLCJlbWFpbCI6ImJlaG5hbS5uaWtiYWtodEBnbWFpbC5jb20iLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwiYXRfaGFzaCI6IjlLUFlvNXBsM2JFckp2Zi1kbDk0YXciLCJuYW1lIjoiQmVobmFtIE5pa2Jha2h0IiwicGljdHVyZSI6Imh0dHBzOi8vbGgzLmdvb2dsZXVzZXJjb250ZW50LmNvbS9hL0FBVFhBSnpNZ3A5azZ2TVFTNXBoQmhXVGlnTGhOanl3MFkwTDN6dGs4MXo4Unc9czk2LWMiLCJnaXZlbl9uYW1lIjoiQmVobmFtIiwiZmFtaWx5X25hbWUiOiJOaWtiYWtodCIsImxvY2FsZSI6ImVuIiwiaWF0IjoxNjI2NzE0OTY4LCJleHAiOjE2MjY3MTg1Njh9.C3RPQLP_r6b4yPWpzB6pUyvLYkEOQifM8axn3j7981CZo_dzUeT1C_ecF7r1qrP7kB3-IhVhPcTBT_dbXMZl9RtzHUFogTm7zR46-GD-scKmKh_MM8ZjR6LppMtOmaLC28pgGPixqllw8a-EFraRQu1q-fF4LTj6tvNvnvH57_kgFWR-l-JI9vZskGA9_Px1pEJCypq5mXFZKhXRZvI8z54PCGtL1277gwRkWzAXYIS9Pby2SMfMglL57WmbjR6jZzsJDIV5CaFo6ZfNcQBo9k4ZcbLwQhIlqaiuIgf8Kx2gtrZInpkBy2lBCF9GrzLzHAkbWaDomb0obU9oQgzsCA"
	user,_ := jwtUtils.ValidateUnsigned(token, false)
	fmt.Println(user)

	val := "{\"Id\":\"a7082937-43ca-d44d-8fb7-44b19c32fac0\",\"Name\":\"\",\"FirstName\":\"\",\"LastName\":\"\",\"Email\":\"behnam.nikbakht@gmail.com\",\"PhoneNumber\":\"\",\"DeviceID\":\"\",\"Roles\":[2],\"Claims\":[{\"claim_type\":4,\"identifier\":\"10219370295430496\"}]}"
	var tokenUser User
	err := json.Unmarshal([]byte(val), &tokenUser)
	if err != nil{
		fmt.Errorf("err %v", err)
	}
}