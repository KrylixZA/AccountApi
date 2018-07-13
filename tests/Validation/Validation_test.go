package validation

import (
	"reflect"
	"testing"

	"../../src/Diagnostics"
	"../TestHelpers/Builders"
)

func TestValidateCreateAccountRequest_GivenNilRequest_ShouldReturnFalseAndExpectedErrorReponse(t *testing.T) {
	// Arrange
	responseBuilder := builders.ErrorResponseBuilder{}

	expectedErrorCode := diagnostics.BadRequest
	expectedErrorMessage := "Bad request."
	expectedErrorDescription := diagnostics.GetErrorDescription(expectedErrorCode)
	expectedErrorResponse := responseBuilder.WithCode(expectedErrorCode).WithMessage(expectedErrorMessage).WithDescription(expectedErrorDescription).Build()

	validation := getSystemUnderTestValidation()

	// Act
	validRequest, actualErrorResponse := validation.ValidateCreateAccountRequest(nil)

	// Assert
	if validRequest != false {
		t.Error("Expected valid request to be false. Actual valid request was true.")
	}
	if expectedErrorCode != actualErrorResponse.Code {
		t.Errorf("Expected error code was %d. Actual error code was %d.", expectedErrorCode, actualErrorResponse.Code)
	}
	if expectedErrorMessage != actualErrorResponse.Message {
		t.Errorf("Expected error message was %s. Actual error message was %s.", expectedErrorMessage, actualErrorResponse.Message)
	}
	if expectedErrorDescription != actualErrorResponse.Description {
		t.Errorf("Expected error description was %s. Actual error description was %s.", expectedErrorDescription, actualErrorResponse.Description)
	}
	if !reflect.DeepEqual(expectedErrorResponse, actualErrorResponse) {
		t.Error("Deep equal shows struct properties not matching.")
	}
}

func TestValidateCreateAccountRequest_GivenRequestWithBadUsername_ShouldReturnFalseAndExpectedErrorReponse(t *testing.T) {
	// Arrange
	requestBuilder := builders.CreateAccountRequestBuilder{}
	responseBuilder := builders.ErrorResponseBuilder{}

	username := ""
	email := "test@test.com"
	firstName := "Simon"
	surname := "Headley"
	password := "test1234"
	request := requestBuilder.WithUsername(username).WithEmail(email).WithFirstName(firstName).WithSurname(surname).WithPassword(password).Build()

	expectedErrorCode := diagnostics.BadUsername
	expectedErrorMessage := "Bad username."
	expectedErrorDescription := diagnostics.GetErrorDescription(expectedErrorCode)
	expectedErrorResponse := responseBuilder.WithCode(expectedErrorCode).WithMessage(expectedErrorMessage).WithDescription(expectedErrorDescription).Build()

	validation := getSystemUnderTestValidation()

	// Act
	validRequest, actualErrorResponse := validation.ValidateCreateAccountRequest(request)

	// Assert
	if validRequest != false {
		t.Error("Expected valid request to be false. Actual valid request was true.")
	}
	if expectedErrorCode != actualErrorResponse.Code {
		t.Errorf("Expected error code was %d. Actual error code was %d.", expectedErrorCode, actualErrorResponse.Code)
	}
	if expectedErrorMessage != actualErrorResponse.Message {
		t.Errorf("Expected error message was %s. Actual error message was %s.", expectedErrorMessage, actualErrorResponse.Message)
	}
	if expectedErrorDescription != actualErrorResponse.Description {
		t.Errorf("Expected error description was %s. Actual error description was %s.", expectedErrorDescription, actualErrorResponse.Description)
	}
	if !reflect.DeepEqual(expectedErrorResponse, actualErrorResponse) {
		t.Error("Deep equal shows struct properties not matching.")
	}
}

