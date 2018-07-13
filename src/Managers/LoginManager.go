package managers

import (
	"net/http"

	"../Diagnostics"
	"../Models/Responses"
)

func (manager *AccountManager) Login(login string, password string) (int, interface{}) {
	account, success := manager.DataAccess.Login(login, password)

	if !success {
		errorCode := diagnostics.LoginFailed
		errorMessage := "Not found."
		errorDescription := diagnostics.GetErrorDescription(errorCode)
		errorResponse := responses.ErrorResponse{Code: errorCode, Message: errorMessage, Description: errorDescription}
		return http.StatusNotFound, errorResponse
	}

	return http.StatusOK, account
}
