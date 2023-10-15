package command

import users "github.com/OrIX219/SomethingSocial/internal/users/domain/user"

type AddUser struct {
	UserId int64
	Name   string
}

type AddUserHandler struct {
	repo users.Repository
}

func NewAddUserHandler(repo users.Repository) AddUserHandler {
	if repo == nil {
		panic("AddUserHandler nil repo")
	}

	return AddUserHandler{
		repo: repo,
	}
}

func (h AddUserHandler) Handle(cmd AddUser) error {
	user, err := users.NewUser(cmd.UserId, cmd.Name)
	if err != nil {
		return err
	}

	return h.repo.AddUser(user)
}
