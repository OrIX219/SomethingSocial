package command

import (
	"context"

	users "github.com/OrIX219/SomethingSocial/internal/users/domain/user"
)

type UnfollowUser struct {
	UserId   int64
	TargetId int64
}

type UnfollowUserHandler struct {
	repo users.Repository
}

func NewUnfollowUserHandler(repo users.Repository) UnfollowUserHandler {
	if repo == nil {
		panic("UnfollowUserHandler nil repo")
	}

	return UnfollowUserHandler{
		repo: repo,
	}
}

func (h UnfollowUserHandler) Handle(ctx context.Context, cmd UnfollowUser) error {
	_, err := h.repo.GetUser(cmd.TargetId)
	if err != nil {
		return err
	}

	return h.repo.UnfollowUser(cmd.UserId, cmd.TargetId)
}
