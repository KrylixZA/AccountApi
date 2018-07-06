package managers

import (
	"encoding/json"
	"net/http"

	interfaces "../Interfaces"
	requests "../Models/Requests"
	responses "../Models/Responses"
)

func ResetPassword(dataAccessor interfaces.IAccountDataAccess, request requests.ResetPasswordRequest) (int, []byte) {
	account, success := dataAccessor.ResetPassword(request)

	if success {
		serializedResponse, err := json.Marshal(account)
		if err != nil {
			panic(err)
		}

		return http.StatusOK, serializedResponse
	}

	errorResponse := responses.ErrorResponse{Code: 4, Message: "Not found.", Description: "Failed to reset account password. Could not find account to reset the password for."}
	serializedResponse, err := json.Marshal(errorResponse)
	if err != nil {
		panic(err)
	}
	return http.StatusNotFound, serializedResponse
}
