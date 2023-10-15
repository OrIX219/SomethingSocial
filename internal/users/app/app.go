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
	AddUser     command.AddUserHandler
	UpdateKarma command.UpdateKarmaHandler
}

type Queries struct {
	GetKarma query.GetKarmaHandler
}
