package main

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/johynpapin/cruciforme/server/handlers"
)

func AuthRequiredMiddleware(jwtSigningKey []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := c.GetHeader("Authorization")

		if accessToken == "" {
			handlers.AuthUnauthorizedError.AbortContext(c)
			return
		}

		if _, err := handlers.ValidateToken(accessToken, jwtSigningKey); err != nil {
			if validationError, ok := err.(*jwt.ValidationError); ok {
				if validationError.Errors&jwt.ValidationErrorExpired != 0 {
					handlers.AuthAccessTokenExpiredError.AbortContext(c)
					return
				}
			}

			handlers.AuthUnauthorizedError.AbortContext(c)
			return
		}
	}
}

func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		lastPublicError := c.Errors.ByType(gin.ErrorTypePublic).Last()

		if lastPublicError != nil {
			if apiError, ok := lastPublicError.Err.(handlers.APIError); ok {
				c.JSON(apiError.Status(), gin.H{
					"error": apiError.Code(),
				})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "internal",
				})
			}

			return
		}

		lastPrivateError := c.Errors.ByType(gin.ErrorTypePrivate).Last()

		if lastPrivateError != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "internal",
			})
		}
	}
}
