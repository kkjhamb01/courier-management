package business

import (
	"math/rand"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/hashicorp/go-uuid"
	"github.com/kkjhamb01/courier-management/common/config"
	"github.com/kkjhamb01/courier-management/common/logger"
	"github.com/kkjhamb01/courier-management/uaa/db"
	"github.com/kkjhamb01/courier-management/uaa/security"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"golang.org/x/oauth2/google"
)

const (
	REDIS_KEY_OTP             = "t"
	REDIS_KEY_REMAINING_RETRY = "r"
	REDIS_KEY_PHONE_NUMBER    = "p"
	REDIS_KEY_OAUTH_TYPE      = "o"
	REDIS_KEY_USER_TYPE       = "u"
	REDIS_KEY_USER_IDENTIFIER = "i"
	REDIS_KEY_CLAIM_TYPE      = "c"
	REDIS_KEY_USER_ID         = "d"
	REDIS_KEY_USER            = "e"
	REDIS_KEY_OAUTH_RESPONSE  = "n"
	USER_TYPE_COURIER         = 1
	USER_TYPE_USER            = 2
)

var (
	ALPHA_NUM_CHARS = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	NUM_CHARS = []rune("0123456789")

	driverScopes = []string{
		"https://www.googleapis.com/auth/userinfo.profile",
		"https://www.googleapis.com/auth/userinfo.email",
	}

	passengerScopes = []string{
		"https://www.googleapis.com/auth/userinfo.profile",
		"https://www.googleapis.com/auth/userinfo.email",
	}
)

type RegistrationService struct {
	config                config.UaaData
	jwtUtils              security.JWTUtils
	redis                 *redis.Client
	partyAPI              PartyAPI
	googleCourierConfig   *oauth2.Config
	googleUserConfig      *oauth2.Config
	facebookCourierConfig *oauth2.Config
	facebookUserConfig    *oauth2.Config
}

type facebookUserInfo struct {
	Id             string `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	Email          string `json:"email,omitempty"`
	FirstName      string `json:"first_name,omitempty"`
	LastName       string `json:"last_name,omitempty"`
	ProfilePicture []byte
}

func (s *RegistrationService) generateOtp() string {
	otp := s.config.Otp
	if otp.Static != "" {
		return otp.Static
	}
	l := otp.Length
	b := make([]rune, l)
	for i := range b {
		b[i] = NUM_CHARS[rand.Intn(len(NUM_CHARS))]
	}
	return string(b)
}

func (s *RegistrationService) sendOtp(phoneNumber string, otp string) error {
	return nil
}

func NewRegistrationService(config config.UaaData, jwtConfig config.JwtData, partyApi PartyAPI) *RegistrationService {
	jwtUtils, err := security.NewJWTUtils(jwtConfig)
	if err != nil {
		logger.Fatalf("cannot create jwtutils ", err)
	}

	facebookCourierConfig := &oauth2.Config{
		ClientID:     config.Facebook.Driver.ClientID,
		ClientSecret: config.Facebook.Driver.ClientSecret,
		RedirectURL:  config.Facebook.Driver.RedirectURL,
		Scopes:       []string{"email"},
		Endpoint:     facebook.Endpoint,
	}

	facebookUserConfig := &oauth2.Config{
		ClientID:     config.Facebook.Passenger.ClientID,
		ClientSecret: config.Facebook.Passenger.ClientSecret,
		RedirectURL:  config.Facebook.Passenger.RedirectURL,
		Scopes:       []string{"email"},
		Endpoint:     facebook.Endpoint,
	}

	googleCourierConfig := &oauth2.Config{
		ClientID:     config.Google.Driver.ClientID,
		ClientSecret: config.Google.Driver.ClientSecret,
		RedirectURL:  config.Google.Driver.RedirectURL,
		Scopes:       driverScopes,
		Endpoint:     google.Endpoint,
	}
	googleUserConfig := &oauth2.Config{
		ClientID:     config.Google.Passenger.ClientID,
		ClientSecret: config.Google.Passenger.ClientSecret,
		RedirectURL:  config.Google.Passenger.RedirectURL,
		Scopes:       passengerScopes,
		Endpoint:     google.Endpoint,
	}

	rand.Seed(time.Now().UnixNano())
	return &RegistrationService{
		config:                config,
		jwtUtils:              jwtUtils,
		redis:                 db.NewRedisConnection(config.Redis),
		partyAPI:              partyApi,
		facebookCourierConfig: facebookCourierConfig,
		facebookUserConfig:    facebookUserConfig,
		googleCourierConfig:   googleCourierConfig,
		googleUserConfig:      googleUserConfig,
	}
}

func generateId(isCourier bool) string {
	buf, err := uuid.GenerateRandomBytes(16)
	if err != nil {
		return ""
	}
	buf[0] = buf[0] / 2
	if isCourier {
		buf[0] = buf[0] | 128
	}
	id, _ := uuid.FormatUUID(buf)
	return id
}
