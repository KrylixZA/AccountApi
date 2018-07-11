package managers

import "../Interfaces"

type AccountManager struct {
	DataAccess interfaces.IAccountDataAccess
}

func NewAccountManager(dataAccess interfaces.IAccountDataAccess) *AccountManager {
	return &AccountManager{DataAccess: dataAccess}
}
