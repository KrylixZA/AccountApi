package managers

import (
	"net/http"

	"../Models/Requests"
	"../Models/Responses"
)

func (manager *AccountManager) ResetPassword(accountID int, request requests.ResetPasswordRequest) (int, interface{}) {
	account, success := manager.DataAccess.ResetPassword(accountID, request)

	if success {
		return http.StatusOK, account
	}

	errorResponse := &responses.ErrorResponse{Code: 4, Message: "Not found.", Description: "Failed to reset account password. Could not find account to reset the password for."}
	return http.StatusNotFound, errorResponse
}
