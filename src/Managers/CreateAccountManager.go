package managers

import (
	"net/http"

	"../Models/Requests"
	"../Models/Responses"
)

func (manager *AccountManager) CreateAccount(request requests.CreateAccountRequest) (int, interface{}) {
	account, success := manager.DataAccess.CreateAccount(request)

	if success {
		return http.StatusCreated, account
	}

	errorResponse := &responses.ErrorResponse{Code: 3, Message: "Internal Server Error.", Description: "Failed to create account."}
	return http.StatusInternalServerError, errorResponse
}
