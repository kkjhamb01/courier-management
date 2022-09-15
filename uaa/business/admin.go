package business

import (
	"context"
	"math/rand"
	"time"

	"github.com/kkjhamb01/courier-management/common/config"
	"github.com/kkjhamb01/courier-management/common/logger"
	pb "github.com/kkjhamb01/courier-management/uaa/proto"
	"github.com/kkjhamb01/courier-management/uaa/security"
)

type AdminService struct {
	config   config.UaaData
	jwtUtils security.JWTUtils
	partyAPI PartyAPI
}

func (s *AdminService) AdminLogin(ctx context.Context, in *pb.AdminLoginRequest) (*pb.AdminLoginResponse, error) {
	logger.Debugf("AdminLogin username = %v, password = %v", in.GetUsername(), in.GetPassword())
	existingUser, err := s.partyAPI.AdminLogin(in.GetUsername(), in.GetPassword())
	logger.Debugf("existingUser = %v", existingUser)
	if err != nil {
		logger.Errorf("cannot query party to get amdin user", err)
		return nil, pb.Unauthenticated.ErrorMsg(err.Error())
	}
	if existingUser == nil {
		logger.Debugf("cannot find user %v", in.GetUsername())
		return nil, pb.Unauthenticated.ErrorMsg("user not found")
	}

	var name = existingUser.Name
	var userId = existingUser.Id
	var claims = existingUser.Claims

	user := security.User{
		Id:     userId,
		Roles:  []security.Role{security.Role_ADMIN},
		Name:   name,
		Claims: claims,
	}

	token, err := s.jwtUtils.GenerateToken(user)

	logger.Debugf("token = %v", token)

	return &pb.AdminLoginResponse{
		Token: token,
	}, err
}

func NewAdminService(config config.UaaData, jwtConfig config.JwtData, partyApi PartyAPI) *AdminService {
	jwtUtils, err := security.NewJWTUtils(jwtConfig)
	if err != nil {
		logger.Fatalf("cannot create jwtutils ", err)
	}
	rand.Seed(time.Now().UnixNano())
	return &AdminService{
		config:   config,
		jwtUtils: jwtUtils,
		partyAPI: partyApi,
	}
}
