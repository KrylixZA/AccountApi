package managers

import (
	"net/http"

	interfaces "../Interfaces"
	responses "../Models/Responses"
)

func (manager *AccountManager) Login(dataAccessor interfaces.IAccountDataAccess, login string, password string) (int, interface{}) {
	account, success := dataAccessor.Login(login, password)

	if success {
		return http.StatusOK, account
	}

	errorResponse := responses.ErrorResponse{Code: 1, Message: "Not found.", Description: "Login failed as no account was found with the given login and password."}
	return http.StatusNotFound, errorResponse
}
