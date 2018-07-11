package controllers

import (
	"github.com/gin-gonic/gin"
)

func (controller *AccountController) Login(ctx *gin.Context) {
	login := ctx.Param("login")
	password := ctx.Param("password")

	statusCode, response := controller.AccountManager.Login(login, password)
	ctx.JSON(statusCode, response)
}
