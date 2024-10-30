package routes

import (
	"context"
	"net/http"

	pb "github.com/asadlive84/bizspace/proto/shop/pb" 
	"github.com/gin-gonic/gin"
)

type CreateShopRequestBody struct {
	ShopName string `json:"shop_name"`
	OwnerID  string `json:"owner_id"`
}

func CreateShop(ctx *gin.Context, d pb.ShopServiceClient) {
	body := CreateShopRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := d.CreateShop(context.Background(), &pb.CreateShopRequest{
		Name:    body.ShopName,
		OwnerId: body.OwnerID,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
