package query

import (
	"context"

	posts "github.com/OrIX219/SomethingSocial/internal/posts/domain/post"
)

type GetFeed struct {
	UserId int64
}

type GetFeedHandler struct {
	repo posts.Repository
}

func NewGetFeedHandler(repo posts.Repository) GetFeedHandler {
	if repo == nil {
		panic("GetFeedHandler nil repo")
	}

	return GetFeedHandler{
		repo: repo,
	}
}

func (h GetFeedHandler) Handle(ctx context.Context,
	query GetFeed) ([]*posts.Post, error) {
	return h.repo.GetFeed(query.UserId)
}
