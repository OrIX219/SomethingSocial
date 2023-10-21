package command

import (
	"context"
	"time"

	posts "github.com/OrIX219/SomethingSocial/internal/posts/domain/post"
)

type EditPost struct {
	PostId   string
	Content  string
	EditDate time.Time
	Author   int64
}

type EditPostHandler struct {
	repo posts.Repository
}

func NewEditPostHandler(repo posts.Repository) EditPostHandler {
	if repo == nil {
		panic("EditPostHandler nil repo")
	}

	return EditPostHandler{
		repo: repo,
	}
}

func (h EditPostHandler) Handle(ctx context.Context, cmd EditPost) error {
	post, err := posts.NewPost(cmd.PostId, cmd.Content, cmd.EditDate, cmd.Author)
	if err != nil {
		return err
	}

	return h.repo.EditPost(cmd.Author, post)
}
