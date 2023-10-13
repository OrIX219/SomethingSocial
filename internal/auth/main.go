package main

import (
	"net/http"

	"github.com/OrIX219/SomethingSocial/internal/auth/adapters"
	"github.com/OrIX219/SomethingSocial/internal/auth/ports"
	"github.com/OrIX219/SomethingSocial/internal/common/logs"
	"github.com/OrIX219/SomethingSocial/internal/common/server"
	"github.com/go-chi/chi/v5"
)

func main() {
	logs.Init()

	repo := adapters.NewUsersInMemoryRepository(10)

	server.RunHTTPServer(func(router chi.Router) http.Handler {
		return ports.HandlerFromMux(ports.NewHttpServer(repo), router)
	})
}
