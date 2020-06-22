package store

import (
	"encoding/json"
)

type User struct {
	Email          string
	Verified       bool
	HashedPassword string
}

func (s *Store) PutUser(user *User) error {
	userBytes, err := json.Marshal(user)
	if err != nil {
		return err
	}

	return s.set([]byte("users/"+user.Email), userBytes)
}

func (s *Store) GetUserByEmail(email string) (*User, error) {
	userBytes, err := s.get([]byte("users/" + email))
	if err != nil {
		return nil, err
	} else if userBytes == nil {
		return nil, nil
	}

	user := &User{}
	if err = json.Unmarshal(userBytes, user); err != nil {
		return nil, err
	}

	return user, nil
}
