package command

import (
	"context"
	"time"

	users "github.com/OrIX219/SomethingSocial/internal/users/domain/user"
)

type UpdateLastLogIn struct {
	UserId int64
	Time   time.Time
}

type UpdateLastLogInHandler struct {
	repo users.Repository
}

func NewUpdateLastLogInHandler(repo users.Repository) UpdateLastLogInHandler {
	if repo == nil {
		panic("UpdateLastLogInHandler nil repo")
	}

	return UpdateLastLogInHandler{
		repo: repo,
	}
}

func (h UpdateLastLogInHandler) Handle(ctx context.Context,
	cmd UpdateLastLogIn) error {
	return h.repo.UpdateUser(cmd.UserId,
		func(user *users.User) (*users.User, error) {
			user.LogInAt(cmd.Time)
			return user, nil
		})
}
