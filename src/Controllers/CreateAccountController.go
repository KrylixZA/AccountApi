package controllers

import (
	"encoding/json"
	"net/http"

	"../Models/Requests"
	"github.com/gin-gonic/gin"
)

// CreateAccount godoc
// @Summary Exposes functionality to create a new account.
// @Description Creates an account for a new user if the account doesn't already exist.
// @Accept json
// @Produce json
// @Param createAccountRequest body requests.CreateAccountRequest true "The payload require to create the account."
// @Success 201 {object} models.Account
// @Failure 400 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /accounts [post]
func (controller *AccountController) CreateAccount(ctx *gin.Context) {
	var request requests.CreateAccountRequest
	_ = json.NewDecoder(ctx.Request.Body).Decode(&request)

	validRequest, errorResponse := controller.requestValidation.ValidateCreateAccountRequest(&request)
	if !validRequest {
		statusCode := http.StatusBadRequest
		ctx.JSON(statusCode, errorResponse)
		return
	}

	statusCode, response := controller.accountManager.CreateAccount(&request)
	ctx.JSON(statusCode, response)
}
