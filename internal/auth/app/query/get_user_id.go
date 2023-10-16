package query

import (
	"context"

	auth "github.com/OrIX219/SomethingSocial/internal/auth/domain/user"
)

type GetUserId struct {
	Username string
	Password string
}

type GetUserIdHadler struct {
	repo auth.Repository
}

func NewGetUserIdHandler(repo auth.Repository) GetUserIdHadler {
	if repo == nil {
		panic("GetUserHandler nil repo")
	}

	return GetUserIdHadler{
		repo: repo,
	}
}

func (h GetUserIdHadler) Handle(ctx context.Context, query GetUserId) (int64, error) {
	user, err := auth.NewUser(query.Username, query.Password)
	if err != nil {
		return -1, err
	}

	return h.repo.GetUserId(user)
}
