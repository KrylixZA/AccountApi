package mocks

import (
	"../../../src/Models"
	"../../../src/Models/Requests"
)

type AccountDataAccessMock struct {
	GetAccountDetailsFunc func(accountID int) (bool, *models.Account)
	LoginFunc             func(login string, password string) (bool, *models.Account)
	CreateAccountFunc     func(request *requests.CreateAccountRequest) (bool, *models.Account)
	ResetPasswordFunc     func(request *requests.ResetPasswordRequest) (bool, *models.Account)
}

func (mock AccountDataAccessMock) GetAccountDetails(accountID int) (bool, *models.Account) {
	return mock.GetAccountDetailsFunc(accountID)
}

func (mock AccountDataAccessMock) Login(login string, password string) (bool, *models.Account) {
	return mock.LoginFunc(login, password)
}

func (mock AccountDataAccessMock) CreateAccount(request *requests.CreateAccountRequest) (bool, *models.Account) {
	return mock.CreateAccountFunc(request)
}

func (mock AccountDataAccessMock) ResetPassword(accountID int, request *requests.ResetPasswordRequest) (bool, *models.Account) {
	return mock.ResetPasswordFunc(request)
}
