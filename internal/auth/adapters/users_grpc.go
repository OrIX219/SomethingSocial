package adapters

import (
	"context"
	"time"

	"github.com/OrIX219/SomethingSocial/internal/common/genproto/users"
	"google.golang.org/protobuf/types/known/timestamppb"
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

func (s UsersGrpc) UpdateLastLogIn(ctx context.Context,
	userId int64, time time.Time) error {
	_, err := s.client.UpdateLastLogIn(ctx, &users.UpdateLastLogInRequest{
		UserId: userId,
		Time:   timestamppb.New(time),
	})

	return err
}
