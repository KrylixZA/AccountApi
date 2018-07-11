package builders

import (
	"bytes"
	"encoding/json"
	"net/http"

	"../../../src/Models/Requests"
	"github.com/gin-gonic/gin"
)

type ContextBuilder struct {
	Request *http.Request
	Params  gin.Params
}

func (builder *ContextBuilder) WithCreateAccountRequest(request *requests.CreateAccountRequest) *ContextBuilder {
	jsonRequest, _ := json.Marshal(request)
	byteRequest := []byte(jsonRequest)
	builder.Request, _ = http.NewRequest("POST", "someTestUrl", bytes.NewReader(byteRequest))
	return builder
}

func (builder *ContextBuilder) WithResetPasswordRequest(request *requests.ResetPasswordRequest) *ContextBuilder {
	jsonRequest, _ := json.Marshal(request)
	byteRequest := []byte(jsonRequest)
	builder.Request, _ = http.NewRequest("POST", "someTestUrl", bytes.NewReader(byteRequest))
	return builder
}

func (builder *ContextBuilder) WithParam(key string, value string) *ContextBuilder {
	newIndex := len(builder.Params) + 1
	builder.Params[newIndex] = gin.Param{Key: key, Value: value}
	return builder
}

func (builder *ContextBuilder) WithParams(params gin.Params) *ContextBuilder {
	builder.Params = params
	return builder
}

func (builder *ContextBuilder) Build() *gin.Context {
	return &gin.Context{Params: builder.Params, Request: builder.Request}
}
