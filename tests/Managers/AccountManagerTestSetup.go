package managers

import (
	"../../src/Interfaces"
	managers "../../src/Managers"
)

func getSystemUnderTestAccountManager(dataAccess interfaces.IAccountDataAccess) *managers.AccountManager {
	return managers.NewAccountManager(dataAccess)
}
