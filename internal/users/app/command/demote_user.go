package command

import (
	"context"

	users "github.com/OrIX219/SomethingSocial/internal/users/domain/user"
)

type DemoteUser struct {
	UserId       int64
	TargetUserId int64
}

type DemoteUserHandler struct {
	repo users.Repository
}

func NewDemoteUserHandler(repo users.Repository) DemoteUserHandler {
	if repo == nil {
		panic("DemoteUserHandler nil repo")
	}

	return DemoteUserHandler{
		repo: repo,
	}
}

func (h DemoteUserHandler) Handle(ctx context.Context, cmd DemoteUser) error {
	initiator, err := h.repo.GetUser(cmd.UserId)
	if err != nil {
		return err
	}

	return h.repo.UpdateUser(cmd.TargetUserId,
		func(user *users.User) (*users.User, error) {
			err := user.Demote(initiator)
			if err != nil {
				return nil, err
			}

			return user, nil
		})
}
