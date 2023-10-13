package adapters

import (
	auth "github.com/OrIX219/SomethingSocial/internal/auth/domain/user"
)

type UserModel struct {
	Id       int64
	Username string
	Password string
}

type UsersInMemoryRepository struct {
	users []UserModel
}

func NewUsersInMemoryRepository(capacity int) *UsersInMemoryRepository {
	return &UsersInMemoryRepository{
		users: make([]UserModel, 0, capacity),
	}
}

func (r *UsersInMemoryRepository) AddUser(user *auth.User) (int64, error) {
	id := int64(len(r.users))

	userModel := r.marshalUser(user)
	userModel.Id = id
	r.users = append(r.users, userModel)

	return id, nil
}

func (r *UsersInMemoryRepository) GetUserId(user *auth.User) (int64, error) {
	for i := range r.users {
		if r.users[i].Username == user.Username() &&
			r.users[i].Password == user.Password() {
			return r.users[i].Id, nil
		}
	}

	return -1, auth.UserNotFoundError{
		Username: user.Username(),
		Password: user.Password(),
	}
}

func (r *UsersInMemoryRepository) GetUserByUsername(username string) (*auth.User, error) {
	for i := range r.users {
		if r.users[i].Username == username {
			return r.unmarshalUser(r.users[i])
		}
	}

	return nil, auth.UserNotFoundError{
		Username: username,
		Password: "",
	}
}

func (r *UsersInMemoryRepository) marshalUser(user *auth.User) UserModel {
	userModel := UserModel{
		Username: user.Username(),
		Password: user.Password(),
	}

	return userModel
}

func (r *UsersInMemoryRepository) unmarshalUser(user UserModel) (*auth.User, error) {
	return auth.UnmarshalFromRepository(user.Id, user.Username, user.Password)
}
