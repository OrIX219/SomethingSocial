package query

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/OrIX219/SomethingSocial/internal/common/auth"
	"github.com/dgrijalva/jwt-go"
)

type GenerateToken struct {
	UserId int64
}

type GenerateTokenHandler struct{}

func NewGenerateTokenHandler() GenerateTokenHandler {
	return GenerateTokenHandler{}
}

func (h GenerateTokenHandler) Handle(ctx context.Context, query GenerateToken) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &auth.JWTClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId: query.UserId,
	})

	if mock, _ := strconv.ParseBool(os.Getenv("MOCK_AUTH")); mock {
		return token.SignedString([]byte("mock_secret"))
	} else {
		return token.SignedString([]byte(os.Getenv("AUTH_SECRET")))
	}
}
