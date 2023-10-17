package adapters

import (
	"database/sql"
	"fmt"

	auth "github.com/OrIX219/SomethingSocial/internal/auth/domain/user"
	"github.com/jmoiron/sqlx"
)

const (
	usersTable = "users"
)

type UserModel struct {
	Id           int64  `db:"id"`
	Username     string `db:"username"`
	PasswordHash string `db:"password_hash"`
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

func (r *UsersPostgresRepository) AddUser(user *auth.User) (int64, error) {
	var id int64
	query := fmt.Sprintf(`INSERT INTO %s (username, password_hash)
		VALUES ($1, $2) RETURNING id`, usersTable)
	row := r.db.QueryRow(query, user.Username(), user.Password())
	if err := row.Scan(&id); err != nil {
		return -1, err
	}

	return id, nil
}

func (r *UsersPostgresRepository) GetUserId(user *auth.User) (int64, error) {
	var id int64
	query := fmt.Sprintf(`SELECT id FROM %s
		WHERE username=$1 AND password_hash=$2`, usersTable)
	if err := r.db.Get(&id, query, user.Username(), user.Password()); err != nil {
		return -1, auth.UserNotFoundError{
			Username: user.Username(),
			Password: user.Password(),
		}
	}

	return id, nil
}

func (r *UsersPostgresRepository) GetUserByUsername(username string) (*auth.User, error) {
	var userModel UserModel
	query := fmt.Sprintf(`SELECT * FROM %s WHERE username=$1`, usersTable)
	err := r.db.Get(&userModel, query, username)
	if err == sql.ErrNoRows {
		return nil, auth.UserNotFoundError{
			Username: username,
		}
	}

	return r.unmarshalUser(userModel)
}

func (r *UsersPostgresRepository) marshalUser(user *auth.User) UserModel {
	userModel := UserModel{
		Username:     user.Username(),
		PasswordHash: user.Password(),
	}

	return userModel
}

func (r *UsersPostgresRepository) unmarshalUser(user UserModel) (*auth.User, error) {
	return auth.UnmarshalFromRepository(user.Id, user.Username, user.PasswordHash)
}
