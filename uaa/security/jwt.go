package security

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	guid "github.com/google/uuid"
	"github.com/kkjhamb01/courier-management/common/config"
	"github.com/kkjhamb01/courier-management/common/logger"
	model "github.com/kkjhamb01/courier-management/uaa/proto"
)

const (
	CLAIM_AUDIENCE     = "aud"
	CLAIM_ID           = "id"
	CLAIM_ISSUED_AT    = "iat"
	CLAIM_ISSUER       = "iss"
	CLAIM_TYPE         = "typ"
	CLAIM_SUB          = "sub"
	CLAIM_EMAIL        = "email"
	CLAIM_ROLES        = "roles"
	CLAIM_PHONE_NUMBER = "phone_number"
	CLAIM_STATUS       = "status"
	CLAIM_EXP          = "exp"
	CLAIM_KID          = "kid"
	CLAIM_DEVICE       = "device"
	CLAIM_DEVICE_ID    = "device_id"
	LOCAL_USER         = "user"
	CLAIM_AUTHORIZED   = "authorized"
)

type JWTUtils struct {
	config       config.JwtData
	verifyKeyStr string
	signKey      *rsa.PrivateKey
	verifyKey    *rsa.PublicKey
	parser       jwt.Parser
	Kid          string
	Alg          string
	Kty          string
}

func (jwtUtils JWTUtils) init(config config.JwtData) (JWTUtils, error) {
	workingDir := os.Getenv("COURIER_MANAGEMENT_WORKING_DIR")
	if workingDir == "" {
		workingDir = "."
	}
	signBytes, err := ioutil.ReadFile(workingDir + "/" + config.SignKey)
	if err != nil {
		logger.Fatal("error in loading sign key file")
	}

	jwtUtils.signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		logger.Fatal("error in loading sign key")
	}

	verifyBytes, err := ioutil.ReadFile(workingDir + "/" + config.VerifyKey)
	if err != nil {
		logger.Fatal("error in loading verify key file")
	}

	jwtUtils.verifyKeyStr = string(verifyBytes)

	jwtUtils.verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		logger.Fatal("error in loading verify key")
	}

	jwtUtils.config = config

	jwtUtils.parser = jwt.Parser{}

	jwtUtils.Kid = "2021jwt"
	jwtUtils.Alg = "RS256"
	jwtUtils.Kty = "RSA"

	return jwtUtils, nil
}

func (jwtUtils JWTUtils) GenerateToken(user User) (*model.Token, error) {
	token := jwt.New(jwt.GetSigningMethod(jwtUtils.config.SigningMethod))
	claims := token.Claims.(jwt.MapClaims)
	claims[CLAIM_ID] = guid.New().String()
	claims[CLAIM_AUDIENCE] = "courier-management-uaa"
	claims[CLAIM_ISSUED_AT] = time.Now().Unix()
	claims[CLAIM_ISSUER] = jwtUtils.config.Issuer
	claims[CLAIM_TYPE] = "bearer"
	claims[CLAIM_KID] = jwtUtils.Kid
	if len(user.Claims) > 0 {
		claims[CLAIM_AUTHORIZED] = user.Claims
	}
	if user.Id != "" {
		claims[CLAIM_SUB] = user.Id
	}
	if user.DeviceID != "" {
		claims[CLAIM_DEVICE] = user.DeviceID
	}
	if user.PhoneNumber != "" {
		claims[CLAIM_PHONE_NUMBER] = user.PhoneNumber
	}
	if user.Email != "" {
		claims[CLAIM_EMAIL] = user.Email
	}
	var roles []string
	for _, role := range user.Roles {
		roles = append(roles, role.String())
	}
	if len(roles) > 0 {
		claims[CLAIM_ROLES] = roles
	}
	claims[CLAIM_EXP] = time.Now().Add(time.Minute * jwtUtils.config.ExpirationMinutes).Unix()
	accessToken, err := token.SignedString(jwtUtils.signKey)
	if err != nil {
		return nil, err
	}

	token = jwt.New(jwt.GetSigningMethod(jwtUtils.config.SigningMethod))
	claims = token.Claims.(jwt.MapClaims)
	if user.Id != "" {
		claims[CLAIM_SUB] = user.Id
	}
	if len(roles) > 0 {
		claims[CLAIM_ROLES] = roles
	}
	claims[CLAIM_EXP] = time.Now().Add(time.Hour * jwtUtils.config.RefreshExpirationHours).Unix()

	refreshToken, err := token.SignedString(jwtUtils.signKey)
	if err != nil {
		return nil, err
	}

	return &model.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

type CustomClaims struct {
	Email       string   `json:"email"`
	Roles       []string `json:"roles"`
	Name        string   `json:"name"`
	GivenName   string   `json:"given_name"`
	FamilyName  string   `json:"family_name"`
	Status      string   `json:"status"`
	PhoneNumber string   `json:"phone_number"`
	Authorized  []Claim  `json:"authorized"`
	jwt.StandardClaims
}

func (c CustomClaims) Valid() error {
	return nil
}

func (jwtUtils JWTUtils) ValidateHeader(auth string) (*User, error) {
	if len(auth) <= 7 || strings.ToLower(auth[:6]) != "bearer" {
		return nil, errors.New("invalid authorization header")
	}
	auth = auth[7:]
	return jwtUtils.Validate(auth)
}
func (jwtUtils JWTUtils) Validate(auth string) (*User, error) {
	token, err := jwt.ParseWithClaims(auth, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtUtils.verifyKey, nil
	})
	var u User
	if err != nil {
		logger.Error("invalid token ", err)
	} else {
		claims := token.Claims.(*CustomClaims)

		var userRoles []Role
		for _, role := range claims.Roles {
			userRoles = append(userRoles, Role(Role_value[role]))
		}

		u = User{
			Id:          claims.Subject,
			PhoneNumber: claims.PhoneNumber,
			Email:       claims.Email,
			Roles:       userRoles,
			Claims:      claims.Authorized,
		}
	}
	return &u, err
}

