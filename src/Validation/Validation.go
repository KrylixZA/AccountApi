package validation

import (
	"../Diagnostics"
	"../Models/Requests"
	"../Models/Responses"
)

type Validation struct {
}

func NewValidation() *Validation {
	return &Validation{}
}

func (*Validation) ValidateCreateAccountRequest(request *requests.CreateAccountRequest) (bool, *responses.ErrorResponse) {
	if request == nil {
		errorCode := diagnostics.BadRequest
		errorMessage := "Bad request."
		errorDescription := diagnostics.GetErrorDescription(errorCode)
		errorResponse := &responses.ErrorResponse{Code: errorCode, Message: errorMessage, Description: errorDescription}

		return false, errorResponse
	}

	if request.Username == "" {
		errorCode := diagnostics.BadUsername
		errorMessage := "Bad username."
		errorDescription := diagnostics.GetErrorDescription(errorCode)
		errorResponse := &responses.ErrorResponse{Code: errorCode, Message: errorMessage, Description: errorDescription}

		return false, errorResponse
	}

	if request.Email == "" {
		errorCode := diagnostics.BadEmail
		errorMessage := "Bad email."
		errorDescription := diagnostics.GetErrorDescription(errorCode)
		errorResponse := &responses.ErrorResponse{Code: errorCode, Message: errorMessage, Description: errorDescription}

		return false, errorResponse
	}

	if request.FirstName == "" {
		errorCode := diagnostics.BadFirstName
		errorMessage := "Bad first name."
		errorDescription := diagnostics.GetErrorDescription(errorCode)
		errorResponse := &responses.ErrorResponse{Code: errorCode, Message: errorMessage, Description: errorDescription}

		return false, errorResponse
	}

	if request.Surname == "" {
		errorCode := diagnostics.BadSurname
		errorMessage := "Bad surname."
		errorDescription := diagnostics.GetErrorDescription(errorCode)
		errorResponse := &responses.ErrorResponse{Code: errorCode, Message: errorMessage, Description: errorDescription}

		return false, errorResponse
	}

	if request.Password == "" {
		errorCode := diagnostics.BadPassword
		errorMessage := "Bad password."
		errorDescription := diagnostics.GetErrorDescription(errorCode)
		errorResponse := &responses.ErrorResponse{Code: errorCode, Message: errorMessage, Description: errorDescription}

		return false, errorResponse
	}

	return true, nil
}

func (*Validation) ValidateResetPasswordRequest(request *requests.ResetPasswordRequest) (bool, *responses.ErrorResponse) {
	if request == nil {
		errorCode := diagnostics.BadRequest
		errorMessage := "Bad request."
		errorDescription := diagnostics.GetErrorDescription(errorCode)
		errorResponse := &responses.ErrorResponse{Code: errorCode, Message: errorMessage, Description: errorDescription}

		return false, errorResponse
	}

	if request.CurrentPassword == "" {
		errorCode := diagnostics.BadCurrentPassword
		errorMessage := "Bad current password."
		errorDescription := diagnostics.GetErrorDescription(errorCode)
		errorResponse := &responses.ErrorResponse{Code: errorCode, Message: errorMessage, Description: errorDescription}

		return false, errorResponse
	}

	if request.NewPassword == "" {
		errorCode := diagnostics.BadNewPassword
		errorMessage := "Bad new password."
		errorDescription := diagnostics.GetErrorDescription(errorCode)
		errorResponse := &responses.ErrorResponse{Code: errorCode, Message: errorMessage, Description: errorDescription}

		return false, errorResponse
	}

	return true, nil
}

func (*Validation) ValidateLoginAndPassword(login string, password string) (bool, *responses.ErrorResponse) {
	if login == "" {
		errorCode := diagnostics.BadLogin
		errorMessage := "Bad login."
		errorDescription := diagnostics.GetErrorDescription(errorCode)
		errorResponse := &responses.ErrorResponse{Code: errorCode, Message: errorMessage, Description: errorDescription}

		return false, errorResponse
	}

	if password == "" {
		errorCode := diagnostics.BadPassword
		errorMessage := "Bad password."
		errorDescription := diagnostics.GetErrorDescription(errorCode)
		errorResponse := &responses.ErrorResponse{Code: errorCode, Message: errorMessage, Description: errorDescription}

		return false, errorResponse
	}

	return true, nil
}
