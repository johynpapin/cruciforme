package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/johynpapin/cruciforme/server/store"
)

type userVerificationRequest struct {
	VerificationToken string `json:"verificationToken"`
}

func (h *Handlers) UserVerificationHandler(c *gin.Context) {
	var request userVerificationRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		NewAPIError(InvalidPayloadError, err).AbortContext(c)
		return
	}

	email, err := h.Store.UseToken(request.VerificationToken)
	if errors.Is(err, store.TokenExpiredError) {
		NewAPIError(UsersVerificationTokenExpiredError, fmt.Errorf("the verification token has expired, you need to get another one")).AbortContext(c)
		return
	} else if errors.Is(err, store.TokenInvalidError) {
		NewAPIError(UsersVerificationTokenInvalidError, fmt.Errorf("this verification token cannot be used to verify this resource")).AbortContext(c)
		return
	} else if err != nil {
		c.Error(fmt.Errorf("could not check the verification token: %w", err))
		c.Abort()
		return
	}

	user, err := h.Store.GetUserByEmail(email)
	if err != nil {
		c.Error(fmt.Errorf("could not get the user by email: %w", err))
		c.Abort()
		return
	}

	user.Verified = true

	if err = h.Store.PutUser(user); err != nil {
		c.Error(fmt.Errorf("could not update the user: %w", err))
		c.Abort()
		return
	}

	accessToken, refreshToken, err := h.newTokenPair(email)
	if err != nil {
		c.Error(fmt.Errorf("could not create a new token pair: %w", err))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
}
