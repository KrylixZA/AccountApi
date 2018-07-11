package interfaces

import (
	"../Models/Requests"
)

type IAccountManager interface {
	GetAccountDetails(accountID int) (int, interface{})
	Login(login string, password string) (int, interface{})
	CreateAccount(request requests.CreateAccountRequest) (int, interface{})
	ResetPassword(accountID int, request requests.ResetPasswordRequest) (int, interface{})
}
