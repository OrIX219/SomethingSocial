package query

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type GenerateToken struct {
	UserId int64
}

type GenerateTokenHandler struct{}

func NewGenerateTokenHandler() GenerateTokenHandler {
	return GenerateTokenHandler{}
}

func (h GenerateTokenHandler) Handle(query GenerateToken) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		query.UserId,
	})

	return token.SignedString([]byte("mock_secret"))
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int64 `json:"user_id"`
}
