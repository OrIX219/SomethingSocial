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

func (s UsersGrpc) UpdateKarma(ctx context.Context, userId, delta int64) error {
	_, err := s.client.UpdateKarma(ctx, &users.UpdateKarmaRequest{
		UserId: userId,
		Delta:  delta,
	})

	return err
}

func (s UsersGrpc) GetFollowing(ctx context.Context, userId int64) ([]int64, error) {
	res, err := s.client.GetFollowing(ctx, &users.GetFollowingRequest{
		UserId: userId,
	})

	return res.Users, err
}