func TestValidateCreateAccountRequest_GivenRequestWithBadEmail_ShouldReturnFalseAndExpectedErrorReponse(t *testing.T) {
	// Arrange
	requestBuilder := builders.CreateAccountRequestBuilder{}
	responseBuilder := builders.ErrorResponseBuilder{}

	username := "test"
	email := ""
	firstName := "Simon"
	surname := "Headley"
	password := "test1234"
	request := requestBuilder.WithUsername(username).WithEmail(email).WithFirstName(firstName).WithSurname(surname).WithPassword(password).Build()

	expectedErrorCode := diagnostics.BadEmail
	expectedErrorMessage := "Bad email."
	expectedErrorDescription := diagnostics.GetErrorDescription(expectedErrorCode)
	expectedErrorResponse := responseBuilder.WithCode(expectedErrorCode).WithMessage(expectedErrorMessage).WithDescription(expectedErrorDescription).Build()

	validation := getSystemUnderTestValidation()

	// Act
	validRequest, actualErrorResponse := validation.ValidateCreateAccountRequest(request)

	// Assert
	if validRequest != false {
		t.Error("Expected valid request to be false. Actual valid request was true.")
	}
	if expectedErrorCode != actualErrorResponse.Code {
		t.Errorf("Expected error code was %d. Actual error code was %d.", expectedErrorCode, actualErrorResponse.Code)
	}
	if expectedErrorMessage != actualErrorResponse.Message {
		t.Errorf("Expected error message was %s. Actual error message was %s.", expectedErrorMessage, actualErrorResponse.Message)
	}
	if expectedErrorDescription != actualErrorResponse.Description {
		t.Errorf("Expected error description was %s. Actual error description was %s.", expectedErrorDescription, actualErrorResponse.Description)
	}
	if !reflect.DeepEqual(expectedErrorResponse, actualErrorResponse) {
		t.Error("Deep equal shows struct properties not matching.")
	}
}

func TestValidateCreateAccountRequest_GivenRequestWithBadFirstName_ShouldReturnFalseAndExpectedErrorReponse(t *testing.T) {
	// Arrange
	requestBuilder := builders.CreateAccountRequestBuilder{}
	responseBuilder := builders.ErrorResponseBuilder{}

	username := "test"
	email := "test@test.com"
	firstName := ""
	surname := "Headley"
	password := "test1234"
	request := requestBuilder.WithUsername(username).WithEmail(email).WithFirstName(firstName).WithSurname(surname).WithPassword(password).Build()

	expectedErrorCode := diagnostics.BadFirstName
	expectedErrorMessage := "Bad first name."
	expectedErrorDescription := diagnostics.GetErrorDescription(expectedErrorCode)
	expectedErrorResponse := responseBuilder.WithCode(expectedErrorCode).WithMessage(expectedErrorMessage).WithDescription(expectedErrorDescription).Build()

	validation := getSystemUnderTestValidation()

	// Act
	validRequest, actualErrorResponse := validation.ValidateCreateAccountRequest(request)

	// Assert
	if validRequest != false {
		t.Error("Expected valid request to be false. Actual valid request was true.")
	}
	if expectedErrorCode != actualErrorResponse.Code {
		t.Errorf("Expected error code was %d. Actual error code was %d.", expectedErrorCode, actualErrorResponse.Code)
	}
	if expectedErrorMessage != actualErrorResponse.Message {
		t.Errorf("Expected error message was %s. Actual error message was %s.", expectedErrorMessage, actualErrorResponse.Message)
	}
	if expectedErrorDescription != actualErrorResponse.Description {
		t.Errorf("Expected error description was %s. Actual error description was %s.", expectedErrorDescription, actualErrorResponse.Description)
	}
	if !reflect.DeepEqual(expectedErrorResponse, actualErrorResponse) {
		t.Error("Deep equal shows struct properties not matching.")
	}
}

func TestValidateCreateAccountRequest_GivenRequestWithBadSurname_ShouldReturnFalseAndExpectedErrorReponse(t *testing.T) {
	// Arrange
	requestBuilder := builders.CreateAccountRequestBuilder{}
	responseBuilder := builders.ErrorResponseBuilder{}

	username := "test"
	email := "test@test.com"
	firstName := "Simon"
	surname := ""
	password := "test1234"
	request := requestBuilder.WithUsername(username).WithEmail(email).WithFirstName(firstName).WithSurname(surname).WithPassword(password).Build()

	expectedErrorCode := diagnostics.BadSurname
	expectedErrorMessage := "Bad surname."
	expectedErrorDescription := diagnostics.GetErrorDescription(expectedErrorCode)
	expectedErrorResponse := responseBuilder.WithCode(expectedErrorCode).WithMessage(expectedErrorMessage).WithDescription(expectedErrorDescription).Build()

	validation := getSystemUnderTestValidation()

	// Act
	validRequest, actualErrorResponse := validation.ValidateCreateAccountRequest(request)

	// Assert
	if validRequest != false {
		t.Error("Expected valid request to be false. Actual valid request was true.")
	}
	if expectedErrorCode != actualErrorResponse.Code {
		t.Errorf("Expected error code was %d. Actual error code was %d.", expectedErrorCode, actualErrorResponse.Code)
	}
	if expectedErrorMessage != actualErrorResponse.Message {
		t.Errorf("Expected error message was %s. Actual error message was %s.", expectedErrorMessage, actualErrorResponse.Message)
	}
	if expectedErrorDescription != actualErrorResponse.Description {
		t.Errorf("Expected error description was %s. Actual error description was %s.", expectedErrorDescription, actualErrorResponse.Description)
	}
	if !reflect.DeepEqual(expectedErrorResponse, actualErrorResponse) {
		t.Error("Deep equal shows struct properties not matching.")
	}
}

