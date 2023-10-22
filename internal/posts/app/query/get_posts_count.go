package query

import (
	"context"

	posts "github.com/OrIX219/SomethingSocial/internal/posts/domain/post"
)

type GetPostsCount struct {
	UserId int64
}

type GetPostsCountHandler struct {
	repo posts.Repository
}

func NewGetPostsCountHandler(repo posts.Repository) GetPostsCountHandler {
	if repo == nil {
		panic("GetPostsCountHandler nil repo")
	}

	return GetPostsCountHandler{
		repo: repo,
	}
}

func (h GetPostsCountHandler) Handle(ctx context.Context,
	query GetPostsCount) (int64, error) {
	return h.repo.GetPostsCount(query.UserId)
}
