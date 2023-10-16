package query

import (
	"context"

	users "github.com/OrIX219/SomethingSocial/internal/users/domain/user"
)

type GetFollowers struct {
	UserId int64
}

type GetFollowersHandler struct {
	repo users.Repository
}

func NewGetFollowersHandler(repo users.Repository) GetFollowersHandler {
	if repo == nil {
		panic("GetFollowersHandler nil repo")
	}

	return GetFollowersHandler{
		repo: repo,
	}
}

func (h GetFollowersHandler) Handle(ctx context.Context,
	query GetFollowers) ([]*users.User, error) {
	return h.repo.GetFollowers(query.UserId)
}
