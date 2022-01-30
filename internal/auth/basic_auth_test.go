package auth_test

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/mock"
	"messages_api/internal/auth"
	"messages_api/internal/test"
	"messages_api/internal/users/model"
	"messages_api/internal/users/repository/mocks"
	"net/http"
	"testing"
)

func TestBasicAuth(t *testing.T) {
	tests := []struct {
		name               string
		expectedStatusCode int
		request            *http.Request
		getAdminUser       *model.User
		getAdminError      error
	}{
		{name: "missing authorization header", request: &http.Request{Header: map[string][]string{}}, expectedStatusCode: http.StatusUnauthorized},
		{
			name:               "repository returns error",
			request:            &http.Request{Header: validAuthorizationHeader("user", "pass")},
			getAdminError:      errors.New("error"),
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			name:               "repository did not return a user",
			request:            &http.Request{Header: validAuthorizationHeader("user", "pass")},
			expectedStatusCode: http.StatusUnauthorized,
		},
		{
			name:               "credentials did not match",
			request:            &http.Request{Header: validAuthorizationHeader("user", "pass")},
			getAdminUser:       &model.User{Password: "pass2"},
			expectedStatusCode: http.StatusUnauthorized,
		},
		{
			name:               "authorized",
			request:            &http.Request{Header: validAuthorizationHeader("user", "pass")},
			getAdminUser:       &model.User{Password: "pass"},
			expectedStatusCode: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repositoryMock := &mocks.Repository{}
			repositoryMock.
				On("GetAdmin", mock.Anything, mock.Anything).
				Return(tt.getAdminUser, tt.getAdminError)

			middleware := auth.BasicAuth(repositoryMock)

			writterMock := test.NewResponseWritterMock()
			ginCtx := &gin.Context{
				Request: tt.request,
				Writer:  writterMock,
			}
			middleware(ginCtx)

			assert.Equal(t, writterMock.StatusCode, tt.expectedStatusCode)
		})
	}
}

func validAuthorizationHeader(username, password string) map[string][]string {
	authHeader := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", username, password)))
	return map[string][]string{
		"Authorization": {
			fmt.Sprintf("Basic %s", authHeader),
		},
	}
}
