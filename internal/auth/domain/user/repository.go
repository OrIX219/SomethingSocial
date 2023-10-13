package auth

import "fmt"

type Repository interface {
	AddUser(user *User) (int64, error)
	GetUser(username, password string) (*User, error)
	GetUserByUsername(username string) (*User, error)
	GetUserById(userId int) (*User, error)
}

type UserNotFoundError struct {
	Username string
	Password string
}

func (e UserNotFoundError) Error() string {
	return fmt.Sprintf("User %s:%s not found", e.Username, e.Password)
}

type UserIdNotFoundError struct {
	Id int
}

func (e UserIdNotFoundError) Error() string {
	return fmt.Sprintf("User with id %d not found", e.Id)
}
