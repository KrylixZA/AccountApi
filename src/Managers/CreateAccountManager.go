package managers

import (
	"encoding/json"
	"net/http"

	interfaces "../Interfaces"
	requests "../Models/Requests"
	responses "../Models/Responses"
)

func (manager *AccountManager) CreateAccount(dataAccessor interfaces.IAccountDataAccess, request requests.CreateAccountRequest) (int, []byte) {
	accounts, success := dataAccessor.CreateAccount(request)

	if success {
		serializedResponse, err := json.Marshal(accounts)
		if err != nil {
			panic(err)
		}

		return http.StatusCreated, serializedResponse
	}

	errorResponse := responses.ErrorResponse{Code: 3, Message: "Internal Server Error.", Description: "Failed to create account."}
	serializedResponse, err := json.Marshal(errorResponse)
	if err != nil {
		panic(err)
	}
	return http.StatusInternalServerError, serializedResponse
}
