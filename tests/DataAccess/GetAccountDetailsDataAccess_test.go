package dataaccess

import (
	"reflect"
	"testing"

	dataAccess "../../src/DataAccess"
	models "../../src/Models"
)

func TestGetAccountDetails_GivenAccountIdInMap_ShouldReturnAccountAndSuccessfulResult(t *testing.T) {
	// Arrange
	accountID := 1

	expectedAccount := &models.Account{AccountID: 1, Login: "headleysj@gmail.com", Password: "password1234", FirstName: "Simon", Surname: "Headley"}

	dataAccess := dataAccess.AccountDataAccess{}

	// Act
	actualAccount, success := dataAccess.GetAccountDetails(accountID)

	// Assert
	if !success {
		t.Error("Expected successful result. Actual result was unsuccessful.")
	}
	if expectedAccount.AccountID != actualAccount.AccountID {
		t.Errorf("Expected AccountID was %d. Actual AccountID was %d", expectedAccount.AccountID, actualAccount.AccountID)
	}
	if expectedAccount.Login != actualAccount.Login {
		t.Errorf("Expected Login was %s. Actual Login was %s", expectedAccount.Login, actualAccount.Login)
	}
	if expectedAccount.Password != actualAccount.Password {
		t.Errorf("Expected Password was %s. Actual Password was %s", expectedAccount.Password, actualAccount.Password)
	}
	if expectedAccount.FirstName != actualAccount.FirstName {
		t.Errorf("Expected FirstName was %s. Actual FirstName was %s", expectedAccount.FirstName, actualAccount.FirstName)
	}
	if expectedAccount.Surname != actualAccount.Surname {
		t.Errorf("Expected AccountID was %s. Actual AccountID was %s", expectedAccount.Surname, actualAccount.Surname)
	}
	if !reflect.DeepEqual(expectedAccount, actualAccount) {
		t.Error("Deep equal shows structs are not matching.")
	}
}

func TestGetAccountGetails_GivenAccountIdNotInMap_ShouldReturnNilAndUnsuccessfulResult(t *testing.T) {
	// Arrange
	accountID := 2

	dataAccess := dataAccess.AccountDataAccess{}

	// Act
	account, success := dataAccess.GetAccountDetails(accountID)

	// Assert
	if success {
		t.Error("Expected unsuccessful result. Actual result was successful.")
	}
	if account != nil {
		t.Error("Expected nil account. Actual account was not nil.")
	}
}
