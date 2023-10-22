package query

import (
	"context"

	users "github.com/OrIX219/SomethingSocial/internal/users/domain/user"
)

type GetUser struct {
	UserId int64
}

type GetUserHandler struct {
	repo users.Repository
}

func NewGetUserHandler(repo users.Repository) GetUserHandler {
	if repo == nil {
		panic("GetUserHandler nil repo")
	}

	return GetUserHandler{
		repo: repo,
	}
}

func (h GetUserHandler) Handle(ctx context.Context,
	query GetUser) (*users.User, error) {
	return h.repo.GetUser(query.UserId)
}
