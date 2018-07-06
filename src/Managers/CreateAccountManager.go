package managers

import (
	"encoding/json"
	"net/http"

	interfaces "../Interfaces"
	requests "../Models/Requests"
	responses "../Models/Responses"
)

func CreateAccount(dataAccessor interfaces.IAccountDataAccess, request requests.CreateAccountRequest) (int, []byte) {
	accounts, success := dataAccessor.CreateAccount(request)

	if success {
		serializedResponse, err := json.Marshal(accounts)
		if err != nil {
			panic(err)
		}

		return http.StatusOK, serializedResponse
	}

	errorResponse := responses.ErrorResponse{Code: 3, Message: "Internal Server Error.", Description: "Failed to create account."}
	serializedResponse, err := json.Marshal(errorResponse)
	if err != nil {
		panic(err)
	}
	return http.StatusInternalServerError, serializedResponse
}
