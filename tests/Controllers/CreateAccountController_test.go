package controllers

import (
	"net/http"
	"testing"

	"../../src/Models"
	"../../src/Models/Requests"
	"../../src/Models/Responses"
	"../TestHelpers/Builders"
	"../TestHelpers/Mocks"
)

/************************************************************************
*	As it turns out, it's not possible to test Gin-based controllers.	*
*	However, if it were, the below tests would suffice.					*
************************************************************************/

func TestCreateAccount_Given500AndErrorResponseFromManager_ShouldReturn500AndErrorResponseJson(t *testing.T) {
	// Arrange
	requestBuilder := builders.CreateAccountRequestBuilder{}
	contextBuilder := builders.ContextBuilder{}
	// responseBuilder := builders.ErrorResponseBuilder{}

	request := requestBuilder.WithUsername("test@test.com").WithEmail("test@test.com").WithFirstName("Simon").WithSurname("Headley").WithPassword("Test1234").Build()
	context := contextBuilder.WithCreateAccountRequest(request).Build()

	// expectedStatusCode := http.StatusInternalServerError
	// expectedResponse := responseBuilder.WithCode(3).WithMessage("Internal Server Error.").WithDescription("Failed to create account.").Build()

	mockedAccountManager := mocks.AccountManagerMock{
		CreateAccountFunc: func(request requests.CreateAccountRequest) (int, interface{}) {
			errorResponse := responses.ErrorResponse{Code: 3, Message: "Internal Server Error.", Description: "Failed to create account."}
			return http.StatusInternalServerError, &errorResponse
		},
	}

	controller := getSystemUnderTestAccountController(mockedAccountManager)

	// Act & Assert
	controller.CreateAccount(context)

	// Ideally we would like to assert something here about the context or the API response by the writermem responseWriter property on the Context struct is private.
}

func TestCreateAccount_Given201AndAccountFromManager_ShouldReturn201AndAccountJson(t *testing.T) {
	// Arrange
	requestBuilder := builders.CreateAccountRequestBuilder{}
	contextBuilder := builders.ContextBuilder{}
	// responseBuilder := builders.ErrorResponseBuilder{}

	request := requestBuilder.WithUsername("test@test.com").WithEmail("test@test.com").WithFirstName("Simon").WithSurname("Headley").WithPassword("Test1234").Build()
	context := contextBuilder.WithCreateAccountRequest(request).Build()

	// expectedStatusCode := http.StatusInternalServerError
	// expectedResponse := responseBuilder.WithCode(3).WithMessage("Internal Server Error.").WithDescription("Failed to create account.").Build()

	mockedAccountManager := mocks.AccountManagerMock{
		CreateAccountFunc: func(request requests.CreateAccountRequest) (int, interface{}) {
			account := models.Account{AccountID: 1, Login: request.Email, Password: request.Password, FirstName: request.FirstName, Surname: request.Surname, Email: request.Email}
			return http.StatusCreated, &account
		},
	}

	controller := getSystemUnderTestAccountController(mockedAccountManager)

	// Act & Assert
	controller.CreateAccount(context)

	// Ideally we would like to assert something here about the context or the API response by the writermem responseWriter property on the Context struct is private.
}
