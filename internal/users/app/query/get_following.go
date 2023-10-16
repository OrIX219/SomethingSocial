package query

import (
	"context"

	users "github.com/OrIX219/SomethingSocial/internal/users/domain/user"
)

type GetFollowing struct {
	UserId int64
}

type GetFollowingHandler struct {
	repo users.Repository
}

func NewGetFollowingHandler(repo users.Repository) GetFollowingHandler {
	if repo == nil {
		panic("GetFollowingHandler nil repo")
	}

	return GetFollowingHandler{
		repo: repo,
	}
}

func (h GetFollowingHandler) Handle(ctx context.Context,
	query GetFollowing) ([]*users.User, error) {
	return h.repo.GetFollowing(query.UserId)
}
