package users

import "fmt"

type Repository interface {
	AddUser(user *User) error
	GetUserById(userId int64) (*User, error)
	UpdateUser(userId int64, updateFn func(user *User) (*User, error)) error
}

type UserNotFoundError struct {
	UserId int64
}

func (e UserNotFoundError) Error() string {
	return fmt.Sprintf("User %d not found", e.UserId)
}
