// /Users/headleysj/go/bin/swag

// @title Account API
// @version 1.0.0
// @description Provides basic functionality to perform operations against accounts.

// @host localhost
// @basePath /api/v1

package main

import (
	ctrls "../src/Controllers"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func main() {
	r := gin.Default()

	accountController := ctrls.NewController()

	v1 := r.Group("/api/v1")
	{
		accounts := v1.Group("/accounts")
		{
			accounts.GET("/id/:accountId", accountController.GetAccountDetails)
			accounts.GET("/login/:login/password/:password", accountController.Login)
			accounts.POST("", accountController.CreateAccount)
			accounts.PATCH("/id/:accountId", accountController.ResetPassword)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}
