package business

import (
	"context"
	"gitlab.artin.ai/backend/courier-management/common/config"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	"gitlab.artin.ai/backend/courier-management/party/proto"
	pb "gitlab.artin.ai/backend/courier-management/uaa/proto"
	"gitlab.artin.ai/backend/courier-management/uaa/security"
	"math/rand"
	"time"
)

type TokenService struct {
	config   config.UaaData
	jwtUtils security.JWTUtils
	partyAPI PartyAPI
}

func (s *TokenService) RefreshToken(ctx context.Context, in *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error){
	user,err := s.jwtUtils.ValidateRefreshToken(in.RefreshToken)
	if err != nil || user == nil{
		return nil, err
	}

	var userType proto.UserType

	for _,role := range user.Roles{
		if role == security.Role_COURIER{
			userType = proto.UserType_USER_TYPE_CURIOUR
		} else if role == security.Role_ADMIN{
			userType = proto.UserType_USER_TYPE_ADMIN
		} else if role == security.Role_CLIENT{
			userType = proto.UserType_USER_TYPE_PASSENGER
		}
		break
	}

	validUser,err := s.partyAPI.GetUserByUserId(user.Id, userType)
	if err != nil || validUser == nil{
		return nil, err
	}

	validUser.Roles = user.Roles

	token,err := s.jwtUtils.GenerateToken(*validUser)

	return &pb.RefreshTokenResponse{
		Token: token,
	}, err
}

func (s *TokenService) GetJwks(ctx context.Context, in *pb.GetJwksRequest) (*pb.GetJwksResponse, error){
	return &pb.GetJwksResponse{
		Jwks: []*pb.JwksItem{
			{
				Kid: s.jwtUtils.Kid,
				Alg: s.jwtUtils.Alg,
				Kty: s.jwtUtils.Kty,
				X5C: s.jwtUtils.GetX5c(),
						},
		},
	}, nil
}


func NewTokenService(config config.UaaData, jwtConfig config.JwtData, partyApi PartyAPI) *TokenService {
	jwtUtils, err := security.NewJWTUtils(jwtConfig)
	if err != nil{
		logger.Fatalf("cannot create jwtutils ", err)
	}
	rand.Seed(time.Now().UnixNano())
	return &TokenService{
		config: config,
		jwtUtils: jwtUtils,
		partyAPI: partyApi,
	}
}