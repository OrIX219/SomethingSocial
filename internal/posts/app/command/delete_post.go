package command

import (
	"context"

	posts "github.com/OrIX219/SomethingSocial/internal/posts/domain/post"
)

type DeletePost struct {
	PostId string
	UserId int64
}

type DeletePostHandler struct {
	repo         posts.Repository
	usersService UsersService
}

func NewDeletePostHandler(repo posts.Repository,
	usersService UsersService) DeletePostHandler {
	if repo == nil {
		panic("DeletePostHandler nil repo")
	}
	if usersService == nil {
		panic("DeletePostHandler nil usersService")
	}

	return DeletePostHandler{
		repo:         repo,
		usersService: usersService,
	}
}

func (h DeletePostHandler) Handle(ctx context.Context, cmd DeletePost) error {
	err := h.repo.DeletePost(cmd.PostId, cmd.UserId)
	if err != nil {
		return err
	}

	return h.usersService.UpdatePostsCount(ctx, cmd.UserId, -1)
}
