package main

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"github.com/johynpapin/cruciforme/server/handlers"
)

func authRequiredMiddleware(jwtSigningKey []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := c.GetHeader("Authorization")

		if accessToken == "" {
			handlers.NewAPIError(handlers.AuthUnauthorizedError, fmt.Errorf("authentication is required but you did not send an access token")).AbortContext(c)
			return
		}

		claims, err := handlers.ValidateAccessToken(accessToken, jwtSigningKey)
		if err != nil {
			if validationError, ok := err.(*jwt.ValidationError); ok {
				if validationError.Errors&jwt.ValidationErrorExpired != 0 {
					handlers.NewAPIError(handlers.AuthAccessTokenExpiredError, fmt.Errorf("the access token has expired, you need to refresh it")).AbortContext(c)
					return
				}
			}

			handlers.NewAPIError(handlers.AuthUnauthorizedError, fmt.Errorf("the access token is invalid: %w", err)).AbortContext(c)
			return
		}

		c.Set("userId", claims.Email)

		if hub := sentrygin.GetHubFromContext(c); hub != nil {
			hub.Scope().SetUser(sentry.User{Email: claims.Email})
		}
	}
}

func errorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		lastPublicError := c.Errors.ByType(gin.ErrorTypePublic).Last()

		if lastPublicError != nil {
			if apiError, ok := lastPublicError.Err.(*handlers.APIError); ok {
				c.JSON(apiError.Status(), gin.H{
					"error": map[string]interface{}{
						"code":    apiError.Code(),
						"message": apiError.Error(),
					},
				})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": map[string]interface{}{
						"code":    "internal",
						"message": lastPublicError.Error(),
					},
				})
			}

			return
		}

		lastPrivateError := c.Errors.ByType(gin.ErrorTypePrivate).Last()

		if lastPrivateError != nil {
			if hub := sentrygin.GetHubFromContext(c); hub != nil {
				hub.CaptureException(lastPrivateError)
			}

			c.JSON(http.StatusInternalServerError, gin.H{
				"error": map[string]interface{}{
					"code":    "internal",
					"message": "an internal error occured, you can try the request again",
				},
			})
		}
	}
}
