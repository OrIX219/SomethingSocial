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

type UserSet map[int64]struct{}

func (set *UserSet) Add(user int64) {
	var void struct{}
	(*set)[user] = void
}

func (set *UserSet) Remove(user int64) {
	delete(*set, user)
}

func (set *UserSet) Exists(user int64) bool {
	if _, ok := (*set)[user]; ok {
		return true
	}
	return false
}

type UsersInMemoryRepository struct {
	users     map[int64]UserModel
	following map[int64]*UserSet
	followers map[int64]*UserSet
}

func NewUsersInMemoryRepository(capacity int) *UsersInMemoryRepository {
	return &UsersInMemoryRepository{
		users:     make(map[int64]UserModel, capacity),
		following: make(map[int64]*UserSet, capacity),
		followers: make(map[int64]*UserSet, capacity),
	}
}

func (r *UsersInMemoryRepository) AddUser(user *users.User) error {
	if user == nil {
		return errors.New("nil user")
	}

	userModel := r.marshalUser(user)
	r.users[userModel.Id] = userModel
	r.following[userModel.Id] = new(UserSet)
	r.followers[userModel.Id] = new(UserSet)

	return nil
}

func (r *UsersInMemoryRepository) GetUser(userId int64) (*users.User, error) {
	user, ok := r.users[userId]
	if !ok {
		return nil, users.UserNotFoundError{
			UserId: userId,
		}
	}

	return r.unmarshalUser(user)
}

func (r *UsersInMemoryRepository) GetKarma(userId int64) (int64, error) {
	user, ok := r.users[userId]
	if !ok {
		return 0, users.UserNotFoundError{
			UserId: userId,
		}
	}

	return user.Karma, nil
}

func (r *UsersInMemoryRepository) UpdateUser(userId int64,
	updateFn func(user *users.User) (*users.User, error)) error {
	user, err := r.GetUser(userId)
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

func (r *UsersInMemoryRepository) FollowUser(userId, targetId int64) error {
	if _, ok := r.users[targetId]; !ok {
		return users.UserNotFoundError{
			UserId: targetId,
		}
	}

	r.following[userId].Add(targetId)
	r.followers[targetId].Add(userId)

	return nil
}

func (r *UsersInMemoryRepository) UnfollowUser(userId, targetId int64) error {
	if _, ok := r.users[targetId]; !ok {
		return users.UserNotFoundError{
			UserId: targetId,
		}
	}

	r.following[userId].Remove(targetId)
	r.followers[targetId].Remove(userId)

	return nil
}

func (r *UsersInMemoryRepository) GetFollowing(userId int64) ([]*users.User, error) {
	if _, ok := r.users[userId]; !ok {
		return nil, users.UserNotFoundError{
			UserId: userId,
		}
	}

	following := make([]*users.User, 0, len(*r.following[userId]))
	for id := range *r.following[userId] {
		user, err := r.unmarshalUser(r.users[id])
		if err == nil {
			following = append(following, user)
		}
	}

	return following, nil
}

func (r *UsersInMemoryRepository) GetFollowers(userId int64) ([]*users.User, error) {
	if _, ok := r.users[userId]; !ok {
		return nil, users.UserNotFoundError{
			UserId: userId,
		}
	}

	followers := make([]*users.User, 0, len(*r.followers[userId]))
	for id := range *r.followers[userId] {
		user, err := r.unmarshalUser(r.users[id])
		if err == nil {
			followers = append(followers, user)
		}
	}

	return followers, nil
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
