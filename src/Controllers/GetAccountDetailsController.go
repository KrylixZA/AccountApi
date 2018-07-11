package controllers

import (
	"net/http"
	"strconv"

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
		errorResponse := responses.ErrorResponse{Code: 2, Message: "Internal Server Error.", Description: "Failed to convert the parameter to an integer."}
		ctx.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	statusCode, response := controller.AccountManager.GetAccountDetails(accountID)
	ctx.JSON(statusCode, response)
}
