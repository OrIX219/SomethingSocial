package adapters

import (
	"errors"
	"time"

	users "github.com/OrIX219/SomethingSocial/internal/users/domain/user"
)

type UserModel struct {
	Id               int64
	Name             string
	RegistrationDate time.Time
	LastLogin        time.Time
	Karma            int64
	PostsCount       int64
}

type UsersInMemoryRepository struct {
	users map[int64]UserModel
}

func NewUsersInMemoryRepository(capacity int) *UsersInMemoryRepository {
	return &UsersInMemoryRepository{
		users: make(map[int64]UserModel, capacity),
	}
}

func (r *UsersInMemoryRepository) AddUser(user *users.User) error {
	if user == nil {
		return errors.New("nil user")
	}

	userModel := r.marshalUser(user)
	r.users[userModel.Id] = userModel

	return nil
}

func (r *UsersInMemoryRepository) GetUserById(userId int64) (*users.User, error) {
	user, ok := r.users[userId]
	if !ok {
		return nil, users.UserNotFoundError{
			UserId: userId,
		}
	}

	return r.unmarshalUser(user)
}

func (r *UsersInMemoryRepository) UpdateUser(userId int64,
	updateFn func(user *users.User) (*users.User, error)) error {
	user, err := r.GetUserById(userId)
	if err != nil {
		return err
	}

	updatedUser, err := updateFn(user)
	if err != nil {
		return err
	}

	r.users[updatedUser.Id()] = r.marshalUser(updatedUser)

	return nil
}

func (r *UsersInMemoryRepository) marshalUser(user *users.User) UserModel {
	return UserModel{
		Id:               user.Id(),
		Name:             user.Name(),
		RegistrationDate: user.RegistrationDate(),
		LastLogin:        user.LastLogin(),
		Karma:            user.Karma(),
		PostsCount:       user.PostsCount(),
	}
}

func (r *UsersInMemoryRepository) unmarshalUser(user UserModel) (*users.User, error) {
	return users.UnmarshalFromRepository(user.Id, user.Name,
		user.RegistrationDate, user.LastLogin, user.Karma, user.PostsCount)
}
