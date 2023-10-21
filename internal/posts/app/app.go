package app

import (
	"github.com/OrIX219/SomethingSocial/internal/posts/app/command"
	"github.com/OrIX219/SomethingSocial/internal/posts/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreatePost     command.CreatePostHandler
	DeletePost     command.DeletePostHandler
	UpvotePost     command.UpvotePostHandler
	RemoveUpvote   command.RemoveUpvoteHandler
	DownvotePost   command.DownvotePostHandler
	RemoveDownvote command.RemoveDownvoteHandler
}

type Queries struct {
	GetPost       query.GetPostHandler
	GetPosts      query.GetPostsHandler
	GetPostsCount query.GetPostsCountHandler
	GetFeed       query.GetFeedHandler
}
