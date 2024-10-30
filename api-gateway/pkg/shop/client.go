package shop

import (
	"fmt"

	// "github.com/asadlive84/bizspace/api-gateway/pkg/auth/pb"
	"github.com/asadlive84/bizspace/api-gateway/pkg/config"
	pb "github.com/asadlive84/bizspace/proto/shop/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.ShopServiceClient
}

func InitServiceClient(c *config.Config) pb.ShopServiceClient {
	cc, err := grpc.Dial(c.ShopSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Could not connect:", err)
	}
	return pb.NewShopServiceClient(cc)

}
