package auth

import (
	"fmt"

	// "github.com/asadlive84/bizspace/api-gateway/pkg/auth/pb"
	"github.com/asadlive84/bizspace/api-gateway/pkg/config"
	pb "github.com/asadlive84/bizspace/proto/auth/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServiceClient(c *config.Config) pb.AuthServiceClient {
	cc, err := grpc.Dial(c.AuthSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Could not connect:", err)
	}
	return pb.NewAuthServiceClient(cc)

}
