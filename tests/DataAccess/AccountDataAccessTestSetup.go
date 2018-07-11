package dataaccess

import dataAccess "../../src/DataAccess"

func getSystemUnderTestAccountDataAccess() *dataAccess.AccountDataAccess {
	return dataAccess.NewAccountDataAccess()
}
