package service

import (
	"github.com/OrIX219/SomethingSocial/internal/posts/app"
	"github.com/OrIX219/SomethingSocial/internal/posts/app/command"
	"github.com/OrIX219/SomethingSocial/internal/posts/app/query"
)

func NewApplication() app.Application {
	// db, err := client.NewPostgres(os.Getenv("POSTS_POSTGRES_ADDR"))
	// if err != nil {
	// 	panic(err)
	// }
	// repo := adapters.NewUsersPostgresRepository(db)

	return newApplication()
}

func newApplication() app.Application {
	return app.Application{
		Commands: app.Commands{
			CreatePost:   command.NewCreatePostHandler(repo),
			DeletePost:   command.NewDeletePostHandler(repo),
			UpvotePost:   command.NewUpvotePostHandler(repo),
			DownvotePost: command.NewDownvotePostHandler(repo),
		},
		Queries: app.Queries{
			GetPost:       query.NewGetPostHandler(repo),
			GetPosts:      query.NewGetPostsHandler(repo),
			GetPostsCount: query.NewGetPostsCountHandler(repo),
			GetFeed:       query.NewGetFeedHandler(repo),
		},
	}
}
