package users

import (
	"errors"
	"time"
)

type User struct {
	id               int64
	name             string
	registrationDate time.Time
	lastLogin        time.Time
	karma            int64
	postsCount       int64
}

func NewUser(id int64, name string) (*User, error) {
	if err := validateUserData(id, name); err != nil {
		return nil, err
	}

	return &User{
		id:               id,
		name:             name,
		registrationDate: time.Now(),
		lastLogin:        time.Now(),
		karma:            0,
		postsCount:       0,
	}, nil
}

func (u User) Id() int64 {
	return u.id
}

func (u User) Name() string {
	return u.name
}

func (u User) RegistrationDate() time.Time {
	return u.registrationDate
}

func (u User) LastLogin() time.Time {
	return u.lastLogin
}

func (u User) Karma() int64 {
	return u.karma
}

func (u User) PostsCount() int64 {
	return u.postsCount
}

func (u *User) UpdateKarma(delta int64) {
	u.karma += delta
}

func UnmarshalFromRepository(id int64, name string,
	regDate, lastLogin time.Time, karma, postsCount int64) (*User, error) {
	user, err := NewUser(id, name)
	if err != nil {
		return nil, err
	}

	user.registrationDate = regDate
	user.lastLogin = lastLogin
	user.karma = karma
	user.postsCount = postsCount

	return user, nil
}

func validateUserData(id int64, name string) error {
	if id < 0 {
		return errors.New("Invalid user id")
	}
	if name == "" {
		return errors.New("Empty user name")
	}

	return nil
}
