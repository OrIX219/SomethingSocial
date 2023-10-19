package command

import (
	"context"
	"fmt"
	"sort"

	posts "github.com/OrIX219/SomethingSocial/internal/posts/domain/post"
	"golang.org/x/exp/slices"
	"google.golang.org/appengine/user"
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
	upvoters, err := h.repo.GetUpvoters(cmd.PostId)
	if err != nil {
		return err
	}

	sort.Slice(upvoters, func(i, j int) bool {
		return upvoters[i] < upvoters[j]
	})
	_, found := slices.BinarySearch[int64](upvoters, cmd.UserId)
	if found {
		return AlreadyUpvotedError{
			PostId: cmd.PostId,
		}
	}

	err = h.repo.UpdatePost(cmd.PostId, func(post *posts.Post) (*posts.Post, error) {
		post.Upvote()
		return post, nil
	})
	if err != nil {
		return err
	}

	author, err := h.repo.GetAuthor(cmd.PostId)
	if err != nil {
		return err
	}

	return h.usersService.UpdateKarma(ctx, author, 1)
}

type AlreadyUpvotedError struct {
	PostId string
}

func (e AlreadyUpvotedError) Error() string {
	return fmt.Sprintf("Post %s already upvoted", e.PostId)
}
