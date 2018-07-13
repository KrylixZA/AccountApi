package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"../Models/Requests"
	"../Models/Responses"
	"github.com/gin-gonic/gin"
)

// ResetPassword godoc
// @Summary Exposes functionality to reset a user's password.
// @Description Based on the account identifier and the current password, the system will attempt to update the user's password.
// @AccountID accountId
// @Param accountId path string true "The user's account identifier."
// @Param resetPasswordRequest body requests.ResetPasswordRequest true "The payload containing the existing password and the new password to be updated."
// @Success 200 {object} models.Account
// @Failure 400 {object} responses.ErrorResponse
// @Failure 404 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /accounts/id/{accountId} [patch]
func (controller *AccountController) ResetPassword(ctx *gin.Context) {
	accountID, error := strconv.Atoi(ctx.Param("accountId"))

	if error != nil {
		errorResponse := responses.ErrorResponse{Code: 2, Message: "Internal Server Error.", Description: "Failed to convert the parameter to an integer."}
		ctx.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	var request requests.ResetPasswordRequest
	_ = json.NewDecoder(ctx.Request.Body).Decode(&request)

	validRequest, errorResponse := controller.requestValidation.ValidateResetPasswordRequest(&request)
	if !validRequest {
		statusCode := http.StatusBadRequest
		ctx.JSON(statusCode, errorResponse)
		return
	}

	statusCode, response := controller.accountManager.ResetPassword(accountID, &request)
	ctx.JSON(statusCode, response)
}
