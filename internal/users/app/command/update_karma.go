package command

import users "github.com/OrIX219/SomethingSocial/internal/users/domain/user"

type UpdateKarma struct {
	UserId int64
	Delta  int64
}

type UpdateKarmaHandler struct {
	repo users.Repository
}

func NewUpdateKarmaHandler(repo users.Repository) UpdateKarmaHandler {
	if repo == nil {
		panic("UpdateKarmaHandler nil repo")
	}

	return UpdateKarmaHandler{
		repo: repo,
	}
}

func (h UpdateKarmaHandler) Handle(cmd UpdateKarma) error {
	return h.repo.UpdateUser(cmd.UserId,
		func(user *users.User) (*users.User, error) {
			user.UpdateKarma(cmd.Delta)
			return user, nil
		})
}
