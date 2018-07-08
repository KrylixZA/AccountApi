package dataaccess

import (
	models "../Models"
	requests "../Models/Requests"
)

func (AccountDataAccess) CreateAccount(request requests.CreateAccountRequest) (map[int]models.Account, bool) {
	// TODO: Replace this setup with a DB call.
	setupAccounts()

	newIndex := len(accounts) + 1
	newAccount := models.Account{AccountID: newIndex, Login: request.Email, Password: request.Password, FirstName: request.FirstName, Surname: request.Surname, Email: request.Email}

	accounts[newIndex] = newAccount

	return accounts, true
}
