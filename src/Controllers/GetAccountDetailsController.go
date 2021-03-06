package controllers

import (
	"net/http"
	"strconv"

	"../Diagnostics"
	"../Models/Responses"
	"github.com/gin-gonic/gin"
)

// GetAccountDetails godoc
// @Summary Exposes functionality to fetch the details of an account once a user has logged in.
// @Description Based off the given account identifier in the URI, the account details are looked up.
// @AccountID accountId
// @Produce json
// @Param accountId path int true "The user's account identifier returned on login."
// @Success 200 {object} models.Account
// @Failure 404 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /accounts/id/{accountId} [get]
func (controller *AccountController) GetAccountDetails(ctx *gin.Context) {
	accountID, error := strconv.Atoi(ctx.Param("accountId"))

	if error != nil {
		errorCode := diagnostics.MalformedAccountID
		errorMessage := "Internal server error."
		errorDescription := diagnostics.GetErrorDescription(errorCode)

		errorResponse := responses.ErrorResponse{Code: errorCode, Message: errorMessage, Description: errorDescription}
		ctx.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	statusCode, response := controller.accountManager.GetAccountDetails(accountID)
	ctx.JSON(statusCode, response)
}
