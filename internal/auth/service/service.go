package service

import (
	"github.com/OrIX219/SomethingSocial/internal/auth/adapters"
	"github.com/OrIX219/SomethingSocial/internal/auth/app"
	"github.com/OrIX219/SomethingSocial/internal/auth/app/command"
	"github.com/OrIX219/SomethingSocial/internal/auth/app/query"
)

func NewApplication() app.Application {
	return newApplication()
}

func newApplication() app.Application {
	repo := adapters.NewUsersInMemoryRepository(10)

	return app.Application{
		Commands: app.Commands{
			AddUser: command.NewAddUserHandler(repo),
		},
		Queries: app.Queries{
			GetUserId:     query.NewGetUserIdHandler(repo),
			GenerateToken: query.NewGenerateTokenHandler(),
		},
	}
}
