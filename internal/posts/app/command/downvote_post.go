package command

import (
	"context"

	posts "github.com/OrIX219/SomethingSocial/internal/posts/domain/post"
)

type DownvotePost struct {
	PostId string
	UserId int64
}

type DownvotePostHandler struct {
	repo         posts.Repository
	usersService UsersService
}

func NewDownvotePostHandler(repo posts.Repository,
	usersService UsersService) DownvotePostHandler {
	if repo == nil {
		panic("DownvotePostHandler nil repo")
	}
	if usersService == nil {
		panic("UpvotePostHandler nil usersService")
	}

	return DownvotePostHandler{
		repo:         repo,
		usersService: usersService,
	}
}

func (h DownvotePostHandler) Handle(ctx context.Context, cmd DownvotePost) error {
	karmaDelta, err := h.repo.DownvotePost(cmd.PostId, cmd.UserId)
	if err != nil {
		return err
	}

	author, err := h.repo.GetAuthor(cmd.PostId)
	if err != nil {
		return err
	}

	return h.usersService.UpdateKarma(ctx, author, int64(karmaDelta))
}
