package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/johynpapin/cruciforme/server/email"
	"github.com/johynpapin/cruciforme/server/store"
	"golang.org/x/crypto/bcrypt"
)

type accessTokenClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

type refreshTokenClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

type signInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type signUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type refreshRequest struct {
	RefreshToken string `json:"refreshToken"`
}

func (h *Handlers) newTokenPair(email string) (string, string, error) {
	accessTokenExpiresAt := time.Now().Add(15 * time.Minute)
	refreshTokenExpiresAt := time.Now().Add(14 * 24 * time.Hour)

	accessTokenClaims := accessTokenClaims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: accessTokenExpiresAt.Unix(),
		},
	}
	refreshTokenClaims := refreshTokenClaims{
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

func ValidateAccessToken(signedToken string, jwtSigningKey []byte) (*accessTokenClaims, error) {
	claims := &accessTokenClaims{}
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

func validateRefreshToken(signedToken string, jwtSigningKey []byte) (*refreshTokenClaims, error) {
	claims := &refreshTokenClaims{}
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
	var request signUpRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		NewAPIError(InvalidPayloadError, err).AbortContext(c)
		return
	}

	userExists, err := h.userExists(request.Email)
	if err != nil {
		c.Error(fmt.Errorf("could not check if the user exists: %w", err))
		c.Abort()
		return
	}

	if userExists {
		NewAPIError(AuthEmailAlreadyInUseError, fmt.Errorf("an account using this email address already exists")).AbortContext(c)
		return
	}

	hashedPassword, err := hashPassword(request.Password)
	if err != nil {
		c.Error(fmt.Errorf("could not hash the password: %w", err))
		c.Abort()
		return
	}

	if err = h.Store.PutUser(&store.User{
		Email:          request.Email,
		HashedPassword: hashedPassword,
	}); err != nil {
		c.Error(fmt.Errorf("could not create the user: %w", err))
		c.Abort()
		return
	}

	verificationToken, err := h.Store.CreateToken(request.Email)
	if err != nil {
		c.Error(fmt.Errorf("could not create the verification token: %w", err))
		c.Abort()
		return
	}

	if err = h.EmailSender.Send(email.EmailAddress{
		Email: request.Email,
	}, &email.VerificationEmail{
		Link: "https://crucifor.me/auth/verify?token=" + verificationToken,
	}); err != nil {
		c.Error(fmt.Errorf("could not send the verification email: %w", err))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handlers) SignInHandler(c *gin.Context) {
	var request signInRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		NewAPIError(InvalidPayloadError, err).AbortContext(c)
		return
	}

	user, err := h.Store.GetUserByEmail(request.Email)
	if err != nil {
		c.Error(fmt.Errorf("could not get the user by email: %w", err))
		c.Abort()
		return
	}

	if user == nil {
		NewAPIError(AuthBadCredentialsError, fmt.Errorf("this email is not used by any account")).AbortContext(c)
		return
	}

	if !checkHashedPassword(request.Password, user.HashedPassword) {
		NewAPIError(AuthBadCredentialsError, fmt.Errorf("the password is incorect")).AbortContext(c)
		return
	}

	if !user.Verified {
		NewAPIError(AuthUserNotVerifiedError, fmt.Errorf("please verify the email of the user before signing in")).AbortContext(c)
		return
	}

	accessToken, refreshToken, err := h.newTokenPair(request.Email)
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

func (h *Handlers) RefreshHandler(c *gin.Context) {
	var request refreshRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		NewAPIError(InvalidPayloadError, err).AbortContext(c)
		return
	}

	claims, err := validateRefreshToken(request.RefreshToken, h.JwtSigningKey)
	if err != nil {
		if validationError, ok := err.(*jwt.ValidationError); ok {
			if validationError.Errors&jwt.ValidationErrorExpired != 0 {
				NewAPIError(AuthRefreshTokenExpiredError, fmt.Errorf("the refresh token has expired, you need to sign in again")).AbortContext(c)
				return
			}
		}

		NewAPIError(AuthUnauthorizedError, fmt.Errorf("the refresh token is invalid: ", err)).AbortContext(c)
		return
	}

	accessToken, refreshToken, err := h.newTokenPair(claims.Email)
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

func hashPassword(password string) (string, error) {
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPasswordBytes), err
}

func checkHashedPassword(password string, hashedPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}
