package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"../Models/Requests"
	"../Models/Responses"
	"github.com/gin-gonic/gin"
)

func (controller *AccountController) ResetPassword(ctx *gin.Context) {
	accountID, error := strconv.Atoi(ctx.Param("accountId"))

	if error != nil {
		errorResponse := responses.ErrorResponse{Code: 2, Message: "Internal Server Error.", Description: "Failed to convert the parameter to an integer."}
		ctx.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	var request requests.ResetPasswordRequest
	_ = json.NewDecoder(ctx.Request.Body).Decode(&request)

	statusCode, response := controller.AccountManager.ResetPassword(accountID, request)
	ctx.JSON(statusCode, response)
}
