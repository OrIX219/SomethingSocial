package command

import (
	"context"
	"time"

	posts "github.com/OrIX219/SomethingSocial/internal/posts/domain/post"
)

type UpdatePost struct {
	PostId     string
	Content    string
	UpdateDate time.Time
	Author     int64
}

type UpdatePostHandler struct {
	repo posts.Repository
}

func NewUpdatePostHandler(repo posts.Repository) UpdatePostHandler {
	if repo == nil {
		panic("UpdatePostHandler nil repo")
	}

	return UpdatePostHandler{
		repo: repo,
	}
}

func (h UpdatePostHandler) Handle(ctx context.Context, cmd UpdatePost) error {
	post, err := posts.NewPost(cmd.PostId, cmd.Content, cmd.UpdateDate, cmd.Author)
	if err != nil {
		return err
	}

	return h.repo.UpdatePost(cmd.Author, post)
}
