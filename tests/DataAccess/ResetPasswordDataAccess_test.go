package dataaccess

import (
	"reflect"
	"testing"

	"../TestHelpers/Builders"
)

func TestResetPassword_GivenRequestWithIncorrectAccountIdAndIncorrectCurrentPassword_ShouldReturnNilAndUnsuccessfulResult(t *testing.T) {
	// Arrange
	builder := builders.ResetPasswordRequestBuilder{}
	request := builder.WithCurrentPassword("test1234").WithNewPassword("test1234").Build()

	dataAccess := getSystemUnderTestAccountDataAccess()

	// Act
	actualAccount, success := dataAccess.ResetPassword(-1, *request)

	// Assert
	if success {
		t.Error("Expected unsuccessful result. Actual result was successful.")
	}
	if actualAccount != nil {
		t.Error("Expected nil account. Actual account was not nil.")
	}
}

func TestResetPassword_GivenRequestWithIncorrectAccountIdAndCorrectCurrentPassword_ShouldReturnNilAndUnsuccessfulResult(t *testing.T) {
	// Arrange
	builder := builders.ResetPasswordRequestBuilder{}
	request := builder.WithCurrentPassword("password1234").WithNewPassword("test1234").Build()

	dataAccess := getSystemUnderTestAccountDataAccess()

	// Act
	actualAccount, success := dataAccess.ResetPassword(-1, *request)

	// Assert
	if success {
		t.Error("Expected unsuccessful result. Actual result was successful.")
	}
	if actualAccount != nil {
		t.Error("Expected nil account. Actual account was not nil.")
	}
}

func TestResetPassword_GivenRequestWithCorrectAccountIdAndIncorrectCurrentPassword_ShouldReturnNilAndUnsuccessfulResult(t *testing.T) {
	// Arrange
	builder := builders.ResetPasswordRequestBuilder{}
	request := builder.WithCurrentPassword("test1234").WithNewPassword("test1234").Build()

	dataAccess := getSystemUnderTestAccountDataAccess()

	// Act
	actualAccount, success := dataAccess.ResetPassword(1, *request)

	// Assert
	if success {
		t.Error("Expected unsuccessful result. Actual result was successful.")
	}
	if actualAccount != nil {
		t.Error("Expected nil account. Actual account was not nil.")
	}
}

func TestResetPassword_GivenRequestWithCorrectAccountIdAndCorrectCurrentPassword_ShouldReturnAccountAndSuccessfulResult(t *testing.T) {
	// Arrange
	requestBuilder := builders.ResetPasswordRequestBuilder{}
	request := requestBuilder.WithCurrentPassword("password1234").WithNewPassword("test1234").Build()

	accountBuilder := builders.AccountBuilder{}
	expectedAccount := accountBuilder.WithAccountID(1).WithLogin("headleysj@gmail.com").WithPassword("test1234").WithFirstName("Simon").WithSurname("Headley").WithEmail("headleysj@gmail.com").Build()

	dataAccess := getSystemUnderTestAccountDataAccess()

	// Act
	actualAccount, success := dataAccess.ResetPassword(1, *request)

	// Assert
	if !success {
		t.Error("Expected successful result. Actual result was unsuccessful.")
	}
	if expectedAccount.AccountID != actualAccount.AccountID {
		t.Errorf("Expected AccountID was %d. Actual AccountID was %d.", expectedAccount.AccountID, actualAccount.AccountID)
	}
	if expectedAccount.Login != actualAccount.Login {
		t.Errorf("Expected Login was %s. Actual Login was %s.", expectedAccount.Login, actualAccount.Login)
	}
	if expectedAccount.Password != actualAccount.Password {
		t.Errorf("Expected Password was %s. Actual Password was %s.", expectedAccount.Password, actualAccount.Password)
	}
	if expectedAccount.FirstName != actualAccount.FirstName {
		t.Errorf("Expected FirstName was %s. Actual FirstName was %s.", expectedAccount.FirstName, actualAccount.FirstName)
	}
	if expectedAccount.Surname != actualAccount.Surname {
		t.Errorf("Expected Surname was %s. Actual Surname was %s.", expectedAccount.Surname, actualAccount.Surname)
	}
	if expectedAccount.Email != actualAccount.Email {
		t.Errorf("Expected Email was %s. Actual Email was %s.", expectedAccount.Email, actualAccount.Email)
	}
	if !reflect.DeepEqual(expectedAccount, actualAccount) {
		t.Error("Deep equal shows structs are not matching.")
	}
}
