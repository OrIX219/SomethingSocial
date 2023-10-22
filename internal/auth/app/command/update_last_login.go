package command

import (
	"context"
	"time"
)

type UpdateLastLogIn struct {
	UserId int64
	Time   time.Time
}

type UpdateLastLogInHandler struct {
	usersService UsersService
}

func NewUpdateLastLogInHandler(usersService UsersService) UpdateLastLogInHandler {
	if usersService == nil {
		panic("UpdateLastLogInHandler nil usersService")
	}

	return UpdateLastLogInHandler{
		usersService: usersService,
	}
}

func (h UpdateLastLogInHandler) Handle(ctx context.Context,
	cmd UpdateLastLogIn) error {
	return h.usersService.UpdateLastLogIn(ctx, cmd.UserId, cmd.Time)
}
