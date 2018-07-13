package mocks

import "../../../src/Models/Requests"

type AccountManagerMock struct {
	GetAccountDetailsFunc func(accountID int) (int, interface{})
	LoginFunc             func(login string, password string) (int, interface{})
	CreateAccountFunc     func(request *requests.CreateAccountRequest) (int, interface{})
	ResetPasswordFunc     func(accountID int, request *requests.ResetPasswordRequest) (int, interface{})
}

func (mock AccountManagerMock) GetAccountDetails(accountID int) (int, interface{}) {
	return mock.GetAccountDetailsFunc(accountID)
}

func (mock AccountManagerMock) Login(login string, password string) (int, interface{}) {
	return mock.LoginFunc(login, password)
}

func (mock AccountManagerMock) CreateAccount(request *requests.CreateAccountRequest) (int, interface{}) {
	return mock.CreateAccountFunc(request)
}

func (mock AccountManagerMock) ResetPassword(accountID int, request *requests.ResetPasswordRequest) (int, interface{}) {
	return mock.ResetPassword(accountID, request)
}
