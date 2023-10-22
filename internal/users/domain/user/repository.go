package users

import "fmt"

type Repository interface {
	AddUser(user *User) error
	GetUser(userId int64) (*User, error)
	GetKarma(userId int64) (int64, error)
	UpdateUser(userId int64, updateFn func(user *User) (*User, error)) error
	FollowUser(userId, targetId int64) error
	UnfollowUser(userId, targetId int64) error
	GetFollowing(userId int64) ([]*User, error)
	GetFollowers(userId int64) ([]*User, error)
}

type UserNotFoundError struct {
	UserId int64
}

func (e UserNotFoundError) Error() string {
	return fmt.Sprintf("User %d not found", e.UserId)
}
