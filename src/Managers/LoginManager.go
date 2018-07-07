package managers

import (
	"encoding/json"
	"net/http"

	interfaces "../Interfaces"
	responses "../Models/Responses"
)

func (manager *AccountManager) Login(dataAccessor interfaces.IAccountDataAccess, login string, password string) (int, []byte) {
	account, success := dataAccessor.Login(login, password)

	if success {
		serializedResponse, err := json.Marshal(account)
		if err != nil {
			panic(err)
		}

		return http.StatusOK, serializedResponse
	}

	errorResponse := responses.ErrorResponse{Code: 1, Message: "Not found.", Description: "Login failed as no account was found with the given login and password."}
	serializedResponse, err := json.Marshal(errorResponse)
	if err != nil {
		panic(err)
	}
	return http.StatusNotFound, serializedResponse
}