func TestValidateCreateAccountRequest_GivenRequestWithBadPassword_ShouldReturnFalseAndExpectedErrorReponse(t *testing.T) {
	// Arrange
	requestBuilder := builders.CreateAccountRequestBuilder{}
	responseBuilder := builders.ErrorResponseBuilder{}

	username := "test"
	email := "test@test.com"
	firstName := "Simon"
	surname := "Headley"
	password := ""
	request := requestBuilder.WithUsername(username).WithEmail(email).WithFirstName(firstName).WithSurname(surname).WithPassword(password).Build()

	expectedErrorCode := diagnostics.BadPassword
	expectedErrorMessage := "Bad password."
	expectedErrorDescription := diagnostics.GetErrorDescription(expectedErrorCode)
	expectedErrorResponse := responseBuilder.WithCode(expectedErrorCode).WithMessage(expectedErrorMessage).WithDescription(expectedErrorDescription).Build()

	validation := getSystemUnderTestValidation()

	// Act
	validRequest, actualErrorResponse := validation.ValidateCreateAccountRequest(request)

	// Assert
	if validRequest != false {
		t.Error("Expected valid request to be false. Actual valid request was true.")
	}
	if expectedErrorCode != actualErrorResponse.Code {
		t.Errorf("Expected error code was %d. Actual error code was %d.", expectedErrorCode, actualErrorResponse.Code)
	}
	if expectedErrorMessage != actualErrorResponse.Message {
		t.Errorf("Expected error message was %s. Actual error message was %s.", expectedErrorMessage, actualErrorResponse.Message)
	}
	if expectedErrorDescription != actualErrorResponse.Description {
		t.Errorf("Expected error description was %s. Actual error description was %s.", expectedErrorDescription, actualErrorResponse.Description)
	}
	if !reflect.DeepEqual(expectedErrorResponse, actualErrorResponse) {
		t.Error("Deep equal shows struct properties not matching.")
	}
}

func TestValidateCreateAccountRequest_GivenWellFormedRequest_ShouldReturnTrueAndNil(t *testing.T) {
	// Arrange
	requestBuilder := builders.CreateAccountRequestBuilder{}

	username := "test"
	email := "test@test.com"
	firstName := "Simon"
	surname := "Headley"
	password := "test1234"
	request := requestBuilder.WithUsername(username).WithEmail(email).WithFirstName(firstName).WithSurname(surname).WithPassword(password).Build()

	validation := getSystemUnderTestValidation()

	// Act
	validRequest, actualErrorResponse := validation.ValidateCreateAccountRequest(request)

	// Assert
	if validRequest != true {
		t.Error("Expected valid request to be true. Actual valid request was false.")
	}
	if actualErrorResponse != nil {
		t.Error("Expected actual error response to be nil. Actual was not nil.")
	}
}

func TestResetPasswordRequest_GivenNilRequest_ShouldReturnFalseAndExpectedErrorResponse(t *testing.T) {
	// Arrange
	responseBuilder := builders.ErrorResponseBuilder{}

	expectedErrorCode := diagnostics.BadRequest
	expectedErrorMessage := "Bad request."
	expectedErrorDescription := diagnostics.GetErrorDescription(expectedErrorCode)
	expectedErrorResponse := responseBuilder.WithCode(expectedErrorCode).WithMessage(expectedErrorMessage).WithDescription(expectedErrorDescription).Build()

	validation := getSystemUnderTestValidation()

	// Act
	validRequest, actualErrorResponse := validation.ValidateResetPasswordRequest(nil)

	// Assert
	if validRequest != false {
		t.Error("Expected valid request to be false. Actual valid request was true.")
	}
	if expectedErrorCode != actualErrorResponse.Code {
		t.Errorf("Expected error code was %d. Actual error code was %d.", expectedErrorCode, actualErrorResponse.Code)
	}
	if expectedErrorMessage != actualErrorResponse.Message {
		t.Errorf("Expected error message was %s. Actual error message was %s.", expectedErrorMessage, actualErrorResponse.Message)
	}
	if expectedErrorDescription != actualErrorResponse.Description {
		t.Errorf("Expected error description was %s. Actual error description was %s.", expectedErrorDescription, actualErrorResponse.Description)
	}
	if !reflect.DeepEqual(expectedErrorResponse, actualErrorResponse) {
		t.Error("Deep equal shows struct properties not matching.")
	}
}

