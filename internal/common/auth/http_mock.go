package auth

import (
	"context"
	"net/http"

	"github.com/OrIX219/SomethingSocial/internal/common/server/httperr"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

func HttpMockMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var claims jwt.MapClaims
		token, err := request.ParseFromRequest(
			r,
			request.AuthorizationHeaderExtractor,
			func(token *jwt.Token) (any, error) {
				return []byte("mock_secret"), nil
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
			Id:   claims["user_id"].(string),
			Role: claims["role"].(string),
		})
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
