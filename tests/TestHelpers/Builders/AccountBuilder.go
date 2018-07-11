package builders

import "../../../src/Models"

type AccountBuilder struct {
	AccountID int
	Login     string
	Password  string
	FirstName string
	Surname   string
	Email     string
}

func (builder *AccountBuilder) WithAccountID(accountID int) *AccountBuilder {
	builder.AccountID = accountID
	return builder
}

func (builder *AccountBuilder) WithLogin(login string) *AccountBuilder {
	builder.Login = login
	return builder
}

func (builder *AccountBuilder) WithPassword(password string) *AccountBuilder {
	builder.Password = password
	return builder
}

func (builder *AccountBuilder) WithFirstName(firstName string) *AccountBuilder {
	builder.FirstName = firstName
	return builder
}

func (builder *AccountBuilder) WithSurname(surname string) *AccountBuilder {
	builder.Surname = surname
	return builder
}

func (builder *AccountBuilder) WithEmail(email string) *AccountBuilder {
	builder.Email = email
	return builder
}

func (builder *AccountBuilder) Build() *models.Account {
	return &models.Account{AccountID: builder.AccountID, Login: builder.Login, Password: builder.Password, FirstName: builder.FirstName, Surname: builder.Surname, Email: builder.Email}
}
