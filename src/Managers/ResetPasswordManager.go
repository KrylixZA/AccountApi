package managers

import (
	"net/http"

	interfaces "../Interfaces"
	requests "../Models/Requests"
	responses "../Models/Responses"
)

func (manager *AccountManager) ResetPassword(dataAccessor interfaces.IAccountDataAccess, accountID int, request requests.ResetPasswordRequest) (int, interface{}) {
	account, success := dataAccessor.ResetPassword(accountID, request)

	if success {
		return http.StatusOK, account
	}

	errorResponse := responses.ErrorResponse{Code: 4, Message: "Not found.", Description: "Failed to reset account password. Could not find account to reset the password for."}
	return http.StatusNotFound, errorResponse
}
