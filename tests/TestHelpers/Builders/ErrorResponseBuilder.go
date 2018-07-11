package builders

import "../../../src/Models/Responses"

type ErrorResponseBuilder struct {
	Code        int
	Message     string
	Description string
}

func (builder *ErrorResponseBuilder) WithCode(code int) *ErrorResponseBuilder {
	builder.Code = code
	return builder
}

func (builder *ErrorResponseBuilder) WithMessage(message string) *ErrorResponseBuilder {
	builder.Message = message
	return builder
}

func (builder *ErrorResponseBuilder) WithDescription(description string) *ErrorResponseBuilder {
	builder.Description = description
	return builder
}

func (builder *ErrorResponseBuilder) Build() *responses.ErrorResponse {
	return &responses.ErrorResponse{Code: builder.Code, Message: builder.Message, Description: builder.Description}
}
