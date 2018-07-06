package builders

import requests "../../../src/Models/Requests"

type CreateAccountRequestBuilder struct {
	username  string
	email     string
	firstName string
	surname   string
	password  string
}

func (builder *CreateAccountRequestBuilder) WithUsername(username string) *CreateAccountRequestBuilder {
	builder.username = username
	return builder
}

func (builder *CreateAccountRequestBuilder) WithEmail(email string) *CreateAccountRequestBuilder {
	builder.email = email
	return builder
}

func (builder *CreateAccountRequestBuilder) WithFirstName(firstName string) *CreateAccountRequestBuilder {
	builder.firstName = firstName
	return builder
}

func (builder *CreateAccountRequestBuilder) WithSurname(surname string) *CreateAccountRequestBuilder {
	builder.surname = surname
	return builder
}

func (builder *CreateAccountRequestBuilder) WithPassword(password string) *CreateAccountRequestBuilder {
	builder.password = password
	return builder
}

func (builder *CreateAccountRequestBuilder) Build() requests.CreateAccountRequest {
	return requests.CreateAccountRequest{Username: builder.username, Email: builder.email, FirstName: builder.firstName, Surname: builder.surname, Password: builder.password}
}
