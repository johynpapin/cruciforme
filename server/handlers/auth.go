package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/johynpapin/cruciforme/server/store"
	"golang.org/x/crypto/bcrypt"
)

type claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

type signInForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type signUpForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type refreshForm struct {
	RefreshToken string `json:"refreshToken"`
}

func (h *Handlers) newTokenPair(email string) (string, string, error) {
	accessTokenExpiresAt := time.Now().Add(15 * time.Minute)
	refreshTokenExpiresAt := time.Now().Add(14 * 24 * time.Hour)

	accessTokenClaims := claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: accessTokenExpiresAt.Unix(),
		},
	}
	refreshTokenClaims := claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: refreshTokenExpiresAt.Unix(),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)

	accessTokenString, err := accessToken.SignedString(h.JwtSigningKey)
	if err != nil {
		return "", "", err
	}
	refreshTokenString, err := refreshToken.SignedString(h.JwtSigningKey)
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}

func ValidateToken(signedToken string, jwtSigningKey []byte) (*claims, error) {
	claims := &claims{}
	_, err := jwt.ParseWithClaims(signedToken, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return jwtSigningKey, nil
	})
	if err != nil {
		return nil, err
	}

	return claims, nil
}

func (h *Handlers) userExists(email string) (bool, error) {
	user, err := h.Store.GetUserByEmail(email)
	if err != nil {
		return false, err
	}

	return user != nil, nil
}

func (h *Handlers) SignUpHandler(c *gin.Context) {
	var form signUpForm
	if err := c.ShouldBindJSON(&form); err != nil {
		InvalidPayloadError.AbortContext(c)
		return
	}

	userExists, err := h.userExists(form.Email)
	if err != nil {
		InternalError.AbortContext(c)
		return
	}

	if userExists {
		InvalidPayloadError.AbortContext(c)
		return
	}

	hashedPassword, err := hashPassword(form.Password)
	if err != nil {
		InternalError.AbortContext(c)
		return
	}

	if err = h.Store.CreateUser(&store.User{
		Email:          form.Email,
		HashedPassword: hashedPassword,
	}); err != nil {
		InternalError.AbortContext(c)
		return
	}

	accessToken, refreshToken, err := h.newTokenPair(form.Email)
	if err != nil {
		InternalError.AbortContext(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
}

func (h *Handlers) SignInHandler(c *gin.Context) {
	var form signInForm
	if err := c.ShouldBindJSON(&form); err != nil {
		InvalidPayloadError.AbortContext(c)
		return
	}

	user, err := h.Store.GetUserByEmail(form.Email)
	if err != nil {
		InternalError.AbortContext(c)
		return
	}

	if user == nil {
		AuthBadCredentialsError.AbortContext(c)
		return
	}

	if !checkHashedPassword(form.Password, user.HashedPassword) {
		AuthBadCredentialsError.AbortContext(c)
		return
	}

	accessToken, refreshToken, err := h.newTokenPair(form.Email)
	if err != nil {
		InternalError.AbortContext(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
}

func (h *Handlers) RefreshHandler(c *gin.Context) {
	var form refreshForm
	if err := c.ShouldBindJSON(&form); err != nil {
		InvalidPayloadError.AbortContext(c)
		return
	}

	claims, err := ValidateToken(form.RefreshToken, h.JwtSigningKey)
	if err != nil {
		if validationError, ok := err.(*jwt.ValidationError); ok {
			if validationError.Errors&jwt.ValidationErrorExpired != 0 {
				AuthRefreshTokenExpiredError.AbortContext(c)
				return
			}
		}

		AuthUnauthorizedError.AbortContext(c)
		return
	}

	accessToken, refreshToken, err := h.newTokenPair(claims.Email)
	if err != nil {
		InternalError.AbortContext(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
}

func hashPassword(password string) (string, error) {
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPasswordBytes), err
}

func checkHashedPassword(password string, hashedPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}
