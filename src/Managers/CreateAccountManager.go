package managers

import (
	"net/http"

	interfaces "../Interfaces"
	requests "../Models/Requests"
	responses "../Models/Responses"
)

func (manager *AccountManager) CreateAccount(dataAccessor interfaces.IAccountDataAccess, request requests.CreateAccountRequest) (int, interface{}) {
	accounts, success := dataAccessor.CreateAccount(request)

	if success {
		return http.StatusCreated, accounts
	}

	errorResponse := responses.ErrorResponse{Code: 3, Message: "Internal Server Error.", Description: "Failed to create account."}
	return http.StatusInternalServerError, errorResponse
}
