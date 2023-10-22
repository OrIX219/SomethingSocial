package command

import (
	"context"
	"time"
)

type UsersService interface {
	AddUser(ctx context.Context, userId int64, name string) error
	UpdateLastLogIn(ctx context.Context, userId int64, time time.Time) error
}
