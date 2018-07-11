package mocks

import (
	"../../../src/Models"
	"../../../src/Models/Requests"
)

type AccountDataAccessMock struct {
	GetAccountDetailsFunc func(accountID int) (*models.Account, bool)
	LoginFunc             func(login string, password string) (*models.Account, bool)
	CreateAccountFunc     func(request requests.CreateAccountRequest) (*models.Account, bool)
	ResetPasswordFunc     func(request requests.ResetPasswordRequest) (*models.Account, bool)
}

func (mock AccountDataAccessMock) GetAccountDetails(accountID int) (*models.Account, bool) {
	return mock.GetAccountDetailsFunc(accountID)
}

func (mock AccountDataAccessMock) Login(login string, password string) (*models.Account, bool) {
	return mock.LoginFunc(login, password)
}

func (mock AccountDataAccessMock) CreateAccount(request requests.CreateAccountRequest) (*models.Account, bool) {
	return mock.CreateAccountFunc(request)
}

func (mock AccountDataAccessMock) ResetPassword(accountID int, request requests.ResetPasswordRequest) (*models.Account, bool) {
	return mock.ResetPasswordFunc(request)
}
