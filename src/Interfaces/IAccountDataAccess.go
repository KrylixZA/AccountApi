package interfaces

import (
	models "../Models"
	requests "../Models/Requests"
)

type IAccountDataAccess interface {
	GetAccountDetails(accountID int) (*models.Account, bool)
	Login(login string, password string) (*models.Account, bool)
	CreateAccount(request requests.CreateAccountRequest) (map[int]models.Account, bool)
	ResetPassword(request requests.ResetPasswordRequest) (*models.Account, bool)
}
