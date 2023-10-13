package auth

import "errors"

type User struct {
	id       int64
	username string
	password string
}

func NewUser(id int64, username, password string) (*User, error) {
	if id < 0 {
		return nil, errors.New("Invalid user id")
	}
	if username == "" {
		return nil, errors.New("Empty user username")
	}
	if password == "" {
		return nil, errors.New("Empty user password")
	}

	return &User{
		username: username,
		password: password,
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
	user, err := NewUser(id, username, password)
	if err != nil {
		return nil, err
	}

	return user, nil
}
