package command

import (
	"context"
	"fmt"
	"sort"

	posts "github.com/OrIX219/SomethingSocial/internal/posts/domain/post"
	"golang.org/x/exp/slices"
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
	downvoters, err := h.repo.GetDownvoters(cmd.PostId)
	if err != nil {
		return err
	}

	sort.Slice(downvoters, func(i, j int) bool {
		return downvoters[i] < downvoters[j]
	})
	_, found := slices.BinarySearch[int64](downvoters, cmd.UserId)
	if found {
		return AlreadyDownvotedError{
			PostId: cmd.PostId,
		}
	}

	err = h.repo.UpdatePost(cmd.PostId, func(post *posts.Post) (*posts.Post, error) {
		post.Downvote()
		return post, nil
	})
	if err != nil {
		return err
	}

	author, err := h.repo.GetAuthor(cmd.PostId)
	if err != nil {
		return err
	}
	return h.usersService.UpdateKarma(ctx, author, -1)
}

type AlreadyDownvotedError struct {
	PostId string
}

func (e AlreadyDownvotedError) Error() string {
	return fmt.Sprintf("Post %s already downvoted", e.PostId)
}
