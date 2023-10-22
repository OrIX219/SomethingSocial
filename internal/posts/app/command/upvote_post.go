package command

import (
	"context"

	posts "github.com/OrIX219/SomethingSocial/internal/posts/domain/post"
)

type UpvotePost struct {
	PostId string
	UserId int64
}

type UpvotePostHandler struct {
	repo         posts.Repository
	usersService UsersService
}

func NewUpvotePostHandler(repo posts.Repository,
	usersService UsersService) UpvotePostHandler {
	if repo == nil {
		panic("UpvotePostHandler nil repo")
	}
	if usersService == nil {
		panic("UpvotePostHandler nil usersService")
	}

	return UpvotePostHandler{
		repo:         repo,
		usersService: usersService,
	}
}

func (h UpvotePostHandler) Handle(ctx context.Context, cmd UpvotePost) error {
	karmaDelta, err := h.repo.UpvotePost(cmd.PostId, cmd.UserId)
	if err != nil {
		return err
	}

	author, err := h.repo.GetAuthor(cmd.PostId)
	if err != nil {
		return err
	}

	return h.usersService.UpdateKarma(ctx, author, int64(karmaDelta))
}
