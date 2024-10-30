package routes

import (
	"context"
	"net/http"

	pb "github.com/asadlive84/bizspace/proto/auth/pb"
	"github.com/gin-gonic/gin"
)

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

func Login(ctx *gin.Context, d pb.AuthServiceClient) {
	body := LoginRequestBody{}
	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := d.Login(context.Background(), &pb.LoginRequest{
		Email:    body.Email,
		Password: body.Password,
		Phone:    body.Phone,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)

}
