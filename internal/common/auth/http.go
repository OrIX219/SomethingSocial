package auth

import (
	"context"
	"errors"
	"net/http"
	"os"

	commonerrors "github.com/OrIX219/SomethingSocial/internal/common/errors"
	"github.com/OrIX219/SomethingSocial/internal/common/server/httperr"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

type User struct {
	Id   int64
	Role string
}

type ctxKey int

const (
	userContextKey ctxKey = iota
)

func HttpAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var claims JWTClaims
		token, err := request.ParseFromRequest(
			r,
			request.AuthorizationHeaderExtractor,
			func(token *jwt.Token) (any, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, errors.New("Invalid token signing method")
				}
				return []byte(os.Getenv("AUTH_SECRET")), nil
			},
			request.WithClaims(&claims),
		)

		if err != nil {
			httperr.BadRequest("unable-to-get-jwt", err, w, r)
			return
		}

		if !token.Valid {
			httperr.BadRequest("invalid-jwt-token", nil, w, r)
			return
		}

		ctx := context.WithValue(r.Context(), userContextKey, User{
			Id: claims.UserId,
		})
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

var (
	NoUserInContextError = commonerrors.NewAuthorizationError("No user in context", "no-user-found")
)

func UserFromCtx(ctx context.Context) (User, error) {
	user, ok := ctx.Value(userContextKey).(User)
	if ok {
		return user, nil
	}

	return User{}, NoUserInContextError
}

type JWTClaims struct {
	jwt.StandardClaims
	UserId int64 `json:"user_id"`
}
