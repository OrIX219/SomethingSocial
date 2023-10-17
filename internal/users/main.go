package main

import (
	"net/http"
	"os"
	"strings"

	"github.com/OrIX219/SomethingSocial/internal/common/auth"
	"github.com/OrIX219/SomethingSocial/internal/common/genproto/users"
	"github.com/OrIX219/SomethingSocial/internal/common/logs"
	"github.com/OrIX219/SomethingSocial/internal/common/server"
	"github.com/OrIX219/SomethingSocial/internal/users/ports"
	"github.com/OrIX219/SomethingSocial/internal/users/service"
	"github.com/go-chi/chi/v5"
	"google.golang.org/grpc"
)

func main() {
	logs.Init()

	app := service.NewApplication()

	serverType := strings.ToLower(os.Getenv("SERVER_TYPE"))
	switch serverType {
	case "http":
		server.RunHTTPServer(func(router chi.Router) http.Handler {
			return ports.HandlerFromMux(ports.NewHttpServer(app), router)
		}, auth.HttpMockMiddleware)
	case "grpc":
		server.RunGRPCServer(func(server *grpc.Server) {
			srv := ports.NewGrpcServer(app)
			users.RegisterUsersServiceServer(server, srv)
		})
	}

}
