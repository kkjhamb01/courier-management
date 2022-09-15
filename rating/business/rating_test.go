package business

import (
	"context"
	"gitlab.artin.ai/backend/courier-management/common/config"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	pb "gitlab.artin.ai/backend/courier-management/grpc/offering/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"testing"
	"time"
)

func init(){
	config.InitTestConfig()
	logger.InitLogger()
}

func TestGetRideInfo(t *testing.T) {
	request := &pb.GetOfferCourierAndCustomerRequest{
		OfferId: "2429df86-2b52-477a-8afb-23e93f738f41",
	}

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	opts = append(opts, grpc.WithBlock())
	conn, err := grpc.Dial("185.97.117.71:50003", opts...)
	if err != nil {
		logger.Errorf("cannot connect to offering: %v", err)
	}

	defer conn.Close()

	clientDeadline := time.Now().Add(time.Duration(6000) * time.Millisecond)

	var header = metadata.Pairs("access_token", "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJjb3VyaWVyLW1hbmFnZW1lbnQtdWFhIiwiYXV0aG9yaXplZCI6W3siY2xhaW1fdHlwZSI6MiwiaWRlbnRpZmllciI6Iis0NDE1OTM1NzQ1OTIifV0sImRldmljZSI6ImQxIiwiZXhwIjoxNjI0OTg3ODA1LCJpYXQiOjE2MjQ5MDE0MDUsImlkIjoiY2I3MTE2OTMtMjdiYi00MzZhLTlmNTQtZTM2YjY5ZGU5MzA1IiwiaXNzIjoiaHR0cDovL29rZC5yYXBpZHJvcHMuY29tIiwia2lkIjoiMjAyMWp3dCIsInBob25lX251bWJlciI6Iis0NDE1OTM1NzQ1OTIiLCJyb2xlcyI6WyJDT1VSSUVSIl0sInN1YiI6IjUyMTQwNjBlLTBkM2QtMjU5Yy0wYmUwLTVlZGNkNmQ5ZTY0ZSIsInR5cCI6ImJlYXJlciJ9.KypwnSnMqy27sVasZQnNjmNNw1UC-szyjMC0eba2RtT-EUfwApCrBbrPBW2cuJ-tfn0JlD6fHDGlLOSc8LZ-gsgBOm2Zy6nXTqrbfykZzv8nPsLP9RBIslg3y9W7D64PgrAWrG3wF1OpZTkeUJaWHi7V8Y2eam4DrEuC4axuc_L-Jk8XjSBp_B84hjU6WIn2cRYsdm-sOuYDkLCAbF4g9AtbsgNW14YEhXqrMoqZQzu6_9Gi6Xj-RZmlIFeJRo53oZdJKIAPHWFxVK2kwGTvGIZF3jQWDO1c3aRusF3d1B85ukadKRK5dygWWFvsVQcgEzJgUA2VsD0PFXCSZBxULQ")
	ctx := metadata.NewOutgoingContext(context.Background(), header)

	ctx, cancel := context.WithDeadline(ctx, clientDeadline)
	defer cancel()

	c := pb.NewOfferingClient(conn)

	out, err := c.GetOfferCourierAndCustomer(ctx, request)

	if e, ok := status.FromError(err); ok {
		if e.Code() == 2{
			logger.Debugf("not found")
		}
	}

	if err != nil{
		logger.Debugf("error = %v", err.Error())
	}

	logger.Debugf("out = %v", out)
}