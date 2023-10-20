package query

import (
	"context"

	posts "github.com/OrIX219/SomethingSocial/internal/posts/domain/post"
)

type GetFeed struct {
	UserId int64
}

type GetFeedHandler struct {
	repo         posts.Repository
	usersService UsersService
}

func NewGetFeedHandler(repo posts.Repository, usersService UsersService) GetFeedHandler {
	if repo == nil {
		panic("GetFeedHandler nil repo")
	}
	if usersService == nil {
		panic("GetFeedHandler nil usersService")
	}

	return GetFeedHandler{
		repo:         repo,
		usersService: usersService,
	}
}

func (h GetFeedHandler) Handle(ctx context.Context,
	query GetFeed) ([]*posts.Post, error) {
	following, err := h.usersService.GetFollowing(ctx, query.UserId)
	if err != nil {
		return nil, err
	}

	return h.repo.GetFeed(following)
}
