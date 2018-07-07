package mocks

import (
	models "../../../src/Models"
	requests "../../../src/Models/Requests"
)

type AccountDataAccessMock struct {
	GetAccountDetailsFunc func(accountID int) (*models.Account, bool)
	LoginFunc             func(login string, password string) (*models.Account, bool)
	CreateAccountFunc     func(request requests.CreateAccountRequest) (map[int]models.Account, bool)
	ResetPasswordFunc     func(request requests.ResetPasswordRequest) (*models.Account, bool)
}

func (mock AccountDataAccessMock) GetAccountDetails(accountID int) (*models.Account, bool) {
	return mock.GetAccountDetailsFunc(accountID)
}

func (mock AccountDataAccessMock) Login(login string, password string) (*models.Account, bool) {
	return mock.Login(login, password)
}

func (mock AccountDataAccessMock) CreateAccount(request requests.CreateAccountRequest) (map[int]models.Account, bool) {
	return mock.CreateAccount(request)
}

func (mock AccountDataAccessMock) ResetPassword(request requests.ResetPasswordRequest) (*models.Account, bool) {
	return mock.ResetPassword(request)
}
