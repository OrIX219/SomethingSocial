package ports

import (
	"net/http"
	"time"

	auth "github.com/OrIX219/SomethingSocial/internal/auth/domain/user"
	"github.com/OrIX219/SomethingSocial/internal/common/server/httperr"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/render"
)

type HttpServer struct {
	repo auth.Repository
}

func NewHttpServer(repo auth.Repository) HttpServer {
	return HttpServer{repo}
}

func (h HttpServer) SignUp(w http.ResponseWriter, r *http.Request) {
	signUp := SignUp{}
	if err := render.Decode(r, &signUp); err != nil {
		httperr.BadRequest("invalid-request", err, w, r)
		return
	}

	// TODO: move all this logic to CQRS when it is implemented
	user, err := auth.NewUser(0, signUp.Username, signUp.Password)
	if err != nil {
		httperr.BadRequest("invalid-request", err, w, r)
		return
	}

	if usr, _ := h.repo.GetUserByUsername(signUp.Username); usr != nil {
		render.Respond(w, r, SignUpResult{
			Status: "username already exists",
		})
		return
	}

	// TODO: hash password before storing
	id, err := h.repo.AddUser(user)

	render.Respond(w, r, SignUpResult{
		Status: "ok",
		UserId: &id,
	})
}

func (h HttpServer) SignIn(w http.ResponseWriter, r *http.Request) {
	signIn := SignIn{}
	if err := render.Decode(r, &signIn); err != nil {
		httperr.BadRequest("invalid-request", err, w, r)
		return
	}

	// TODO: move all this logic to CQRS when it is implemented
	user, err := h.repo.GetUser(signIn.Username, signIn.Password)
	if err != nil {
		switch err.(type) {
		case auth.UserNotFoundError:
			render.Respond(w, r, SignInResult{
				Status: "invalid credentials",
			})
		default:
			httperr.InternalError("internal-error", err, w, r)
		}
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id(),
	})

	signedToken, err := token.SignedString([]byte("mock_secret"))
	if err != nil {
		httperr.InternalError("token-error", err, w, r)
		return
	}

	render.Respond(w, r, SignInResult{
		Status: "ok",
		Token:  &signedToken,
	})
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int64 `json:"user_id"`
}
