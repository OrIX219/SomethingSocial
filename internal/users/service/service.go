package service

import (
	"github.com/OrIX219/SomethingSocial/internal/users/adapters"
	"github.com/OrIX219/SomethingSocial/internal/users/app"
	"github.com/OrIX219/SomethingSocial/internal/users/app/command"
	"github.com/OrIX219/SomethingSocial/internal/users/app/query"
)

func NewApplication() app.Application {
	return newApplication()
}

func newApplication() app.Application {
	repo := adapters.NewUsersInMemoryRepository(10)

	return app.Application{
		Commands: app.Commands{
			AddUser:     command.NewAddUserHandler(repo),
			UpdateKarma: command.NewUpdateKarmaHandler(repo),
		},
		Queries: app.Queries{
			GetKarma: query.NewGetKarmaHandler(repo),
		},
	}
}
