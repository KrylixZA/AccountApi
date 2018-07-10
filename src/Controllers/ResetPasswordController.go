package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	dataAccess "../DataAccess"
	managers "../Managers"
	requests "../Models/Requests"
	responses "../Models/Responses"
	"github.com/gin-gonic/gin"
)

func (*AccountController) ResetPassword(ctx *gin.Context) {
	accountID, error := strconv.Atoi(ctx.Param("accountId"))

	if error != nil {
		errorResponse := responses.ErrorResponse{Code: 2, Message: "Internal Server Error.", Description: "Failed to convert the parameter to an integer."}
		ctx.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	var request requests.ResetPasswordRequest
	_ = json.NewDecoder(ctx.Request.Body).Decode(&request)

	// Introduce your strategy
	dataAccess := dataAccess.AccountDataAccess{}
	manager := managers.AccountManager{}
	statusCode, response := manager.ResetPassword(dataAccess, accountID, request)
	ctx.JSON(statusCode, response)
}
