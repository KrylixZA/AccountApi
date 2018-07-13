package dataaccess

import models "../Models"

func (AccountDataAccess) GetAccountDetails(accountID int) (bool, *models.Account) {
	// TODO: Replace this setup with a DB call.
	setupAccounts()

	account, ok := accounts[accountID]
	if ok {
		return true, &account
	}

	return false, nil
}
