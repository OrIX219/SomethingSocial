package command

import (
	"context"

	posts "github.com/OrIX219/SomethingSocial/internal/posts/domain/post"
)

type RemoveUpvote struct {
	PostId string
	UserId int64
}

type RemoveUpvoteHandler struct {
	repo         posts.Repository
	usersService UsersService
}

func NewRemoveUpvoteHandler(repo posts.Repository,
	usersService UsersService) RemoveUpvoteHandler {
	if repo == nil {
		panic("RemoveUpvoteHandler nil repo")
	}
	if usersService == nil {
		panic("RemoveUpvoteHandler nil usersService")
	}

	return RemoveUpvoteHandler{
		repo:         repo,
		usersService: usersService,
	}
}

func (h RemoveUpvoteHandler) Handle(ctx context.Context, cmd RemoveUpvote) error {
	karmaDelta, err := h.repo.RemoveUpvote(cmd.PostId, cmd.UserId)
	if err != nil {
		return err
	}

	author, err := h.repo.GetAuthor(cmd.PostId)
	if err != nil {
		return err
	}

	return h.usersService.UpdateKarma(ctx, author, int64(karmaDelta))
}
