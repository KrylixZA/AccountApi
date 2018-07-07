package dataaccess

import (
	models "../Models"
	requests "../Models/Requests"
)

func (AccountDataAccess) ResetPassword(request requests.ResetPasswordRequest) (*models.Account, bool) {
	// TODO: Replace this setup with a DB call.
	setupAccounts()

	account, ok := accounts[request.AccountID]
	if ok {
		if account.Password == request.CurrentPassword {
			account.Password = request.NewPassword
			return &account, true
		}

		return nil, false
	}

	return nil, false
}
