package controllers

import (
	"encoding/json"

	dataAccess "../DataAccess"
	managers "../Managers"
	requests "../Models/Requests"
	"github.com/gin-gonic/gin"
)

func (*AccountController) CreateAccount(ctx *gin.Context) {
	var request requests.CreateAccountRequest
	_ = json.NewDecoder(ctx.Request.Body).Decode(&request)

	// Introduce your strategy
	dataAccess := dataAccess.AccountDataAccess{}
	manager := managers.AccountManager{}
	statusCode, response := manager.CreateAccount(dataAccess, request)
	ctx.JSON(statusCode, response)
}
