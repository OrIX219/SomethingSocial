package app

import (
	"github.com/OrIX219/SomethingSocial/internal/posts/app/command"
	"github.com/OrIX219/SomethingSocial/internal/posts/app/query"
)

type UsersService interface {
	command.UsersService
	query.UsersService
}
