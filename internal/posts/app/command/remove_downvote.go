package command

import (
	"context"

	posts "github.com/OrIX219/SomethingSocial/internal/posts/domain/post"
)

type RemoveDownvote struct {
	PostId string
	UserId int64
}

type RemoveDownvoteHandler struct {
	repo         posts.Repository
	usersService UsersService
}

func NewRemoveDownvoteHandler(repo posts.Repository,
	usersService UsersService) RemoveDownvoteHandler {
	if repo == nil {
		panic("RemoveDownvoteHandler nil repo")
	}
	if usersService == nil {
		panic("RemoveDownvoteHandler nil usersService")
	}

	return RemoveDownvoteHandler{
		repo:         repo,
		usersService: usersService,
	}
}

func (h RemoveDownvoteHandler) Handle(ctx context.Context, cmd RemoveDownvote) error {
	karmaDelta, err := h.repo.RemoveDownvote(cmd.PostId, cmd.UserId)
	if err != nil {
		return err
	}

	author, err := h.repo.GetAuthor(cmd.PostId)
	if err != nil {
		return err
	}

	return h.usersService.UpdateKarma(ctx, author, int64(karmaDelta))
}
