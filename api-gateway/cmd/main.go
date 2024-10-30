package main

import (
	"log"
	"net/http"

	"github.com/asadlive84/bizspace/api-gateway/pkg/auth"
	"github.com/asadlive84/bizspace/api-gateway/pkg/shop"
	"github.com/asadlive84/bizspace/api-gateway/pkg/config"
	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("###Failed at config", err)
	}

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	})

	// authSvc := *auth.RegisterRoutes(r, &c)
	auth.RegisterRoutes(r, &c)
	shop.RegisterRoutes(r, &c)

	// Run the server
	r.Run(c.Port)

}
