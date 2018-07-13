package managers

import (
	"net/http"

	"../Diagnostics"
	"../Models/Responses"
)

func (manager *AccountManager) GetAccountDetails(accountID int) (int, interface{}) {
	success, account := manager.DataAccess.GetAccountDetails(accountID)

	if !success {
		errorCode := diagnostics.NoAccountDetailsFound
		errorMessage := "Not found."
		errorDescription := diagnostics.GetErrorDescription(errorCode)
		errorResponse := responses.ErrorResponse{Code: errorCode, Message: errorMessage, Description: errorDescription}
		return http.StatusNotFound, &errorResponse
	}

	return http.StatusOK, account
}
