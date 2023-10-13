package auth

import (
	"crypto/sha1"
	"errors"
	"fmt"
)

type User struct {
	id       int64
	username string
	password string
}

func NewUser(username, password string) (*User, error) {
	if err := validateUserData(0, username, password); err != nil {
		return nil, err
	}

	return &User{
		id:       0,
		username: username,
		password: generatePasswordHash(password),
	}, nil
}

func (u User) Id() int64 {
	return u.id
}

func (u User) Username() string {
	return u.username
}

func (u User) Password() string {
	return u.password
}

func UnmarshalFromRepository(id int64, username, password string) (*User, error) {
	if err := validateUserData(id, username, password); err != nil {
		return nil, err
	}

	return &User{
		id:       id,
		username: username,
		password: password,
	}, nil
}

func validateUserData(id int64, username, password string) error {
	if id < 0 {
		return errors.New("Invalid user id")
	}
	if username == "" {
		return errors.New("Empty user username")
	}
	if password == "" {
		return errors.New("Empty user password")
	}
	return nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte("mock_salt")))
}
