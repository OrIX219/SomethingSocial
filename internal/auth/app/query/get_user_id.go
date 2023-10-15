package query

import (
	"context"

	auth "github.com/OrIX219/SomethingSocial/internal/auth/domain/user"
)

type GetUserId struct {
	Username string
	Password string
}

type GetUserIdHadler struct {
	readModel GetUserIdReadModel
}

func NewGetUserIdHandler(readModel GetUserIdReadModel) GetUserIdHadler {
	if readModel == nil {
		panic("GetUserHandler nil readModel")
	}

	return GetUserIdHadler{
		readModel: readModel,
	}
}

type GetUserIdReadModel interface {
	GetUserId(user *auth.User) (int64, error)
}

func (h GetUserIdHadler) Handle(ctx context.Context, query GetUserId) (int64, error) {
	user, err := auth.NewUser(query.Username, query.Password)
	if err != nil {
		return -1, err
	}

	return h.readModel.GetUserId(user)
}
