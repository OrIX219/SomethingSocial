package auth

import (
	"context"

	"github.com/OrIX219/SomethingSocial/internal/common/errors"
)

type User struct {
	Id   string
	Role string
}

type ctxKey int

const (
	userContextKey ctxKey = iota
)

var (
	NoUserInContextError = errors.NewAuthorizationError("No user in context",
		"no-user-found")
)

func UserFromCtx(ctx context.Context) (User, error) {
	user, ok := ctx.Value(userContextKey).(User)
	if ok {
		return user, nil
	}

	return User{}, NoUserInContextError
}
