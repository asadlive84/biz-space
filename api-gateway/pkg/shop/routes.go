package shop

import (
	"github.com/asadlive84/bizspace/api-gateway/pkg/config"
	"github.com/asadlive84/bizspace/api-gateway/pkg/shop/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/shop")
	routes.POST("/shop-create", svc.CreateShop)

	return svc

}

func (svc *ServiceClient) CreateShop(ctx *gin.Context) {
	routes.CreateShop(ctx, svc.Client)
}
