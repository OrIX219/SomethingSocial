package command

import "context"

type UsersService interface {
	AddUser(ctx context.Context, userId int64, name string) error
}
