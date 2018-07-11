package builders

import "../../../src/Models/Requests"

type CreateAccountRequestBuilder struct {
	Username  string
	Email     string
	FirstName string
	Surname   string
	Password  string
}

func (builder *CreateAccountRequestBuilder) WithUsername(username string) *CreateAccountRequestBuilder {
	builder.Username = username
	return builder
}

func (builder *CreateAccountRequestBuilder) WithEmail(email string) *CreateAccountRequestBuilder {
	builder.Email = email
	return builder
}

func (builder *CreateAccountRequestBuilder) WithFirstName(firstName string) *CreateAccountRequestBuilder {
	builder.FirstName = firstName
	return builder
}

func (builder *CreateAccountRequestBuilder) WithSurname(surname string) *CreateAccountRequestBuilder {
	builder.Surname = surname
	return builder
}

func (builder *CreateAccountRequestBuilder) WithPassword(password string) *CreateAccountRequestBuilder {
	builder.Password = password
	return builder
}

func (builder *CreateAccountRequestBuilder) Build() *requests.CreateAccountRequest {
	return &requests.CreateAccountRequest{Username: builder.Username, Email: builder.Email, FirstName: builder.FirstName, Surname: builder.Surname, Password: builder.Password}
}
