package security

type User struct {
	Id                   string
	Name                 string
	FirstName            string
	LastName             string
	Email                string
	PhoneNumber          string
	DeviceID          	 string
	Roles                []Role
	Claims				 []Claim
}

type Claim struct {
	ClaimType 	ClaimType	`json:"claim_type"`
	Identifier	string		`json:"identifier"`
}

type Role int32

const (
	Role_UNKNOWN_ROLE Role = 0
	Role_ADMIN        Role = 1
	Role_CLIENT       Role = 2
	Role_COURIER      Role = 3
)

var Role_name = map[int32]string{
	0: "UNKNOWN_ROLE",
	1: "ADMIN",
	2: "CLIENT",
	3: "COURIER",
}

var Role_value = map[string]int32{
	"UNKNOWN_ROLE": 0,
	"ADMIN":        1,
	"CLIENT":    2,
	"COURIER":       3,
}

func (x Role) String() string {
	return Role_name[int32(x)]
}

type ClaimType int32

const (
	CLAIM_TYPE_UNKNOWN_CLAIM_TYPE ClaimType = 0
	CLAIM_TYPE_EMAIL              ClaimType = 1
	CLAIM_TYPE_PHONE_NUMBER       ClaimType = 2
	CLAIM_TYPE_GOOGLE_ID 		 ClaimType = 3
	CLAIM_TYPE_FACEBOOK_ID 		 ClaimType = 4
)

var ClaimType_name = map[int32]string{
	0: "CLAIM_TYPE_UNKNOWN_CLAIM_TYPE",
	1: "CLAIM_TYPE_EMAIL",
	2: "CLAIM_TYPE_PHONE_NUMBER",
	3: "CLAIM_TYPE_Google",
	4: "CLAIM_TYPE_Facebook",
}

var ClaimType_value = map[string]int32{
	"CLAIM_TYPE_UNKNOWN_CLAIM_TYPE": 0,
	"CLAIM_TYPE_EMAIL":        1,
	"CLAIM_TYPE_PHONE_NUMBER":    2,
	"CLAIM_TYPE_Google":			3,
	"CLAIM_TYPE_Facebook":			4,
}

func (x ClaimType) String() string {
	return ClaimType_name[int32(x)]
}
