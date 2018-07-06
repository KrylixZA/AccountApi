package dataaccess

import models "../Models"

func (AccountDataAccess) Login(login string, password string) (*models.Account, bool) {
	// TODO: Replace this setup with a DB call.
	setupAccounts()

	for _, item := range accounts {
		if item.Login == login && item.Password == password {
			return &item, true
		}
	}

	return nil, false
}
