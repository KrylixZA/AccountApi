package dataaccess

import (
	"../Models"
	"../Models/Requests"
)

func (AccountDataAccess) ResetPassword(accountID int, request requests.ResetPasswordRequest) (*models.Account, bool) {
	// TODO: Replace this setup with a DB call.
	setupAccounts()

	account, ok := accounts[accountID]
	if ok {
		if account.Password == request.CurrentPassword {
			account.Password = request.NewPassword
			return &account, true
		}

		return nil, false
	}

	return nil, false
}
