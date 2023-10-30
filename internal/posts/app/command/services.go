package command

import "context"

type UsersService interface {
	UpdateKarma(ctx context.Context, userId, delta int64) error
	UpdatePostsCount(ctx context.Context, userId, delta int64) error
}
