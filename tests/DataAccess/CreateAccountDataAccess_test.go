package dataaccesss

import (
	"testing"

	dataAccess "../../src/DataAccess"
	models "../../src/Models"
	"../TestHelpers/Builders"
)

func TestCreateAccount_GivenCreateAccountRequest_ShouldAddAccountToMap(t *testing.T) {
	// Arrange
	builder := builders.CreateAccountRequestBuilder{}
	request := builder.WithUsername("Test").WithEmail("test@test.com").WithFirstName("Simon").WithSurname("Headley").WithPassword("Test1234").Build()

	expectedAccounts := make(map[int]models.Account)
	expectedAccounts[1] = models.Account{AccountID: 1, Login: "headleysj@gmail.com", Password: "password1234", FirstName: "Simon", Surname: "Headley"}
	expectedAccounts[2] = models.Account{AccountID: 2, Login: "test@test.com", Password: "Test1234", FirstName: "Simon", Surname: "Headley"}

	dataAccess := dataAccess.AccountDataAccess{}

	// Act
	actualAccounts, success := dataAccess.CreateAccount(request)

	// Assert
	if !success {
		t.Fail()
	}
	if len(actualAccounts) != len(expectedAccounts) {
		t.Fail()
	}
	for i := 1; i < len(actualAccounts); i++ {
		expected := expectedAccounts[i]
		actual := actualAccounts[i]

		if expected.AccountID != actual.AccountID {
			t.Fail()
		}
		if expected.Login != actual.Login {
			t.Fail()
		}
		if expected.Login != actual.Login {
			t.Fail()
		}
		if expected.Password != actual.Password {
			t.Fail()
		}
		if expected.FirstName != actual.FirstName {
			t.Fail()
		}
		if expected.Surname != actual.Surname {
			t.Fail()
		}
	}
}
