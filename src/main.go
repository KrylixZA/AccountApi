package main

import (
	"log"
	"net/http"

	ctrls "../src/Controllers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Account controller endpoints
	router.HandleFunc("/account/{id}", ctrls.GetAccountDetails).Methods("GET")
	router.HandleFunc("/account/{login}/{password}", ctrls.Login).Methods("GET")
	router.HandleFunc("/account", ctrls.CreateAccount).Methods("POST")
	router.HandleFunc("/account", ctrls.ResetPassword).Methods("PATCH")

	log.Fatal(http.ListenAndServe(":8000", router))
}
