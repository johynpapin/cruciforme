package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type APIError int

const (
	AuthBadCredentialsError APIError = iota
	AuthAccessTokenExpiredError
	AuthRefreshTokenExpiredError
	AuthUnauthorizedError
	InvalidPayloadError
	InternalError
)

func (e APIError) Status() int {
	switch e {
	case AuthBadCredentialsError, AuthAccessTokenExpiredError, AuthRefreshTokenExpiredError, AuthUnauthorizedError:
		return http.StatusUnauthorized
	case InvalidPayloadError:
		return http.StatusBadRequest
	case InternalError:
		fallthrough
	default:
		return http.StatusInternalServerError
	}
}

func (e APIError) Code() string {
	switch e {
	case AuthBadCredentialsError:
		return "auth-bad-credentials"
	case AuthAccessTokenExpiredError:
		return "auth-access-token-expired"
	case AuthRefreshTokenExpiredError:
		return "auth-refresh-token-expired"
	case AuthUnauthorizedError:
		return "auth-unauthorized"
	case InvalidPayloadError:
		return "invalid-payload"
	case InternalError:
		fallthrough
	default:
		return "internal"
	}
}

func (e APIError) Error() string {
	return e.Code()
}

func (e APIError) AbortContext(c *gin.Context) {
	c.Error(e).SetType(gin.ErrorTypePublic)
	c.Abort()
}
