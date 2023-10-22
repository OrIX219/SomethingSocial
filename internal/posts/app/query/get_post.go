package query

import (
	"context"

	posts "github.com/OrIX219/SomethingSocial/internal/posts/domain/post"
)

type GetPost struct {
	PostId string
}

type GetPostHandler struct {
	repo posts.Repository
}

func NewGetPostHandler(repo posts.Repository) GetPostHandler {
	if repo == nil {
		panic("GetPostHandler nil repo")
	}

	return GetPostHandler{
		repo: repo,
	}
}

func (h GetPostHandler) Handle(ctx context.Context,
	query GetPost) (*posts.Post, error) {
	return h.repo.GetPost(query.PostId)
}
