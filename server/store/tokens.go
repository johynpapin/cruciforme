package store

import (
	"errors"
	"fmt"

	"github.com/dchest/uniuri"
)

const (
	tokenLength = 20
)

var (
	TokenExpiredError = errors.New("the token has expired")
	TokenInvalidError = errors.New("this token is invalid")
)

func (s *Store) CreateToken(scope string) (string, error) {
	var token string

	if err := s.update(func(txn txn) error {
		var err error
		token, err = generateToken(txn)
		if err != nil {
			return err
		}

		return s.set([]byte("tokens/"+token), []byte(scope))
	}); err != nil {
		return "", err
	}

	return token, nil
}

func (s *Store) UseToken(token string) (string, error) {
	var storedScope []byte

	err := s.update(func(txn txn) (err error) {
		storedScope, err = txn.get([]byte("tokens/" + token))
		if err != nil {
			return err
		}
		if storedScope == nil {
			return TokenInvalidError
		}

		// TODO: check token expiration

		return nil
	})
	if err == TokenExpiredError || err == TokenInvalidError {
		return "", fmt.Errorf("%w", err)
	} else if err != nil {
		return "", err
	}

	return string(storedScope), nil
}

func generateToken(txn txn) (string, error) {
	for {
		token := uniuri.NewLen(tokenLength)

		scope, err := txn.get([]byte("tokens/" + token))
		if err != nil {
			return "", err
		}

		if scope == nil {
			return token, nil
		}
	}
}