func TestResetPasswordRequest_GivenRequestWithBadCurrentPassword_ShouldReturnFalseAndExpectedErrorResponse(t *testing.T) {
	// Arrange
	responseBuilder := builders.ErrorResponseBuilder{}
	requestBuilder := builders.ResetPasswordRequestBuilder{}

	currentPassword := ""
	newPassword := "test5678"
	request := requestBuilder.WithCurrentPassword(currentPassword).WithNewPassword(newPassword).Build()

	expectedErrorCode := diagnostics.BadCurrentPassword
	expectedErrorMessage := "Bad current password."
	expectedErrorDescription := diagnostics.GetErrorDescription(expectedErrorCode)
	expectedErrorResponse := responseBuilder.WithCode(expectedErrorCode).WithMessage(expectedErrorMessage).WithDescription(expectedErrorDescription).Build()

	validation := getSystemUnderTestValidation()

	// Act
	validRequest, actualErrorResponse := validation.ValidateResetPasswordRequest(request)

	// Assert
	if validRequest != false {
		t.Error("Expected valid request to be false. Actual valid request was true.")
	}
	if expectedErrorCode != actualErrorResponse.Code {
		t.Errorf("Expected error code was %d. Actual error code was %d.", expectedErrorCode, actualErrorResponse.Code)
	}
	if expectedErrorMessage != actualErrorResponse.Message {
		t.Errorf("Expected error message was %s. Actual error message was %s.", expectedErrorMessage, actualErrorResponse.Message)
	}
	if expectedErrorDescription != actualErrorResponse.Description {
		t.Errorf("Expected error description was %s. Actual error description was %s.", expectedErrorDescription, actualErrorResponse.Description)
	}
	if !reflect.DeepEqual(expectedErrorResponse, actualErrorResponse) {
		t.Error("Deep equal shows struct properties not matching.")
	}
}

func TestResetPasswordRequest_GivenRequestWithBadNewPassword_ShouldReturnFalseAndExpectedErrorResponse(t *testing.T) {
	// Arrange
	responseBuilder := builders.ErrorResponseBuilder{}
	requestBuilder := builders.ResetPasswordRequestBuilder{}

	currentPassword := "test1234"
	newPassword := ""
	request := requestBuilder.WithCurrentPassword(currentPassword).WithNewPassword(newPassword).Build()

	expectedErrorCode := diagnostics.BadNewPassword
	expectedErrorMessage := "Bad new password."
	expectedErrorDescription := diagnostics.GetErrorDescription(expectedErrorCode)
	expectedErrorResponse := responseBuilder.WithCode(expectedErrorCode).WithMessage(expectedErrorMessage).WithDescription(expectedErrorDescription).Build()

	validation := getSystemUnderTestValidation()

	// Act
	validRequest, actualErrorResponse := validation.ValidateResetPasswordRequest(request)

	// Assert
	if validRequest != false {
		t.Error("Expected valid request to be false. Actual valid request was true.")
	}
	if expectedErrorCode != actualErrorResponse.Code {
		t.Errorf("Expected error code was %d. Actual error code was %d.", expectedErrorCode, actualErrorResponse.Code)
	}
	if expectedErrorMessage != actualErrorResponse.Message {
		t.Errorf("Expected error message was %s. Actual error message was %s.", expectedErrorMessage, actualErrorResponse.Message)
	}
	if expectedErrorDescription != actualErrorResponse.Description {
		t.Errorf("Expected error description was %s. Actual error description was %s.", expectedErrorDescription, actualErrorResponse.Description)
	}
	if !reflect.DeepEqual(expectedErrorResponse, actualErrorResponse) {
		t.Error("Deep equal shows struct properties not matching.")
	}
}

