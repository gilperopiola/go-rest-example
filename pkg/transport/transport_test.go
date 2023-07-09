package transport

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gilperopiola/go-rest-example/pkg/entities"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMakeSignupRequest(t *testing.T) {

	type SignupRequest struct {
		Email          any
		Username       any
		Password       any
		RepeatPassword any `json:"repeat_password"`
	}

	tests := []struct {
		name    string
		request SignupRequest
		want    entities.SignupRequest
		wantErr error
	}{
		/*{
			name:    "error_binding_request",
			request: SignupRequest{Email: 5},
			want:    entities.SignupRequest{},
			wantErr: entities.ErrBindingRequest,
		},
		{
			name:    "error_validating_request",
			request: SignupRequest{Email: "invalid"},
			want:    entities.SignupRequest{},
			wantErr: entities.ErrAllFieldsRequired,
		},*/
		{
			name: "success",
			request: SignupRequest{
				Username:       "test",
				Email:          "test@email.com",
				Password:       "test",
				RepeatPassword: "test",
			},
			want: entities.SignupRequest{
				Username:       "test",
				Email:          "test@email.com",
				Password:       "test",
				RepeatPassword: "test",
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// prepare
			jsonData, _ := json.Marshal(tt.request)
			req, _ := http.NewRequest("POST", "/signup", bytes.NewBuffer(jsonData))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = req

			// act
			signupRequest, err := makeSignupRequest(c)

			// assert
			assert.Equal(t, tt.want, signupRequest)
			assert.ErrorIs(t, err, tt.wantErr)
		})
	}
}
