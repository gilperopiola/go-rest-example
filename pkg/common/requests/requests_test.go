package requests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gilperopiola/go-rest-example/pkg/common"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const (
	validUsername = "valid_username"
	validEmail    = "test@email.com"
	validPassword = "password"
)

func TestMakeSignupRequest(t *testing.T) {
	type SignupBody struct {
		Email          any `json:"email"`
		Username       any `json:"username"`
		Password       any `json:"password"`
		RepeatPassword any `json:"repeat_password"`
	}

	successBody := SignupBody{
		Username:       validUsername,
		Email:          validEmail,
		Password:       validPassword,
		RepeatPassword: validPassword,
	}

	successRequest := SignupRequest{
		Username:       validUsername,
		Email:          validEmail,
		Password:       validPassword,
		RepeatPassword: validPassword,
	}

	tests := []struct {
		name    string
		body    SignupBody
		want    *SignupRequest
		wantErr error
	}{
		{
			name:    "error_binding_request",
			body:    SignupBody{Email: 5},
			want:    nil,
			wantErr: common.ErrBindingRequest,
		},
		{
			name:    "error_validating_request",
			body:    SignupBody{Email: "invalid"},
			want:    nil,
			wantErr: common.ErrAllFieldsRequired,
		},
		{
			name:    "success",
			body:    successBody,
			want:    &successRequest,
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// Prepare
			context := makeTestContextWithHTTPRequest(tt.body, "")

			// Act
			got, err := MakeRequest(context, &SignupRequest{})

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
		UsernameOrEmail: validUsername,
		Password:        validPassword,
	}

	successResponse := LoginRequest{
		UsernameOrEmail: validUsername,
		Password:        validPassword,
	}

	tests := []struct {
		name    string
		body    LoginBody
		want    *LoginRequest
		wantErr error
	}{
		{
			name:    "error_binding_request",
			body:    LoginBody{UsernameOrEmail: 5},
			want:    nil,
			wantErr: common.ErrBindingRequest,
		},
		{
			name:    "error_validating_request",
			body:    LoginBody{UsernameOrEmail: "invalid"},
			want:    nil,
			wantErr: common.ErrAllFieldsRequired,
		},
		{
			name:    "success",
			body:    successBody,
			want:    &successResponse,
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// Prepare
			context := makeTestContextWithHTTPRequest(tt.body, "")

			// Act
			got, err := MakeRequest(context, &LoginRequest{})

			// Assert
			assert.Equal(t, tt.want, got)
			assert.ErrorIs(t, err, tt.wantErr)
		})
	}
}

func TestMakeGetUserRequest(t *testing.T) {
	tests := []struct {
		name      string
		ctxUserID int
		urlUserID string
		want      *GetUserRequest
		wantErr   error
	}{
		{
			name:      "success",
			ctxUserID: 1,
			urlUserID: "1",
			want:      &GetUserRequest{UserID: 1},
			wantErr:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// Prepare
			context := makeTestContextWithHTTPRequest(GetUserRequest{}, "")
			addValueAndParamToContext(context, contextUserIDKey, tt.ctxUserID, pathUserIDKey, tt.urlUserID)

			// Act
			got, err := MakeRequest(context, &GetUserRequest{})

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

	successBody := UpdateUserBody{Username: validUsername}
	successResponse := UpdateUserRequest{UserID: 1, Username: validUsername}

	tests := []struct {
		name    string
		body    UpdateUserBody
		want    *UpdateUserRequest
		wantErr error
	}{
		{
			name:    "error_binding_request",
			body:    UpdateUserBody{Username: 5},
			want:    nil,
			wantErr: common.ErrBindingRequest,
		},
		{
			name:    "error_validating_request",
			body:    UpdateUserBody{Username: validUsername, Email: "invalid"},
			want:    nil,
			wantErr: common.ErrInvalidEmailFormat,
		},
		{
			name:    "success",
			body:    successBody,
			want:    &successResponse,
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// Prepare
			context := makeTestContextWithHTTPRequest(tt.body, "")
			addValueAndParamToContext(context, contextUserIDKey, 1, pathUserIDKey, "1")

			// Act
			got, err := MakeRequest(context, &UpdateUserRequest{})

			// Assert
			assert.Equal(t, tt.want, got)
			assert.ErrorIs(t, err, tt.wantErr)
		})
	}
}

func TestMakeSearchUsersRequest(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		want    *SearchUsersRequest
		wantErr error
	}{
		{
			name:    "default_values",
			want:    &SearchUsersRequest{Username: "", Page: 0, PerPage: 10},
			wantErr: nil,
		},
		{
			name:    "error_invalid_value",
			path:    "/users?username=john&page=0&per_page=",
			want:    nil,
			wantErr: common.Wrap("request.Build", common.ErrInvalidValue("per_page")),
		},
		{
			name:    "success",
			path:    "/users?username=john&page=1&per_page=20",
			want:    &SearchUsersRequest{Username: "john", Page: 1, PerPage: 20},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// Prepare
			context := makeTestContextWithHTTPRequest(SearchUsersRequest{}, tt.path)

			// Act
			got, err := MakeRequest(context, &SearchUsersRequest{})

			// Assert
			assert.Equal(t, tt.want, got)
			if tt.wantErr != nil {
				assert.Error(t, err, tt.wantErr)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestMakeDeleteUserRequest(t *testing.T) {
	tests := []struct {
		name      string
		ctxUserID int
		want      *DeleteUserRequest
		wantErr   error
	}{
		{
			name:      "success",
			ctxUserID: 1,
			want:      &DeleteUserRequest{UserID: 1},
			wantErr:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// Prepare
			context := makeTestContextWithHTTPRequest(DeleteUserRequest{}, "")
			addValueAndParamToContext(context, contextUserIDKey, tt.ctxUserID, pathUserIDKey, strconv.Itoa(tt.ctxUserID))

			// Act
			got, err := MakeRequest(context, &DeleteUserRequest{})

			// Assert
			assert.Equal(t, tt.want, got)
			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

//----------------------------------------------

func makeTestHTTPRequest(body []byte, path string) *http.Request {
	req, _ := http.NewRequest("", path, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	return req
}

func addRequestToContext(request *http.Request) *gin.Context {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = request
	return c
}

func makeTestContextWithHTTPRequest(body any, path string) *gin.Context {
	jsonData, _ := json.Marshal(body)
	httpReq := makeTestHTTPRequest(jsonData, path)
	return addRequestToContext(httpReq)
}

// TODO make everything work with strings, not the ctxValue as an int
func addValueAndParamToContext(context *gin.Context, ctxKey string, ctxValue int, paramKey, paramValue string) {
	context.Set(ctxKey, ctxValue)
	context.Params = append(context.Params, gin.Param{Key: paramKey, Value: paramValue})
}
