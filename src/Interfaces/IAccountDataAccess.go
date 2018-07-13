package interfaces

import (
	"../Models"
	"../Models/Requests"
)

type IAccountDataAccess interface {
	GetAccountDetails(accountID int) (bool, *models.Account)
	Login(login string, password string) (bool, *models.Account)
	CreateAccount(request *requests.CreateAccountRequest) (bool, *models.Account)
	ResetPassword(accountID int, request *requests.ResetPasswordRequest) (bool, *models.Account)
}
