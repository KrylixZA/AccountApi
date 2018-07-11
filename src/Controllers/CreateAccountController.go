package controllers

import (
	"encoding/json"

	"../Models/Requests"
	"github.com/gin-gonic/gin"
)

// CreateAccount godoc
// @Summary Exposes functionality to create a new account.
// @Description Creates an account for a new user if the account doesn't already exist.
// @Accept json
// @Produce json
// @Success 201 {object} models.Account
// @Failure 500 {object} responses.ErrorResponse
// @Router /accounts [post]
func (controller *AccountController) CreateAccount(ctx *gin.Context) {
	var request requests.CreateAccountRequest
	_ = json.NewDecoder(ctx.Request.Body).Decode(&request)

	statusCode, response := controller.AccountManager.CreateAccount(request)
	ctx.JSON(statusCode, response)
}
