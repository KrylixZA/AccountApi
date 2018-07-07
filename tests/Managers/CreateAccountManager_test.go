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

	request := requestBuilder.WithUsername("Test").WithEmail("test@test.com").WithFirstName("Simon").WithSurname("Headley").WithPassword("Test1234").Build()

	expectedStatusCode := http.StatusInternalServerError
	expectedResponse := responseBuilder.WithCode(3).WithMessage("Internal Server Error.").WithDescription("Failed to create account.").Build()

	mockedAccountDataAccess := &mocks.AccountDataAccessMock{
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
