package dataaccess

import (
	"reflect"
	"testing"

	"../TestHelpers/Builders"
)

func TestCreateAccount_GivenCreateAccountRequest_ShouldReturnAccountsAndSuccessfulResult(t *testing.T) {
	// Arrange
	requestBuilder := builders.CreateAccountRequestBuilder{}
	request := requestBuilder.WithUsername("Test").WithEmail("test@test.com").WithFirstName("Simon").WithSurname("Headley").WithPassword("Test1234").Build()

	accountBuilder := builders.AccountBuilder{}
	expectedAccount := accountBuilder.WithAccountID(2).WithLogin("test@test.com").WithPassword("Test1234").WithFirstName("Simon").WithSurname("Headley").WithEmail("test@test.com").Build()

	dataAccess := getSystemUnderTestAccountDataAccess()

	// Act
	actualAccount, success := dataAccess.CreateAccount(*request)

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
		t.Error("Deep equal shows struct properties not matching.")
	}
}
