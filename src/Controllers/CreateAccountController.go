package controllers

import (
	"encoding/json"

	"../Models/Requests"
	"github.com/gin-gonic/gin"
)

func (controller *AccountController) CreateAccount(ctx *gin.Context) {
	var request requests.CreateAccountRequest
	_ = json.NewDecoder(ctx.Request.Body).Decode(&request)

	statusCode, response := controller.AccountManager.CreateAccount(request)
	ctx.JSON(statusCode, response)
}
