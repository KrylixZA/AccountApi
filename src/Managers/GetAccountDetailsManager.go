package managers

import (
	"encoding/json"
	"net/http"

	interfaces "../Interfaces"
	responses "../Models/Responses"
)

func (manager *AccountManager) GetAccountDetails(dataAccessor interfaces.IAccountDataAccess, accountID int) (int, []byte) {
	account, success := dataAccessor.GetAccountDetails(accountID)

	if success {
		serializedResponse, err := json.Marshal(account)
		if err != nil {
			panic(err)
		}

		return http.StatusOK, serializedResponse
	}

	errorResponse := responses.ErrorResponse{Code: 1, Message: "Not found.", Description: "No accounts were found with the given identifier and hence no details could be returned."}
	serializedResponse, err := json.Marshal(errorResponse)
	if err != nil {
		panic(err)
	}
	return http.StatusNotFound, serializedResponse
}
