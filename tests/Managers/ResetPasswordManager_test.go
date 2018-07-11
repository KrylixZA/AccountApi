package managers

import (
	"net/http"
	"reflect"
	"testing"

	"../../src/Models"
	"../../src/Models/Requests"
	"../../src/Models/Responses"
	"../TestHelpers/Builders"
	"../TestHelpers/Mocks"
)

func TestResetPassword_GivenUnsuccessfulResult_ShouldReturn404AndErrorResponse(t *testing.T) {
	// Arrange
	responseBuilder := builders.ErrorResponseBuilder{}
	requestBuilder := builders.ResetPasswordRequestBuilder{}

	request := requestBuilder.WithCurrentPassword("password1234").WithNewPassword("test1234").Build()

	expectedStatusCode := http.StatusNotFound
	expectedResponse := responseBuilder.WithCode(4).WithMessage("Not found.").WithDescription("Failed to reset account password. Could not find account to reset the password for.").Build()

	mockedAccountDataAccess := mocks.AccountDataAccessMock{
		ResetPasswordFunc: func(request requests.ResetPasswordRequest) (*models.Account, bool) {
			return nil, false
		},
	}
	manager := getSystemUnderTestAccountManager(mockedAccountDataAccess)

	// Act
	actualStatusCode, actualResponseInterface := manager.ResetPassword(1, *request)

	// Assert
	actualResponse := actualResponseInterface.(*responses.ErrorResponse)

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

func TestResetPassword_GivenSuccessfulResult_ShouldReturn200AndAccountDetails(t *testing.T) {
	// Arrange
	requestBuilder := builders.ResetPasswordRequestBuilder{}
	accountBuilder := builders.AccountBuilder{}

	request := requestBuilder.WithCurrentPassword("password1234").WithNewPassword("test1234").Build()

	expectedStatusCode := http.StatusOK
	expectedAccount := accountBuilder.WithAccountID(1).WithLogin("test@test.com").WithPassword("Test1234").WithFirstName("Simon").WithSurname("Headley").WithEmail("test@test.com").Build()

	mockedAccountDataAccess := mocks.AccountDataAccessMock{
		ResetPasswordFunc: func(request requests.ResetPasswordRequest) (*models.Account, bool) {
			account := models.Account{AccountID: 1, Login: "test@test.com", Password: "Test1234", FirstName: "Simon", Surname: "Headley", Email: "test@test.com"}
			return &account, true
		},
	}
	manager := getSystemUnderTestAccountManager(mockedAccountDataAccess)

	// Act
	actualStatusCode, actualResponseInterface := manager.ResetPassword(1, *request)

	// Assert
	actualAccount := actualResponseInterface.(*models.Account)

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
