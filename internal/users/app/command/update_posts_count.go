package command

import (
	"context"

	users "github.com/OrIX219/SomethingSocial/internal/users/domain/user"
)

type UpdatePostsCount struct {
	UserId int64
	Delta  int64
}

type UpdatePostsCountHandler struct {
	repo users.Repository
}

func NewUpdatePostsCountHandler(repo users.Repository) UpdatePostsCountHandler {
	if repo == nil {
		panic("UpdatePostsCountHandler nil repo")
	}

	return UpdatePostsCountHandler{
		repo: repo,
	}
}

func (h UpdatePostsCountHandler) Handle(ctx context.Context, cmd UpdatePostsCount) error {
	return h.repo.UpdateUser(cmd.UserId,
		func(user *users.User) (*users.User, error) {
			user.UpdatePostsCount(cmd.Delta)
			return user, nil
		})
}
