package command

import (
	"context"
	"time"

	posts "github.com/OrIX219/SomethingSocial/internal/posts/domain/post"
)

type CreatePost struct {
	PostId   string
	Content  string
	PostDate time.Time
	Author   int64
}

type CreatePostHandler struct {
	repo         posts.Repository
	usersService UsersService
}

func NewCreatePostHandler(repo posts.Repository,
	usersService UsersService) CreatePostHandler {
	if repo == nil {
		panic("CreatePostHandler nil repo")
	}
	if usersService == nil {
		panic("CreatePostHandler nil usersService")
	}

	return CreatePostHandler{
		repo:         repo,
		usersService: usersService,
	}
}

func (h CreatePostHandler) Handle(ctx context.Context, cmd CreatePost) error {
	post, err := posts.NewPost(cmd.PostId, cmd.Content, cmd.PostDate, cmd.Author)
	if err != nil {
		return err
	}

	err = h.repo.AddPost(post)
	if err != nil {
		return err
	}

	return h.usersService.UpdatePostsCount(ctx, cmd.Author, 1)
}
