package command

import (
	"context"

	users "github.com/OrIX219/SomethingSocial/internal/users/domain/user"
)

type FollowUser struct {
	UserId   int64
	TargetId int64
}

type FollowUserHandler struct {
	repo users.Repository
}

func NewFollowUserHandler(repo users.Repository) FollowUserHandler {
	if repo == nil {
		panic("FollowUserHandler nil repo")
	}

	return FollowUserHandler{
		repo: repo,
	}
}

func (h FollowUserHandler) Handle(ctx context.Context, cmd FollowUser) error {
	_, err := h.repo.GetUser(cmd.TargetId)
	if err != nil {
		return err
	}

	return h.repo.FollowUser(cmd.UserId, cmd.TargetId)
}
