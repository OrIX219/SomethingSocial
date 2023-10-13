package app

import (
	"github.com/OrIX219/SomethingSocial/internal/auth/app/command"
	"github.com/OrIX219/SomethingSocial/internal/auth/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	AddUser command.AddUserHandler
}

type Queries struct {
	GetUserId     query.GetUserIdHadler
	GenerateToken query.GenerateTokenHandler
}
