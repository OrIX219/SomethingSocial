package app

import (
	"github.com/OrIX219/SomethingSocial/internal/users/app/command"
	"github.com/OrIX219/SomethingSocial/internal/users/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	AddUser          command.AddUserHandler
	UpdateKarma      command.UpdateKarmaHandler
	UpdatePostsCount command.UpdatePostsCountHandler
	UpdateLastLogIn  command.UpdateLastLogInHandler
	FollowUser       command.FollowUserHandler
	UnfollowUser     command.UnfollowUserHandler
	PromoteUser      command.PromoteUserHandler
	DemoteUser       command.DemoteUserHandler
}

type Queries struct {
	GetKarma     query.GetKarmaHandler
	GetUser      query.GetUserHandler
	GetFollowing query.GetFollowingHandler
	GetFollowers query.GetFollowersHandler
}
