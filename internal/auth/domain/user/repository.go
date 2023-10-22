package auth

import "fmt"

type Repository interface {
	AddUser(user *User) (int64, error)
	GetUserId(user *User) (int64, error)
	GetUserByUsername(username string) (*User, error)
}

type UserNotFoundError struct {
	Username string
	Password string
}

func (e UserNotFoundError) Error() string {
	return fmt.Sprintf("User %s:%s not found", e.Username, e.Password)
}

type UsernameExistsError struct {
	Username string
}

func (e UsernameExistsError) Error() string {
	return fmt.Sprintf("Username '%s' already exists", e.Username)
}
