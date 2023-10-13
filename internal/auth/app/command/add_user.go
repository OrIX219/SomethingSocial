package command

import auth "github.com/OrIX219/SomethingSocial/internal/auth/domain/user"

type AddUser struct {
	Username string
	Password string
}

type AddUserHandler struct {
	repo auth.Repository
}

func NewAddUserHandler(repo auth.Repository) AddUserHandler {
	if repo == nil {
		panic("AddUserHandler nil repo")
	}

	return AddUserHandler{
		repo: repo,
	}
}

func (h AddUserHandler) Handle(cmd AddUser) error {
	user, err := auth.NewUser(cmd.Username, cmd.Password)
	if err != nil {
		return err
	}

	if usr, _ := h.repo.GetUserByUsername(cmd.Username); usr != nil {
		return UsernameExistsError{
			Username: cmd.Username,
		}
	}

	_, err = h.repo.AddUser(user)
	// TODO: gRPC call to Users service to add user
	return err
}

type UsernameExistsError struct {
	Username string
}

func (e UsernameExistsError) Error() string {
	return e.Username
}
