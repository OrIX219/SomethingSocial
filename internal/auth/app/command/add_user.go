package command

import (
	"context"

	auth "github.com/OrIX219/SomethingSocial/internal/auth/domain/user"
)

type AddUser struct {
	Name     string
	Username string
	Password string
}

type AddUserHandler struct {
	repo         auth.Repository
	usersService UsersService
}

func NewAddUserHandler(repo auth.Repository, usersService UsersService) AddUserHandler {
	if repo == nil {
		panic("AddUserHandler nil repo")
	}
	if usersService == nil {
		panic("AddUserHandler nil usersService")
	}

	return AddUserHandler{
		repo:         repo,
		usersService: usersService,
	}
}

func (h AddUserHandler) Handle(ctx context.Context, cmd AddUser) error {
	user, err := auth.NewUser(cmd.Username, cmd.Password)
	if err != nil {
		return err
	}

	if cmd.Name == "" {
		cmd.Name = "Unnamed"
	}

	if usr, _ := h.repo.GetUserByUsername(cmd.Username); usr != nil {
		return UsernameExistsError{
			Username: cmd.Username,
		}
	}

	id, err := h.repo.AddUser(user)
	if err != nil {
		return err
	}

	return h.usersService.AddUser(ctx, id, cmd.Name)
}

type UsernameExistsError struct {
	Username string
}

func (e UsernameExistsError) Error() string {
	return e.Username
}
