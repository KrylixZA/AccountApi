package interfaces

import (
	"../Models/Requests"
	"../Models/Responses"
)

type IValidation interface {
	ValidateCreateAccountRequest(request *requests.CreateAccountRequest) (bool, *responses.ErrorResponse)
	ValidateResetPasswordRequest(request *requests.ResetPasswordRequest) (bool, *responses.ErrorResponse)
	ValidateLoginAndPassword(login string, password string) (bool, *responses.ErrorResponse)
}
