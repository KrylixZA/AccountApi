package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	dataAccess "../DataAccess"
	manager "../Managers"
	responses "../Models/Responses"
	"github.com/gorilla/mux"
)

func GetAccountDetails(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	accountID, error := strconv.Atoi(params["id"])

	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(responses.ErrorResponse{Code: 2, Message: "Internal Server Error.", Description: "Failed to convert the parameter to an integer."})
		return
	}

	// Introduce your strategy
	dataAccess := dataAccess.AccountDataAccess{}
	statusCode, response := manager.GetAccountDetails(dataAccess, accountID)
	w.WriteHeader(statusCode)
	w.Write(response)
}
