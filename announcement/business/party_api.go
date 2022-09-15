package business

import (
	"context"
	"gitlab.artin.ai/backend/courier-management/common/config"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	"gitlab.artin.ai/backend/courier-management/party/proto"
	"gitlab.artin.ai/backend/courier-management/uaa/security"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"time"
)

type PartyAPI struct{
	config    config.UaaData
}


func (api PartyAPI) GetUserByUserId(userId string, userType proto.UserType) (*security.User,error) {
	logger.Debugf("findUser by userId %v", userId)
	query := &proto.FindAccountRequest{
		Filter: &proto.FindAccountRequest_UserId{
			UserId: userId,
		},
		Type: userType,
	}
	users, err := api.findUser(query)
	logger.Debugf("found users are %v", users)
	if users!=nil && len(users) > 0{
		return users[0], nil
	}
	return nil, err
}


func (api PartyAPI) findUser(query *proto.FindAccountRequest) ([]*security.User,error) {
	conn := api.getConn()
	defer conn.Close()
	clientDeadline := time.Now().Add(time.Duration(6000) * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), clientDeadline)
	defer cancel()

	c := proto.NewUaaServiceClient(conn)

	out, err := c.FindAccount(ctx, query)

	if err != nil {
		if e, ok := status.FromError(err); ok {
			if e.Code() == 5{
				return nil, nil
			}
		}
		logger.Errorf("cannot find users from party", err)
		return nil, err
	}
	if len(out.Users) == 0{
		return nil, nil
	}

	var users = make([]*security.User, len(out.Users))
	for i,a := range out.Users{
		var claims []security.Claim
		for _, claim := range a.Claims{
			claims = append(claims, security.Claim{
				ClaimType: security.ClaimType(claim.ClaimType),
				Identifier: claim.Identifier,
			})
		}
		users[i] = &security.User{
			Id: a.UserId,
			Name: a.FirstName + " " + a.LastName,
			Email: a.Email,
			PhoneNumber: a.PhoneNumber,
			Claims: claims,
		}
	}
	return users, nil
}

func (api PartyAPI) getConn() *grpc.ClientConn{
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	opts = append(opts, grpc.WithBlock())
	conn, err := grpc.Dial(api.config.PartyAddress, opts...)
	if err != nil {
		logger.Errorf("cannot connect to party: %v", err)
	}
	return conn
}

func NewPartyAPI(config config.UaaData) PartyAPI{
	return PartyAPI{
		config: config,
	}
}

type PartyError struct{
	Code int32
	Error string
	Desc string
}