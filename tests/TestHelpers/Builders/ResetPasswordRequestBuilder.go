package builders

import "../../../src/Models/Requests"

type ResetPasswordRequestBuilder struct {
	AccountID       int
	CurrentPassword string
	NewPassword     string
}

func (builder *ResetPasswordRequestBuilder) WithAccountID(accountID int) *ResetPasswordRequestBuilder {
	builder.AccountID = accountID
	return builder
}

func (builder *ResetPasswordRequestBuilder) WithCurrentPassword(currentPassword string) *ResetPasswordRequestBuilder {
	builder.CurrentPassword = currentPassword
	return builder
}

func (builder *ResetPasswordRequestBuilder) WithNewPassword(newPassword string) *ResetPasswordRequestBuilder {
	builder.NewPassword = newPassword
	return builder
}

func (builder *ResetPasswordRequestBuilder) Build() requests.ResetPasswordRequest {
	return requests.ResetPasswordRequest{AccountID: builder.AccountID, CurrentPassword: builder.CurrentPassword, NewPassword: builder.NewPassword}
}
