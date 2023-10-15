package requests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	customErrors "github.com/gilperopiola/go-rest-example/pkg/common/errors"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const (
	CTX_KEY_USER_ID   = "ID"
	PARAM_KEY_USER_ID = "user_id"

	VALID_USERNAME = "valid_username"
	VALID_EMAIL    = "test@email.com"
	VALID_PASSWORD = "password"
)

func TestMakeSignupRequest(t *testing.T) {
	type SignupBody struct {
		Email          any `json:"email"`
		Username       any `json:"username"`
		Password       any `json:"password"`
		RepeatPassword any `json:"repeat_password"`
	}

	successBody := SignupBody{
		Username:       VALID_USERNAME,
		Email:          VALID_EMAIL,
		Password:       VALID_PASSWORD,
		RepeatPassword: VALID_PASSWORD,
	}

	successResponse := SignupRequest{
		Username:       VALID_USERNAME,
		Email:          VALID_EMAIL,
		Password:       VALID_PASSWORD,
		RepeatPassword: VALID_PASSWORD,
	}

	tests := []struct {
		name    string
		body    SignupBody
		want    SignupRequest
		wantErr error
	}{
		{
			name:    "error_binding_request",
			body:    SignupBody{Email: 5},
			want:    SignupRequest{},
			wantErr: customErrors.ErrBindingRequest,
		},
		{
			name:    "error_validating_request",
			body:    SignupBody{Email: "invalid"},
			want:    SignupRequest{},
			wantErr: customErrors.ErrAllFieldsRequired,
		},
		{
			name:    "success",
			body:    successBody,
			want:    successResponse,
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// Prepare
			context := makeTestContextWithHTTPRequest(tt.body)

			// Act
			got, err := MakeSignupRequest(context)

			// Assert
			assert.Equal(t, tt.want, got)
			assert.ErrorIs(t, err, tt.wantErr)
		})
	}
}

func TestMakeLoginRequest(t *testing.T) {
	type LoginBody struct {
		UsernameOrEmail any `json:"username_or_email"`
		Password        any `json:"password"`
	}

	successBody := LoginBody{
		UsernameOrEmail: VALID_USERNAME,
		Password:        VALID_PASSWORD,
	}

	successResponse := LoginRequest{
		UsernameOrEmail: VALID_USERNAME,
		Password:        VALID_PASSWORD,
	}

	tests := []struct {
		name    string
		body    LoginBody
		want    LoginRequest
		wantErr error
	}{
		{
			name:    "error_binding_request",
			body:    LoginBody{UsernameOrEmail: 5},
			want:    LoginRequest{},
			wantErr: customErrors.ErrBindingRequest,
		},
		{
			name:    "error_validating_request",
			body:    LoginBody{UsernameOrEmail: "invalid"},
			want:    LoginRequest{},
			wantErr: customErrors.ErrAllFieldsRequired,
		},
		{
			name:    "success",
			body:    successBody,
			want:    successResponse,
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// Prepare
			context := makeTestContextWithHTTPRequest(tt.body)

			// Act
			got, err := MakeLoginRequest(context)

			// Assert
			assert.Equal(t, tt.want, got)
			assert.ErrorIs(t, err, tt.wantErr)
		})
	}
}

func TestMakeGetUserRequest(t *testing.T) {
	tests := []struct {
		name      string
		ctxUserID string
		urlUserID string
		want      GetUserRequest
		wantErr   error
	}{
		{
			name:      "error_invalid_id",
			ctxUserID: "0",
			urlUserID: "0",
			want:      GetUserRequest{},
			wantErr:   customErrors.ErrAllFieldsRequired,
		},
		{
			name:      "success",
			ctxUserID: "1",
			urlUserID: "1",
			want:      GetUserRequest{ID: 1},
			wantErr:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// Prepare
			context := makeTestContextWithHTTPRequest(GetUserRequest{})
			addValueAndParamToContext(context, CTX_KEY_USER_ID, tt.ctxUserID, PARAM_KEY_USER_ID, tt.urlUserID)

			// Act
			got, err := MakeGetUserRequest(context)

			// Assert
			assert.Equal(t, tt.want, got)
			assert.ErrorIs(t, err, tt.wantErr)
		})
	}
}

func TestMakeUpdateUserRequest(t *testing.T) {
	type UpdateUserBody struct {
		ID       any `json:"id"`
		Username any `json:"username"`
		Email    any `json:"email"`
	}

	successBody := UpdateUserBody{Username: VALID_USERNAME}
	successResponse := UpdateUserRequest{ID: 1, Username: VALID_USERNAME}

	tests := []struct {
		name      string
		ctxUserID string
		urlUserID string
		body      UpdateUserBody
		want      UpdateUserRequest
		wantErr   error
	}{
		{
			name:      "error_binding_request",
			ctxUserID: "1",
			urlUserID: "1",
			body:      UpdateUserBody{Username: 5},
			want:      UpdateUserRequest{},
			wantErr:   customErrors.ErrBindingRequest,
		},
		{
			name:      "error_validating_request",
			ctxUserID: "0",
			urlUserID: "0",
			body:      UpdateUserBody{Username: VALID_USERNAME},
			want:      UpdateUserRequest{},
			wantErr:   customErrors.ErrAllFieldsRequired,
		},
		{
			name:      "success",
			ctxUserID: "1",
			urlUserID: "1",
			body:      successBody,
			want:      successResponse,
			wantErr:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// Prepare
			context := makeTestContextWithHTTPRequest(tt.body)
			addValueAndParamToContext(context, CTX_KEY_USER_ID, tt.ctxUserID, PARAM_KEY_USER_ID, tt.urlUserID)

			// Act
			got, err := MakeUpdateUserRequest(context)

			// Assert
			assert.Equal(t, tt.want, got)
			assert.ErrorIs(t, err, tt.wantErr)
		})
	}
}

//----------------------------------------------

func makeTestHTTPRequest(body []byte) *http.Request {
	req, _ := http.NewRequest("", "", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	return req
}

func addRequestToContext(request *http.Request) *gin.Context {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = request
	return c
}

func makeTestContextWithHTTPRequest(body any) *gin.Context {
	jsonData, _ := json.Marshal(body)
	httpReq := makeTestHTTPRequest(jsonData)
	return addRequestToContext(httpReq)
}

func addValueAndParamToContext(context *gin.Context, ctxKey, ctxValue, paramKey, paramValue string) {
	context.Set(ctxKey, ctxValue)
	context.Params = append(context.Params, gin.Param{Key: paramKey, Value: paramValue})
}
