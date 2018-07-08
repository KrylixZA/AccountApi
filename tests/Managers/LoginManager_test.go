package managers

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"

	managers "../../src/Managers"
	models "../../src/Models"
	responses "../../src/Models/Responses"
	"../TestHelpers/Builders"
	"../TestHelpers/Mocks"
)

func TestLogin_GivenUnsuccessfulResult_ShouldReturn404AndErrorResponse(t *testing.T) {
	// Arrange
	login := "testLogin"
	password := "testPassword"

	responseBuilder := builders.ErrorResponseBuilder{}
	expectedStatusCode := http.StatusNotFound
	expectedResponse := responseBuilder.WithCode(1).WithMessage("Not found.").WithDescription("Login failed as no account was found with the given login and password.").Build()

	mockedAccountDataAccess := mocks.AccountDataAccessMock{
		LoginFunc: func(login string, password string) (*models.Account, bool) {
			return nil, false
		},
	}
	manager := managers.AccountManager{}

	// Act
	actualStatusCode, actualResponseByteArray := manager.Login(mockedAccountDataAccess, login, password)

	// Assert
	var actualResponse responses.ErrorResponse
	json.Unmarshal(actualResponseByteArray, &actualResponse)

	if expectedStatusCode != actualStatusCode {
		t.Errorf("Expected HTTP status code was %d. Actual HTTP status code was %d.", expectedStatusCode, actualStatusCode)
	}
	if expectedResponse.Code != actualResponse.Code {
		t.Errorf("Expected error code was %d. Actual error code was %d.", expectedResponse.Code, actualResponse.Code)
	}
	if expectedResponse.Message != actualResponse.Message {
		t.Errorf("Expected error message was %s. Actual error message was %s.", expectedResponse.Message, actualResponse.Message)
	}
	if expectedResponse.Description != actualResponse.Description {
		t.Errorf("Expected error description was %s. Actual error description was %s.", expectedResponse.Description, actualResponse.Description)
	}
	if !reflect.DeepEqual(expectedResponse, actualResponse) {
		t.Error("Deep equal shows struct properties not matching.")
	}
}

func TestLogin_GivenSuccessfulResult_ShouldReturn200AndAccountDetails(t *testing.T) {
	// Arrange
	login := "test@test.com"
	password := "Test1234"

	accountBuilder := builders.AccountBuilder{}

	expectedStatusCode := http.StatusOK
	expectedAccount := accountBuilder.WithAccountID(1).WithLogin("test@test.com").WithPassword("Test1234").WithFirstName("Simon").WithSurname("Headley").WithEmail("test@test.com").Build()

	mockedAccountDataAccess := mocks.AccountDataAccessMock{
		LoginFunc: func(login string, password string) (*models.Account, bool) {
			account := models.Account{AccountID: 1, Login: "test@test.com", Password: "Test1234", FirstName: "Simon", Surname: "Headley", Email: "test@test.com"}
			return &account, true
		},
	}
	manager := managers.AccountManager{}

	// Act
	actualStatusCode, actualResponseByteArray := manager.Login(mockedAccountDataAccess, login, password)

	// Assert
	var actualAccount models.Account
	json.Unmarshal(actualResponseByteArray, &actualAccount)

	if expectedStatusCode != actualStatusCode {
		t.Errorf("Expected HTTP status was %d. Actual HTTP status code was %d.", expectedStatusCode, actualStatusCode)
	}
	if expectedAccount.AccountID != actualAccount.AccountID {
		t.Errorf("Expected AccountID %d. Actual AccountID %d.", expectedAccount.AccountID, actualAccount.AccountID)
	}
	if expectedAccount.Login != actualAccount.Login {
		t.Errorf("Expected Login %s. Actual Login %s.", expectedAccount.Login, actualAccount.Login)
	}
	if expectedAccount.Password != actualAccount.Password {
		t.Errorf("Expected Password %s. Actual Password %s.", expectedAccount.Password, actualAccount.Password)
	}
	if expectedAccount.FirstName != actualAccount.FirstName {
		t.Errorf("Expected FirstName %s. Actual FirstName %s.", expectedAccount.FirstName, actualAccount.FirstName)
	}
	if expectedAccount.Surname != actualAccount.Surname {
		t.Errorf("Expected Surname %s. Actual Surname %s.", expectedAccount.Surname, actualAccount.Surname)
	}
	if expectedAccount.Email != actualAccount.Email {
		t.Errorf("Expected Email %s. Actual Email %s.", expectedAccount.Email, actualAccount.Email)
	}
	if !reflect.DeepEqual(expectedAccount, actualAccount) {
		t.Error("Deep equal shows struct properties not matching.")
	}
}
