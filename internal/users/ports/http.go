package ports

import (
	"net/http"

	"github.com/OrIX219/SomethingSocial/internal/users/app"
)

type HttpServer struct {
	app app.Application
}

func NewHttpServer(app app.Application) HttpServer {
	return HttpServer{app}
}

func (h HttpServer) GetCurrentUser(w http.ResponseWriter, r *http.Request) {

}

func (h HttpServer) GetUser(w http.ResponseWriter, r *http.Request, userId string) {

}

func (h HttpServer) FollowUser(w http.ResponseWriter, r *http.Request, userId string) {

}

func (h HttpServer) UnfollowUser(w http.ResponseWriter, r *http.Request, userId string) {

}

func (h HttpServer) GetUserFollowers(w http.ResponseWriter, r *http.Request, userId string) {

}

func (h HttpServer) GetUserFollows(w http.ResponseWriter, r *http.Request, userId string) {

}
