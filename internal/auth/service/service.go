package service

import (
	"github.com/OrIX219/SomethingSocial/internal/auth/adapters"
	"github.com/OrIX219/SomethingSocial/internal/auth/app"
	"github.com/OrIX219/SomethingSocial/internal/auth/app/command"
	"github.com/OrIX219/SomethingSocial/internal/auth/app/query"
	"github.com/OrIX219/SomethingSocial/internal/common/client"
)

func NewApplication() (app.Application, func()) {
	usersClient, closeUsersClient, err := client.NewUsersClient()
	if err != nil {
		panic(err)
	}

	usersGrpc := adapters.NewUsersGrpc(usersClient)

	return newApplication(usersGrpc), func() {
		_ = closeUsersClient()
	}
}

func newApplication(usersService command.UsersService) app.Application {
	repo := adapters.NewUsersInMemoryRepository(10)

	return app.Application{
		Commands: app.Commands{
			AddUser: command.NewAddUserHandler(repo, usersService),
		},
		Queries: app.Queries{
			GetUserId:     query.NewGetUserIdHandler(repo),
			GenerateToken: query.NewGenerateTokenHandler(),
		},
	}
}
