package routes

import (
	"context"
	"net/http"

	pb "github.com/asadlive84/bizspace/proto/auth/pb"
	"github.com/gin-gonic/gin"
)

type CreateUserRequestBody struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	UserName    string `json:"user_name,omitempty"`
	FullName    string `json:"full_name"`
	Address     string `json:"address,omitempty"`
}

func Register(ctx *gin.Context, d pb.AuthServiceClient) {

	body := CreateUserRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := d.CreateUser(context.Background(), &pb.CreateUserRequest{
		UserName:     body.UserName,
		FullName:     body.FullName,
		PhoneNumber:  body.PhoneNumber,
		Address:      body.Address,
		PasswordHash: body.Password,
		Email:        body.Email,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(int(res.Status), &res)

}
