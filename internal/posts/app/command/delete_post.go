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
	repo posts.Repository
}

func NewDeletePostHandler(repo posts.Repository) DeletePostHandler {
	if repo == nil {
		panic("DeletePostHandler nil repo")
	}

	return DeletePostHandler{
		repo: repo,
	}
}

func (h DeletePostHandler) Handle(ctx context.Context, cmd DeletePost) error {
	return h.repo.DeletePost(cmd.PostId, cmd.UserId)
}
