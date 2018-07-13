package dataaccess

import "../Models"

func (AccountDataAccess) Login(login string, password string) (bool, *models.Account) {
	// TODO: Replace this setup with a DB call.
	setupAccounts()

	for _, item := range accounts {
		if item.Login == login && item.Password == password {
			return true, &item
		}
	}

	return false, nil
}
