package controllers

import (
	"net/http"

	dataAccess "../DataAccess"
	manager "../Managers"
	"github.com/gorilla/mux"
)

func Login(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	login := params["login"]
	password := params["password"]

	// Introduce your strategy
	dataAccess := dataAccess.AccountDataAccess{}
	statusCode, response := manager.Login(dataAccess, login, password)
	w.WriteHeader(statusCode)
	w.Write(response)
}
