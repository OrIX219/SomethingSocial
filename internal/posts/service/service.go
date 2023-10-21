package service

import (
	"os"

	"github.com/OrIX219/SomethingSocial/internal/common/client"
	"github.com/OrIX219/SomethingSocial/internal/posts/adapters"
	"github.com/OrIX219/SomethingSocial/internal/posts/app"
	"github.com/OrIX219/SomethingSocial/internal/posts/app/command"
	"github.com/OrIX219/SomethingSocial/internal/posts/app/query"
	posts "github.com/OrIX219/SomethingSocial/internal/posts/domain/post"
)

func NewApplication() (app.Application, func()) {
	db, err := client.NewPostgres(os.Getenv("POSTS_POSTGRES_ADDR"))
	if err != nil {
		panic(err)
	}
	repo := adapters.NewPostsPostgresRepository(db)

	usersClient, closeUsersClient, err := client.NewUsersClient()
	if err != nil {
		panic(err)
	}
	usersGrpc := adapters.NewUsersGrpc(usersClient)

	return newApplication(repo, usersGrpc), func() {
		_ = closeUsersClient()
	}
}

func newApplication(repo posts.Repository,
	usersService app.UsersService) app.Application {
	return app.Application{
		Commands: app.Commands{
			CreatePost:     command.NewCreatePostHandler(repo),
			DeletePost:     command.NewDeletePostHandler(repo),
			UpdatePost:     command.NewUpdatePostHandler(repo),
			UpvotePost:     command.NewUpvotePostHandler(repo, usersService),
			RemoveUpvote:   command.NewRemoveUpvoteHandler(repo, usersService),
			DownvotePost:   command.NewDownvotePostHandler(repo, usersService),
			RemoveDownvote: command.NewRemoveDownvoteHandler(repo, usersService),
		},
		Queries: app.Queries{
			GetPost:       query.NewGetPostHandler(repo),
			GetPosts:      query.NewGetPostsHandler(repo),
			GetPostsCount: query.NewGetPostsCountHandler(repo),
			GetFeed:       query.NewGetFeedHandler(repo, usersService),
		},
	}
}
