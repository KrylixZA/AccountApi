package controllers

import (
	"encoding/json"
	"net/http"

	dataAccess "../DataAccess"
	manager "../Managers"
	requests "../Models/Requests"
)

func ResetPassword(w http.ResponseWriter, r *http.Request) {
	var request requests.ResetPasswordRequest
	_ = json.NewDecoder(r.Body).Decode(&request)

	// Introduce your strategy
	dataAccess := dataAccess.AccountDataAccess{}
	statusCode, response := manager.ResetPassword(dataAccess, request)
	w.WriteHeader(statusCode)
	w.Write(response)
}
