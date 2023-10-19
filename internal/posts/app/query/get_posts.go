package query

import (
	"context"

	posts "github.com/OrIX219/SomethingSocial/internal/posts/domain/post"
)

type GetPosts struct {
	Filter posts.PostFilter
}

type GetPostsHandler struct {
	repo posts.Repository
}

func NewGetPostsHandler(repo posts.Repository) GetPostsHandler {
	if repo == nil {
		panic("GetPostsHandler nil repo")
	}

	return GetPostsHandler{
		repo: repo,
	}
}

func (h GetPostsHandler) Handle(ctx context.Context,
	query GetPosts) ([]*posts.Post, error) {
	return h.repo.GetPosts(query.Filter)
}
