package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type APIErrorCode string

const (
	AuthBadCredentialsError            APIErrorCode = "auth-bad-credentials"
	AuthEmailAlreadyInUseError         APIErrorCode = "auth-email-already-in-use"
	AuthAccessTokenExpiredError        APIErrorCode = "auth-access-token-expired"
	AuthRefreshTokenExpiredError       APIErrorCode = "auth-refresh-token-expired"
	AuthUserNotVerifiedError           APIErrorCode = "auth-user-not-verified"
	AuthUnauthorizedError              APIErrorCode = "auth-unauthorized"
	UsersVerificationTokenExpiredError APIErrorCode = "users-verification-token-expired"
	UsersVerificationTokenInvalidError APIErrorCode = "users-verification-token-invalid"
	InvalidPayloadError                APIErrorCode = "invalid-payload"
	InternalError                      APIErrorCode = "internal"
)

type APIError struct {
	code APIErrorCode
	err  error
}

func NewAPIError(code APIErrorCode, err error) *APIError {
	return &APIError{
		code: code,
		err:  err,
	}
}

func (e *APIError) Status() int {
	switch e.code {
	case AuthBadCredentialsError, AuthAccessTokenExpiredError, AuthRefreshTokenExpiredError, AuthUnauthorizedError, AuthUserNotVerifiedError:
		return http.StatusUnauthorized
	case InvalidPayloadError, AuthEmailAlreadyInUseError, UsersVerificationTokenExpiredError, UsersVerificationTokenInvalidError:
		return http.StatusBadRequest
	case InternalError:
		fallthrough
	default:
		return http.StatusInternalServerError
	}
}

func (e *APIError) Code() string {
	return string(e.code)
}

func (e *APIError) Error() string {
	return e.err.Error()
}

func (e *APIError) Unwrap() error {
	return e.err
}

func (e *APIError) AbortContext(c *gin.Context) {
	c.Error(e).SetType(gin.ErrorTypePublic)
	c.Abort()
}
