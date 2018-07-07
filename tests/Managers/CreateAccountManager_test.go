package managers

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"

	managers "../../src/Managers"
	models "../../src/Models"
	requests "../../src/Models/Requests"
	responses "../../src/Models/Responses"
	"../TestHelpers/Builders"
	"../TestHelpers/Mocks"
)

func TestCreateAccount_GivenValidRequestAndUnsuccessfulResult_ShouldReturn500AndErrorResponse(t *testing.T) {
	// Arrange
	requestBuilder := builders.CreateAccountRequestBuilder{}
	responseBuilder := builders.ErrorResponseBuilder{}

	request := requestBuilder.WithUsername("test@test.com").WithEmail("test@test.com").WithFirstName("Simon").WithSurname("Headley").WithPassword("Test1234").Build()

	expectedStatusCode := http.StatusInternalServerError
	expectedResponse := responseBuilder.WithCode(3).WithMessage("Internal Server Error.").WithDescription("Failed to create account.").Build()

	mockedAccountDataAccess := mocks.AccountDataAccessMock{
		CreateAccountFunc: func(request requests.CreateAccountRequest) (map[int]models.Account, bool) {
			return nil, false
		},
	}
	manager := managers.AccountManager{}

	// Act
	actualStatusCode, actualResponseByteArray := manager.CreateAccount(mockedAccountDataAccess, request)

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

func TestCreateAccount_GivenValidRequestAndSuccessfulResult_ShouldReturn201AndMapOfWithCreatedAccount(t *testing.T) {
	// Arrange
	requestBuilder := builders.CreateAccountRequestBuilder{}
	accountBuilder := builders.AccountBuilder{}

	request := requestBuilder.WithUsername("test@test.com").WithEmail("test@test.com").WithFirstName("Simon").WithSurname("Headley").WithPassword("Test1234").Build()

	expectedStatusCode := http.StatusCreated
	expectedAccount := accountBuilder.WithAccountID(1).WithLogin("test@test.com").WithPassword("Test1234").WithFirstName("Simon").WithSurname("Headley").WithEmail("test@test.com").Build()

	mockedAccountDataAccess := mocks.AccountDataAccessMock{
		CreateAccountFunc: func(request requests.CreateAccountRequest) (map[int]models.Account, bool) {
			accounts := make(map[int]models.Account)
			accounts[1] = models.Account{AccountID: 1, Login: request.Email, Password: request.Password, FirstName: request.FirstName, Surname: request.Surname, Email: request.Email}
			return accounts, true
		},
	}
	manager := managers.AccountManager{}

	// Act
	actualStatusCode, actualResponseByteArray := manager.CreateAccount(mockedAccountDataAccess, request)

	// Assert
	var actualResponse map[int]models.Account
	json.Unmarshal(actualResponseByteArray, &actualResponse)

	if expectedStatusCode != actualStatusCode {
		t.Errorf("Expected HTTP status was %d. Actual HTTP status code was %d.", expectedStatusCode, actualStatusCode)
	}
	if len(actualResponse) != 1 {
		t.Errorf("Expected to find 1 entry in map of accounts. Actual count was %d.", len(actualResponse))
	}
	actualAccount := actualResponse[1]
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
