package auth

import (
	// "context"
	"net/http"
	"strings"

	// pb "github.com/asadlive84/bizspace/proto/auth/pb"
	"github.com/gin-gonic/gin"
)

type AuthMiddlewareConfig struct {
	svc *ServiceClient
}

func InitAuthMiddlewareServiceClient(svc *ServiceClient) AuthMiddlewareConfig {
	return AuthMiddlewareConfig{svc}
}

func (a *AuthMiddlewareConfig) AuthRequired(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("authorization")

	if authorization == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token := strings.Split(authorization, "Bearer ")

	if len(token) < 2 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// res, err := a.svc.Client.Validate(context.TODO(), &pb.ValidateRequest{
	// 	Token: token[1],
	// })

	// if err != nil {
	// 	ctx.AbortWithStatus(http.StatusUnauthorized)
	// 	return
	// }

	// ctx.Set("userId", res.UserId)

	// ctx.Next()

}
