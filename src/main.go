// @title Account API
// @version 1.0.0
// @description Provides basic functionality to perform operations against accounts.

// @contact.name Simon Headley
// @contact.url None
// @contact.email headleysj@gmail.com

// @license.name None
// @license.url None

// @host localhost:8080
// @basePath /api/v1

package main

import (
	"../src/Controllers"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"../src/DataAccess"
	"../src/Managers"
)

func main() {
	r := gin.Default()

	// TO UPDATE SWAGGO: /Users/headleysj/go/bin/swag

	// introduce your strategy here.
	dataAccess := dataaccess.NewAccountDataAccess()
	manager := managers.NewAccountManager(dataAccess)
	accountController := controllers.NewController(manager)
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
