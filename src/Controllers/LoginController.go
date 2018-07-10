package controllers

import (
	dataAccess "../DataAccess"
	managers "../Managers"
	"github.com/gin-gonic/gin"
)

func (*AccountController) Login(ctx *gin.Context) {
	login := ctx.Param("login")
	password := ctx.Param("password")

	// Introduce your strategy
	dataAccess := dataAccess.AccountDataAccess{}
	manager := managers.AccountManager{}
	statusCode, response := manager.Login(dataAccess, login, password)
	ctx.JSON(statusCode, response)
}
