package dataaccess

import models "../Models"

func (AccountDataAccess) GetAccountDetails(accountID int) (*models.Account, bool) {
	// TODO: Replace this setup with a DB call.
	setupAccounts()

	account, ok := accounts[accountID]
	if ok {
		return &account, true
	}

	return nil, false
}
