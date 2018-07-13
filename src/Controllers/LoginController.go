package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login godoc
// @Summary Exposes functionality to fetch the details of a user account when attempting to log in.
// @Description Based off the login and password supplied, the system will attempt to fetch the details of a user's account.
// @Login login
// @Password password
// @Param login path string true "The user's login."
// @Param password path string true "The user's password that belongs to the login."
// @Success 200 {object} models.Account
// @Failure 400 {object} responses.ErrorResponse
// @Failure 404 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /accounts/login/{login}/password/{password} [get]
func (controller *AccountController) Login(ctx *gin.Context) {
	login := ctx.Param("login")
	password := ctx.Param("password")

	validRequest, errorResponse := controller.requestValidation.ValidateLoginAndPassword(login, password)
	if !validRequest {
		statusCode := http.StatusBadRequest
		ctx.JSON(statusCode, errorResponse)
		return
	}

	statusCode, response := controller.accountManager.Login(login, password)
	ctx.JSON(statusCode, response)
}
