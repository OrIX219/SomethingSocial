package adapters

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	users "github.com/OrIX219/SomethingSocial/internal/users/domain/user"
	"github.com/jmoiron/sqlx"
)

const (
	usersTable     = "users"
	followingTable = "following"
)

type UserModel struct {
	Id               int64     `db:"id"`
	Name             string    `db:"name"`
	RegistrationDate time.Time `db:"registration_date"`
	LastLogin        time.Time `db:"last_login"`
	Karma            int64     `db:"karma"`
	PostsCount       int64     `db:"posts_count"`
	Role             string    `db:"role"`
}

type UsersPostgresRepository struct {
	db *sqlx.DB
}

func NewUsersPostgresRepository(db *sqlx.DB) *UsersPostgresRepository {
	if db == nil {
		panic("UsersPostgresRepository nil db")
	}

	return &UsersPostgresRepository{
		db: db,
	}
}

func (r *UsersPostgresRepository) AddUser(user *users.User) error {
	if user == nil {
		return errors.New("nil user")
	}

	userModel := r.marshalUser(user)

	query := fmt.Sprintf(`INSERT INTO %s
		(id, name, registration_date, last_login, karma, posts_count)
		VALUES ($1, $2, $3, $4, $5, $6)`, usersTable)

	_, err := r.db.Exec(query, userModel.Id, userModel.Name,
		userModel.RegistrationDate, userModel.LastLogin,
		userModel.Karma, userModel.PostsCount)

	return err
}

func (r *UsersPostgresRepository) GetUser(userId int64) (*users.User, error) {
	var userModel UserModel
	query := fmt.Sprintf(`SELECT * FROM %s WHERE id = $1`, usersTable)
	err := r.db.Get(&userModel, query, userId)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, users.UserNotFoundError{
				UserId: userId,
			}
		default:
			return nil, err
		}
	}

	return r.unmarshalUser(userModel)
}

func (r *UsersPostgresRepository) GetKarma(userId int64) (int64, error) {
	var karma int64
	query := fmt.Sprintf(`SELECT karma FROM %s WHERE id = $1`, usersTable)
	err := r.db.Get(&karma, query, userId)
	if err == sql.ErrNoRows {
		return 0, users.UserNotFoundError{
			UserId: userId,
		}
	}

	return karma, nil
}

func (r *UsersPostgresRepository) UpdateUser(userId int64,
	updateFn func(user *users.User) (*users.User, error)) error {
	user, err := r.GetUser(userId)
	if err != nil {
		return err
	}

	updatedUser, err := updateFn(user)
	if err != nil {
		return err
	}

	query := fmt.Sprintf(`UPDATE %s SET name=$1, last_login=$2, karma=$3,
		posts_count=$4 role=$5 WHERE id=$6`, usersTable)
	res, err := r.db.Exec(query, updatedUser.Name(), updatedUser.LastLogin(),
		updatedUser.Karma(), updatedUser.PostsCount(),
		updatedUser.Role().String(), userId)
	if err != nil {
		return err
	}

	if rows, _ := res.RowsAffected(); rows == 0 {
		return users.UserNotFoundError{
			UserId: userId,
		}
	}

	return nil
}

func (r *UsersPostgresRepository) FollowUser(userId, targetId int64) error {
	query := fmt.Sprintf(`INSERT INTO %s (follower_id, follow_id) VALUES ($1, $2)`, followingTable)
	_, err := r.db.Exec(query, userId, targetId)

	return err
}

func (r *UsersPostgresRepository) UnfollowUser(userId, targetId int64) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE follower_id=$1 AND follow_id=$2`,
		followingTable)
	_, err := r.db.Exec(query, userId, targetId)

	return err
}

func (r *UsersPostgresRepository) GetFollowing(userId int64) ([]*users.User, error) {
	var userModels []UserModel
	query := fmt.Sprintf(`SELECT u.* FROM %s u
		INNER JOIN %s f ON u.id=f.follow_id	WHERE f.follower_id=$1`,
		usersTable, followingTable)
	err := r.db.Select(&userModels, query, userId)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, users.UserNotFoundError{
				UserId: userId,
			}
		default:
			return nil, err
		}
	}

	following := make([]*users.User, 0, len(userModels))
	for i := range userModels {
		user, err := r.unmarshalUser(userModels[i])
		if err == nil {
			following = append(following, user)
		}
	}

	return following, nil
}

func (r *UsersPostgresRepository) GetFollowers(userId int64) ([]*users.User, error) {
	var userModels []UserModel
	query := fmt.Sprintf(`SELECT u.* FROM %s u
		INNER JOIN %s f ON u.id=f.follower_id WHERE f.follow_id=$1`,
		usersTable, followingTable)
	err := r.db.Select(&userModels, query, userId)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, users.UserNotFoundError{
				UserId: userId,
			}
		default:
			return nil, err
		}
	}

	followers := make([]*users.User, 0, len(userModels))
	for i := range userModels {
		user, err := r.unmarshalUser(userModels[i])
		if err == nil {
			followers = append(followers, user)
		}
	}

	return followers, nil
}

func (r *UsersPostgresRepository) marshalUser(user *users.User) UserModel {
	return UserModel{
		Id:               user.Id(),
		Name:             user.Name(),
		RegistrationDate: user.RegistrationDate(),
		LastLogin:        user.LastLogin(),
		Karma:            user.Karma(),
		PostsCount:       user.PostsCount(),
		Role:             user.Role().String(),
	}
}

func (r *UsersPostgresRepository) unmarshalUser(user UserModel) (*users.User, error) {
	return users.UnmarshalFromRepository(
		user.Id, user.Name, user.RegistrationDate, user.LastLogin,
		user.Karma, user.PostsCount, user.Role)
}
