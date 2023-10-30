package command

import (
	"context"

	users "github.com/OrIX219/SomethingSocial/internal/users/domain/user"
)

type PromoteUser struct {
	UserId       int64
	TargetUserId int64
}

type PromoteUserHandler struct {
	repo users.Repository
}

func NewPromoteUserHandler(repo users.Repository) PromoteUserHandler {
	if repo == nil {
		panic("PromoteUserHandler nil repo")
	}

	return PromoteUserHandler{
		repo: repo,
	}
}

func (h PromoteUserHandler) Handle(ctx context.Context, cmd PromoteUser) error {
	initiator, err := h.repo.GetUser(cmd.UserId)
	if err != nil {
		return err
	}

	return h.repo.UpdateUser(cmd.TargetUserId,
		func(user *users.User) (*users.User, error) {
			err := user.Promote(initiator)
			if err != nil {
				return nil, err
			}

			return user, nil
		})
}
