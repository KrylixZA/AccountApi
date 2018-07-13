package dataaccess

import (
	"reflect"
	"testing"

	"../TestHelpers/Builders"
)

func TestLogin_GivenIncorrectLoginAndIncorrectPassword_ShouldReturnNilAndUnsuccessfulResult(t *testing.T) {
	// Arrange
	login := "someBadLogin"
	password := "someBadPassword"

	dataAccess := getSystemUnderTestAccountDataAccess()

	// Act
	success, actualAccount := dataAccess.Login(login, password)

	// Assert
	if success {
		t.Error("Expected unsuccessful result. Actual result was successful.")
	}
	if actualAccount != nil {
		t.Error("Expected nil account. Actual account was not nil.")
	}
}

func TestLogin_GivenIncorrectLoginAndCorrectPassword_ShouldReturnNilAndUnsuccessfulResult(t *testing.T) {
	// Arrange
	login := "someBadLogin"
	password := "password1234"

	dataAccess := getSystemUnderTestAccountDataAccess()

	// Act
	success, actualAccount := dataAccess.Login(login, password)

	// Assert
	if success {
		t.Error("Expected unsuccessful result. Actual result was successful.")
	}
	if actualAccount != nil {
		t.Error("Expected nil account. Actual account was not nil.")
	}
}

func TestLogin_GivenCorrectLoginAndIncorrectPassword_ShouldReturnNilAndUnsuccessfulResult(t *testing.T) {
	// Arrange
	login := "headleysj@gmail.com"
	password := "someBadPassword"

	dataAccess := getSystemUnderTestAccountDataAccess()

	// Act
	success, actualAccount := dataAccess.Login(login, password)

	// Assert
	if success {
		t.Error("Expected unsuccessful result. Actual result was successful.")
	}
	if actualAccount != nil {
		t.Error("Expected nil account. Actual account was not nil.")
	}
}

func TestLogin_GivenCorrectLoginAndCorrectPassword_ShouldReturnAccountAndSuccessfulResult(t *testing.T) {
	// Arrange
	login := "headleysj@gmail.com"
	password := "password1234"

	accountBuilder := builders.AccountBuilder{}
	expectedAccount := accountBuilder.WithAccountID(1).WithLogin("headleysj@gmail.com").WithPassword("password1234").WithFirstName("Simon").WithSurname("Headley").WithEmail("headleysj@gmail.com").Build()

	dataAccess := getSystemUnderTestAccountDataAccess()

	// Act
	success, actualAccount := dataAccess.Login(login, password)

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
