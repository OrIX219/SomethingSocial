package service

import (
	"os"

	"github.com/OrIX219/SomethingSocial/internal/common/client"
	"github.com/OrIX219/SomethingSocial/internal/users/adapters"
	"github.com/OrIX219/SomethingSocial/internal/users/app"
	"github.com/OrIX219/SomethingSocial/internal/users/app/command"
	"github.com/OrIX219/SomethingSocial/internal/users/app/query"
	users "github.com/OrIX219/SomethingSocial/internal/users/domain/user"
)

func NewApplication() app.Application {
	db, err := client.NewPostgres(os.Getenv("USERS_POSTGRES_ADDR"))
	if err != nil {
		panic(err)
	}
	repo := adapters.NewUsersPostgresRepository(db)

	return newApplication(repo)
}

func newApplication(repo users.Repository) app.Application {
	return app.Application{
		Commands: app.Commands{
			AddUser:      command.NewAddUserHandler(repo),
			UpdateKarma:  command.NewUpdateKarmaHandler(repo),
			FollowUser:   command.NewFollowUserHandler(repo),
			UnfollowUser: command.NewUnfollowUserHandler(repo),
		},
		Queries: app.Queries{
			GetKarma:     query.NewGetKarmaHandler(repo),
			GetUser:      query.NewGetUserHandler(repo),
			GetFollowing: query.NewGetFollowingHandler(repo),
			GetFollowers: query.NewGetFollowersHandler(repo),
		},
	}
}
