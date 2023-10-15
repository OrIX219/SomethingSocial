package query

import users "github.com/OrIX219/SomethingSocial/internal/users/domain/user"

type GetKarma struct {
	UserId int64
}

type GetKarmaHandler struct {
	repo users.Repository
}

func NewGetKarmaHandler(repo users.Repository) GetKarmaHandler {
	if repo == nil {
		panic("GetKarmaHandler nil repo")
	}

	return GetKarmaHandler{
		repo: repo,
	}
}

func (h GetKarmaHandler) Handle(query GetKarma) (int64, error) {
	user, err := h.repo.GetUserById(query.UserId)
	if err != nil {
		return 0, err
	}

	return user.Karma(), nil
}
