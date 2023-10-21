package ports

import (
	"net/http"

	"github.com/OrIX219/SomethingSocial/internal/common/auth"
	"github.com/OrIX219/SomethingSocial/internal/common/server/httperr"
	"github.com/OrIX219/SomethingSocial/internal/users/app"
	"github.com/OrIX219/SomethingSocial/internal/users/app/command"
	"github.com/OrIX219/SomethingSocial/internal/users/app/query"
	users "github.com/OrIX219/SomethingSocial/internal/users/domain/user"
	"github.com/go-chi/render"
)

type HttpServer struct {
	app app.Application
}

func NewHttpServer(app app.Application) HttpServer {
	return HttpServer{app}
}

func (h HttpServer) GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	currentUser, err := auth.UserFromCtx(r.Context())
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	h.GetUser(w, r, currentUser.Id)
}

func (h HttpServer) GetUser(w http.ResponseWriter, r *http.Request, userId int64) {
	user, err := h.app.Queries.GetUser.Handle(r.Context(), query.GetUser{
		UserId: userId,
	})
	if err != nil {
		switch err.(type) {
		case users.UserNotFoundError:
			httperr.NotFound("user-not-found", err, w, r)
		default:
			httperr.RespondWithSlugError(err, w, r)
		}
		return
	}

	render.Respond(w, r, marshalUser(user))
}

func (h HttpServer) FollowUser(w http.ResponseWriter, r *http.Request, userId int64) {
	currentUser, err := auth.UserFromCtx(r.Context())
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	err = h.app.Commands.FollowUser.Handle(r.Context(), command.FollowUser{
		UserId:   currentUser.Id,
		TargetId: userId,
	})
	if err != nil {
		switch err.(type) {
		case users.UserNotFoundError:
			render.Respond(w, r, FollowStatus{
				Status: "user not found",
			})
		default:
			httperr.RespondWithSlugError(err, w, r)
		}
		return
	}

	render.Respond(w, r, FollowStatus{
		Status: "ok",
	})
}

func (h HttpServer) UnfollowUser(w http.ResponseWriter, r *http.Request, userId int64) {
	currentUser, err := auth.UserFromCtx(r.Context())
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	err = h.app.Commands.UnfollowUser.Handle(r.Context(), command.UnfollowUser{
		UserId:   currentUser.Id,
		TargetId: userId,
	})
	if err != nil {
		switch err.(type) {
		case users.UserNotFoundError:
			render.Respond(w, r, FollowStatus{
				Status: "user not found",
			})
		default:
			httperr.RespondWithSlugError(err, w, r)
		}
		return
	}

	render.Respond(w, r, FollowStatus{
		Status: "ok",
	})
}

func (h HttpServer) GetUserFollowers(w http.ResponseWriter, r *http.Request, userId int64) {
	followers, err := h.app.Queries.GetFollowers.Handle(r.Context(), query.GetFollowers{
		UserId: userId,
	})
	if err != nil {
		switch err.(type) {
		case users.UserNotFoundError:
			httperr.NotFound("user-not-found", err, w, r)
		default:
			httperr.RespondWithSlugError(err, w, r)
		}
		return
	}

	render.Respond(w, r, responseUserArray(followers))
}

func (h HttpServer) GetUserFollows(w http.ResponseWriter, r *http.Request, userId int64) {
	following, err := h.app.Queries.GetFollowing.Handle(r.Context(), query.GetFollowing{
		UserId: userId,
	})
	if err != nil {
		switch err.(type) {
		case users.UserNotFoundError:
			httperr.NotFound("user-not-found", err, w, r)
		default:
			httperr.RespondWithSlugError(err, w, r)
		}
		return
	}

	render.Respond(w, r, responseUserArray(following))
}

func responseUserArray(users []*users.User) UserArray {
	array := make([]User, len(users))
	for i := range users {
		array[i] = marshalUser(users[i])
	}

	return array
}

func marshalUser(user *users.User) User {
	var role UserRole
	switch user.Role() {
	case users.UserRoleUser:
		role = UserRoleUser
	case users.UserRoleModerator:
		role = UserRoleModerator
	case users.UserRoleAdmin:
		role = UserRoleAdmin
	}

	return User{
		Id:               user.Id(),
		Name:             user.Name(),
		RegistrationDate: user.RegistrationDate(),
		LastLogin:        user.LastLogin(),
		Karma:            user.Karma(),
		PostsCount:       user.PostsCount(),
		Role:             role,
	}
}
