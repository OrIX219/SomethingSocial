package adapters

import (
	"context"

	"github.com/OrIX219/SomethingSocial/internal/common/genproto/users"
)

type UsersGrpc struct {
	client users.UsersServiceClient
}

func NewUsersGrpc(client users.UsersServiceClient) UsersGrpc {
	return UsersGrpc{
		client: client,
	}
}

func (s UsersGrpc) AddUser(ctx context.Context, userId int64, name string) error {
	_, err := s.client.AddUser(ctx, &users.AddUserRequest{
		UserId: userId,
		Name:   name,
	})

	return err
}
