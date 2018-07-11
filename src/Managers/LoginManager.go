package managers

import (
	"net/http"

	responses "../Models/Responses"
)

func (manager *AccountManager) Login(login string, password string) (int, interface{}) {
	account, success := manager.DataAccess.Login(login, password)

	if success {
		return http.StatusOK, account
	}

	errorResponse := responses.ErrorResponse{Code: 1, Message: "Not found.", Description: "Login failed as no account was found with the given login and password."}
	return http.StatusNotFound, errorResponse
}
