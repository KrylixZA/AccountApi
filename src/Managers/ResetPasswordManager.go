package managers

import (
	"net/http"

	"../Diagnostics"
	"../Models/Requests"
	"../Models/Responses"
)

func (manager *AccountManager) ResetPassword(accountID int, request *requests.ResetPasswordRequest) (int, interface{}) {
	account, success := manager.DataAccess.ResetPassword(accountID, request)

	if !success {
		errorCode := diagnostics.ResetPasswordFailed
		errorMessage := "Not found."
		errorDescription := diagnostics.GetErrorDescription(errorCode)
		errorResponse := responses.ErrorResponse{Code: errorCode, Message: errorMessage, Description: errorDescription}
		return http.StatusNotFound, errorResponse
	}

	return http.StatusOK, account
}
