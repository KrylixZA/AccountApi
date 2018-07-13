package managers

import (
	"net/http"

	"../Diagnostics"
	"../Models/Requests"
	"../Models/Responses"
)

func (manager *AccountManager) CreateAccount(request *requests.CreateAccountRequest) (int, interface{}) {
	success, account := manager.DataAccess.CreateAccount(request)

	if !success {
		errorCode := diagnostics.FailedToCreateAccount
		errorMessage := "Internal server error."
		errorDescription := diagnostics.GetErrorDescription(errorCode)
		errorResponse := responses.ErrorResponse{Code: errorCode, Message: errorMessage, Description: errorDescription}
		return http.StatusInternalServerError, &errorResponse
	}

	return http.StatusCreated, account
}
