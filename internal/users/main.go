package main

import (
	"net/http"

	"github.com/OrIX219/SomethingSocial/internal/common/logs"
	"github.com/OrIX219/SomethingSocial/internal/common/server"
	"github.com/OrIX219/SomethingSocial/internal/users/ports"
	"github.com/OrIX219/SomethingSocial/internal/users/service"
	"github.com/go-chi/chi/v5"
)

func main() {
	logs.Init()

	app := service.NewApplication()

	server.RunHTTPServer(func(router chi.Router) http.Handler {
		return ports.HandlerFromMux(ports.NewHttpServer(app), router)
	})
}
