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

func (h *Handlers) validateToken(signedToken string) (*claims, error) {
	claims := &claims{}
	_, err := jwt.ParseWithClaims(signedToken, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return h.JwtSigningKey, nil
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
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	userExists, err := h.userExists(form.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	if userExists {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	hashedPassword, err := hashPassword(form.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	if err = h.Store.CreateUser(&store.User{
		Email:          form.Email,
		HashedPassword: hashedPassword,
	}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	accessToken, refreshToken, err := h.newTokenPair(form.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
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
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	user, err := h.Store.GetUserByEmail(form.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{})
		return
	}

	if !checkHashedPassword(form.Password, user.HashedPassword) {
		c.JSON(http.StatusUnauthorized, gin.H{})
		return
	}

	accessToken, refreshToken, err := h.newTokenPair(form.Email)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(200, gin.H{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
}

func (h *Handlers) RefreshHandler(c *gin.Context) {
	var form refreshForm
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	claims, err := h.validateToken(form.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{})
		return
	}

	accessToken, refreshToken, err := h.newTokenPair(claims.Email)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(200, gin.H{
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
