package ports

import (
	"errors"
	"net/http"

	"github.com/OrIX219/SomethingSocial/internal/auth/app"
	"github.com/OrIX219/SomethingSocial/internal/auth/app/command"
	"github.com/OrIX219/SomethingSocial/internal/auth/app/query"
	auth "github.com/OrIX219/SomethingSocial/internal/auth/domain/user"
	"github.com/OrIX219/SomethingSocial/internal/common/server/httperr"
	"github.com/go-chi/render"
)

type HttpServer struct {
	app app.Application
}

func NewHttpServer(app app.Application) HttpServer {
	return HttpServer{app}
}

func (h HttpServer) SignUp(w http.ResponseWriter, r *http.Request) {
	signUp := SignUp{}
	if err := render.Decode(r, &signUp); err != nil {
		httperr.BadRequest("invalid-request", err, w, r)
		return
	}
	if err := validateSignUpRequest(signUp); err != nil {
		httperr.BadRequest("invalid-request", err, w, r)
		return
	}

	cmd := command.AddUser{
		Name:     signUp.Name,
		Username: signUp.Username,
		Password: signUp.Password,
	}
	err := h.app.Commands.AddUser.Handle(r.Context(), cmd)
	if err != nil {
		switch err.(type) {
		case auth.UsernameExistsError:
			render.Respond(w, r, SignUpResult{
				Status: "username already exists",
			})
		default:
			httperr.RespondWithSlugError(err, w, r)
		}
		return
	}

	id, err := h.app.Queries.GetUserId.Handle(r.Context(), query.GetUserId{
		Username: signUp.Username,
		Password: signUp.Password,
	})
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	render.Respond(w, r, SignUpResult{
		Status: "ok",
		UserId: &id,
	})
}

func validateSignUpRequest(signUp SignUp) error {
	if signUp.Username == "" {
		return errors.New("Empty user username")
	}
	if signUp.Password == "" {
		return errors.New("Empty user password")
	}

	return nil
}

func (h HttpServer) SignIn(w http.ResponseWriter, r *http.Request) {
	signIn := SignIn{}
	if err := render.Decode(r, &signIn); err != nil {
		httperr.BadRequest("invalid-request", err, w, r)
		return
	}
	if err := validateSignInRequest(signIn); err != nil {
		httperr.BadRequest("invalid-request", err, w, r)
		return
	}

	id, err := h.app.Queries.GetUserId.Handle(r.Context(), query.GetUserId{
		Username: signIn.Username,
		Password: signIn.Password,
	})
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

	signedToken, err := h.app.Queries.GenerateToken.Handle(r.Context(),
		query.GenerateToken{
			UserId: id,
		})
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	render.Respond(w, r, SignInResult{
		Status: "ok",
		Token:  &signedToken,
	})
}

func validateSignInRequest(signIn SignIn) error {
	if signIn.Username == "" {
		return errors.New("Empty user username")
	}
	if signIn.Password == "" {
		return errors.New("Empty user password")
	}

	return nil
}