func (jwtUtils JWTUtils) ValidateUnsignedHeader(auth string) (*User, error) {
	if len(auth) <= 7 || strings.ToLower(auth[:6]) != "bearer" {
		return nil, errors.New("invalid authorization header")
	}
	auth = auth[7:]
	return jwtUtils.ValidateUnsigned(auth, true)
}
func (jwtUtils JWTUtils) ValidateUnsigned(auth string, validateIssuer bool) (*User, error) {
	token, err := jwtUtils.parseToken(jwtUtils.parser, auth, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtUtils.verifyKey, nil
	})
	if err2, ok := err.(*jwt.ValidationError); ok {
		if jwt.ValidationErrorExpired == err2.Errors {
			return nil, model.TokenIsExpired.ErrorNoMsg()
		}
	}
	if validateIssuer {
		if claims, ok := token.Claims.(*CustomClaims); ok {
			if claims.Issuer != jwtUtils.config.Issuer {
				return nil, model.Unauthenticated.ErrorMsg(fmt.Sprintf("token issuer is not valid, token issuer is %v, should be = %v", claims.Issuer, jwtUtils.config.Issuer))
			}
		}
	}

	var u User
	if err != nil {
		logger.Error("invalid token ", err)
	} else {
		claims := token.Claims.(*CustomClaims)

		var userRoles []Role
		for _, role := range claims.Roles {
			userRoles = append(userRoles, Role(Role_value[role]))
		}

		u = User{
			Id:          claims.Subject,
			PhoneNumber: claims.PhoneNumber,
			Email:       claims.Email,
			Name:        claims.Name,
			Roles:       userRoles,
			Claims:      claims.Authorized,
			FirstName:   claims.GivenName,
			LastName:    claims.FamilyName,
		}
	}
	return &u, err
}

func (jwtUtils JWTUtils) ValidateRefreshToken(refreshToken string) (*User, error) {
	token, err := jwtUtils.parseToken(jwtUtils.parser, refreshToken, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtUtils.verifyKey, nil
	})
	if err2, ok := err.(*jwt.ValidationError); ok {
		if jwt.ValidationErrorExpired == err2.Errors {
			return nil, model.TokenIsExpired.ErrorNoMsg()
		}
	}
	var u User
	if err != nil {
		logger.Error("invalid token ", err)
	} else {
		claims := token.Claims.(*CustomClaims)

		var userRoles []Role
		for _, role := range claims.Roles {
			userRoles = append(userRoles, Role(Role_value[role]))
		}

		u = User{
			Id:          claims.Subject,
			PhoneNumber: claims.PhoneNumber,
			Email:       claims.Email,
			Name:        claims.Name,
			Roles:       userRoles,
		}
	}
	return &u, err
}

func (jwtUtils JWTUtils) parseToken(p jwt.Parser, tokenString string, claims jwt.Claims, keyFunc jwt.Keyfunc) (*jwt.Token, error) {
	token, _, err := p.ParseUnverified(tokenString, claims)
	if err != nil {
		return token, err
	}

	// Verify signing method is in the required set
	if p.ValidMethods != nil {
		var signingMethodValid = false
		var alg = token.Method.Alg()
		for _, m := range p.ValidMethods {
			if m == alg {
				signingMethodValid = true
				break
			}
		}
		if !signingMethodValid {
			// signing method is not in the listed set
			return token, jwt.NewValidationError(fmt.Sprintf("signing method %v is invalid", alg), jwt.ValidationErrorSignatureInvalid)
		}
	}

	// Lookup key
	var _ interface{}
	if keyFunc == nil {
		// keyFunc was not provided.  short circuiting validation
		return token, jwt.NewValidationError("no Keyfunc was provided.", jwt.ValidationErrorUnverifiable)
	}
	if _, err = keyFunc(token); err != nil {
		// keyFunc returned an error
		if ve, ok := err.(*jwt.ValidationError); ok {
			return token, ve
		}
		return token, &jwt.ValidationError{Inner: err, Errors: jwt.ValidationErrorUnverifiable}
	}

	vErr := &jwt.ValidationError{}

	// Validate Claims
	if !p.SkipClaimsValidation {
		if err := token.Claims.Valid(); err != nil {

			// If the Claims Valid returned an error, check if it is a validation error,
			// If it was another error type, create a ValidationError with a generic ClaimsInvalid flag set
			if e, ok := err.(*jwt.ValidationError); !ok {
				vErr = &jwt.ValidationError{Inner: err, Errors: jwt.ValidationErrorClaimsInvalid}
			} else {
				vErr = e
			}
		}
	}

	if vErr.Errors == 0 {
		token.Valid = true
		return token, nil
	}

	return token, vErr
}

func (jwtUtils JWTUtils) Sign(plain string, secret string) (string, error) {
	return jwt.GetSigningMethod(jwtUtils.config.SigningMethod).Sign(plain+secret, jwtUtils.signKey)
}

func (jwtUtils JWTUtils) Verify(plain string, secret string, sign string) error {
	return jwt.GetSigningMethod(jwtUtils.config.SigningMethod).Verify(plain+secret, sign, jwtUtils.verifyKey)
}

func (jwtUtils JWTUtils) GetX5c() string {
	return jwtUtils.verifyKeyStr
}

func NewJWTUtils(config config.JwtData) (JWTUtils, error) {
	var jwtUtils JWTUtils
	return jwtUtils.init(config)
}
