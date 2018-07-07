package dataaccess

import (
	"reflect"
	"testing"

	dataAccess "../../src/DataAccess"
	models "../../src/Models"
	"../TestHelpers/Builders"
)

func TestCreateAccount_GivenCreateAccountRequest_ShouldReturnAccountsAndSuccessfulResult(t *testing.T) {
	// Arrange
	requestBuilder := builders.CreateAccountRequestBuilder{}
	request := requestBuilder.WithUsername("Test").WithEmail("test@test.com").WithFirstName("Simon").WithSurname("Headley").WithPassword("Test1234").Build()

	accountBuilder := builders.AccountBuilder{}
	expectedAccounts := make(map[int]models.Account)
	expectedAccounts[1] = accountBuilder.WithAccountID(1).WithLogin("headleysj@gmail.com").WithPassword("password1234").WithFirstName("Simon").WithSurname("Headley").WithEmail("headleysj@gmail.com").Build()
	expectedAccounts[2] = accountBuilder.WithAccountID(2).WithLogin("test@test.com").WithPassword("Test1234").WithFirstName("Simon").WithSurname("Headley").WithEmail("test@test.com").Build()

	dataAccess := dataAccess.AccountDataAccess{}

	// Act
	actualAccounts, success := dataAccess.CreateAccount(request)

	// Assert
	if !success {
		t.Error("Expected successful result. Actual result was unsuccessful.")
	}
	if len(actualAccounts) != len(expectedAccounts) {
		t.Errorf("Expected map count was %d. Actual map count was %d", len(expectedAccounts), len(actualAccounts))
	}
	for i := 1; i < len(actualAccounts); i++ {
		expected := expectedAccounts[i]
		actual := actualAccounts[i]

		if expected.AccountID != actual.AccountID {
			t.Errorf("Expected AccountID was %d. Actual AccountID was %d.", expected.AccountID, actual.AccountID)
		}
		if expected.Login != actual.Login {
			t.Errorf("Expected Login was %s. Actual Login was %s.", expected.Login, actual.Login)
		}
		if expected.Password != actual.Password {
			t.Errorf("Expected Password was %s. Actual Password was %s.", expected.Password, actual.Password)
		}
		if expected.FirstName != actual.FirstName {
			t.Errorf("Expected FirstName was %s. Actual FirstName was %s.", expected.FirstName, actual.FirstName)
		}
		if expected.Surname != actual.Surname {
			t.Errorf("Expected Surname was %s. Actual Surname was %s.", expected.Surname, actual.Surname)
		}
		if expected.Email != actual.Email {
			t.Errorf("Expected Email was %s. Actual Email was %s.", expected.Email, actual.Email)
		}
		if !reflect.DeepEqual(expected, actual) {
			t.Error("Deep equal shows struct properties not matching.")
		}
	}
}
