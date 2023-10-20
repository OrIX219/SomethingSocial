package query

import "context"

type UsersService interface {
	GetFollowing(ctx context.Context, userId int64) ([]int64, error)
}
