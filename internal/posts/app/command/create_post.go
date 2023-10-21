package command

import (
	"context"
	"time"

	posts "github.com/OrIX219/SomethingSocial/internal/posts/domain/post"
)

type CreatePost struct {
	PostId   string
	Content  string
	PostDate time.Time
	Author   int64
}

type CreatePostHandler struct {
	repo posts.Repository
}

func NewCreatePostHandler(repo posts.Repository) CreatePostHandler {
	if repo == nil {
		panic("CreatePostHandler nil repo")
	}

	return CreatePostHandler{
		repo: repo,
	}
}

func (h CreatePostHandler) Handle(ctx context.Context, cmd CreatePost) error {
	post, err := posts.NewPost(cmd.PostId, cmd.Content, cmd.PostDate, cmd.Author)
	if err != nil {
		return err
	}

	return h.repo.AddPost(post)
}
