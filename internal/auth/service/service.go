package service

import (
	"os"

	"github.com/OrIX219/SomethingSocial/internal/auth/adapters"
	"github.com/OrIX219/SomethingSocial/internal/auth/app"
	"github.com/OrIX219/SomethingSocial/internal/auth/app/command"
	"github.com/OrIX219/SomethingSocial/internal/auth/app/query"
	auth "github.com/OrIX219/SomethingSocial/internal/auth/domain/user"
	"github.com/OrIX219/SomethingSocial/internal/common/client"
)

func NewApplication() (app.Application, func()) {
	db, err := client.NewPostgres(os.Getenv("AUTH_POSTGRES_ADDR"))
	if err != nil {
		panic(err)
	}
	repo := adapters.NewUsersPostgresRepository(db)

	usersClient, closeUsersClient, err := client.NewUsersClient()
	if err != nil {
		panic(err)
	}
	usersGrpc := adapters.NewUsersGrpc(usersClient)

	return newApplication(repo, usersGrpc), func() {
		_ = closeUsersClient()
	}
}

func newApplication(repo auth.Repository, usersService command.UsersService) app.Application {
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