func TestResetPasswordRequest_GivenWellFormedRequest_ShouldReturnFalseAndExpectedErrorResponse(t *testing.T) {
	// Arrange
	requestBuilder := builders.ResetPasswordRequestBuilder{}

	currentPassword := "test1234"
	newPassword := "test5678"
	request := requestBuilder.WithCurrentPassword(currentPassword).WithNewPassword(newPassword).Build()

	validation := getSystemUnderTestValidation()

	// Act
	validRequest, actualErrorResponse := validation.ValidateResetPasswordRequest(request)

	// Assert
	if validRequest != true {
		t.Error("Expected valid request to be true. Actual valid request was false.")
	}
	if actualErrorResponse != nil {
		t.Error("Expected actual error response to be nil. Actual was not nil.")
	}
}

func TestValidateLoginAndPassword_GivenBadLogin_ShouldReturnFalseAndExpectedErrorResponse(t *testing.T) {
	// Arrange
	responseBuilder := builders.ErrorResponseBuilder{}

	login := ""
	password := "test1234"

	expectedErrorCode := diagnostics.BadLogin
	expectedErrorMessage := "Bad login."
	expectedErrorDescription := diagnostics.GetErrorDescription(expectedErrorCode)
	expectedErrorResponse := responseBuilder.WithCode(expectedErrorCode).WithMessage(expectedErrorMessage).WithDescription(expectedErrorDescription).Build()

	validation := getSystemUnderTestValidation()

	// Act
	validRequest, actualErrorResponse := validation.ValidateLoginAndPassword(login, password)

	// Assert
	if validRequest != false {
		t.Error("Expected valid request to be false. Actual valid request was true.")
	}
	if expectedErrorCode != actualErrorResponse.Code {
		t.Errorf("Expected error code was %d. Actual error code was %d.", expectedErrorCode, actualErrorResponse.Code)
	}
	if expectedErrorMessage != actualErrorResponse.Message {
		t.Errorf("Expected error message was %s. Actual error message was %s.", expectedErrorMessage, actualErrorResponse.Message)
	}
	if expectedErrorDescription != actualErrorResponse.Description {
		t.Errorf("Expected error description was %s. Actual error description was %s.", expectedErrorDescription, actualErrorResponse.Description)
	}
	if !reflect.DeepEqual(expectedErrorResponse, actualErrorResponse) {
		t.Error("Deep equal shows struct properties not matching.")
	}
}

func TestValidateLoginAndPassword_GivenBadPassword_ShouldReturnFalseAndExpectedErrorResponse(t *testing.T) {
	// Arrange
	responseBuilder := builders.ErrorResponseBuilder{}

	login := "test"
	password := ""

	expectedErrorCode := diagnostics.BadPassword
	expectedErrorMessage := "Bad password."
	expectedErrorDescription := diagnostics.GetErrorDescription(expectedErrorCode)
	expectedErrorResponse := responseBuilder.WithCode(expectedErrorCode).WithMessage(expectedErrorMessage).WithDescription(expectedErrorDescription).Build()

	validation := getSystemUnderTestValidation()

	// Act
	validRequest, actualErrorResponse := validation.ValidateLoginAndPassword(login, password)

	// Assert
	if validRequest != false {
		t.Error("Expected valid request to be false. Actual valid request was true.")
	}
	if expectedErrorCode != actualErrorResponse.Code {
		t.Errorf("Expected error code was %d. Actual error code was %d.", expectedErrorCode, actualErrorResponse.Code)
	}
	if expectedErrorMessage != actualErrorResponse.Message {
		t.Errorf("Expected error message was %s. Actual error message was %s.", expectedErrorMessage, actualErrorResponse.Message)
	}
	if expectedErrorDescription != actualErrorResponse.Description {
		t.Errorf("Expected error description was %s. Actual error description was %s.", expectedErrorDescription, actualErrorResponse.Description)
	}
	if !reflect.DeepEqual(expectedErrorResponse, actualErrorResponse) {
		t.Error("Deep equal shows struct properties not matching.")
	}
}

func TestValidateLoginAndPassword_GivenWellFormedLoginAndPassword_ShouldReturnTrueAndNil(t *testing.T) {
	// Arrange
	login := "test"
	password := "test1234"

	validation := getSystemUnderTestValidation()

	// Act
	validRequest, actualErrorResponse := validation.ValidateLoginAndPassword(login, password)

	// Assert
	if validRequest != true {
		t.Error("Expected valid request to be true. Actual valid request was false.")
	}
	if actualErrorResponse != nil {
		t.Error("Expected actual error response to be nil. Actual was not nil.")
	}
}
