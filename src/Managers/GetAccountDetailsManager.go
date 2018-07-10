package managers

import (
	"net/http"

	interfaces "../Interfaces"
	responses "../Models/Responses"
)

func (manager *AccountManager) GetAccountDetails(dataAccessor interfaces.IAccountDataAccess, accountID int) (int, interface{}) {
	account, success := dataAccessor.GetAccountDetails(accountID)

	if success {
		return http.StatusOK, account
	}

	errorResponse := responses.ErrorResponse{Code: 1, Message: "Not found.", Description: "No accounts were found with the given identifier and hence no details could be returned."}
	return http.StatusNotFound, errorResponse
}
