package interfaces

import (
	"../Models"
	"../Models/Requests"
)

type IAccountDataAccess interface {
	GetAccountDetails(accountID int) (*models.Account, bool)
	Login(login string, password string) (*models.Account, bool)
	CreateAccount(request *requests.CreateAccountRequest) (*models.Account, bool)
	ResetPassword(accountID int, request *requests.ResetPasswordRequest) (*models.Account, bool)
}
