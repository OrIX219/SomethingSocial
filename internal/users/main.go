package main

import (
	"net/http"
	"os"
	"strconv"
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
		var authMiddleware func(http.Handler) http.Handler
		if mock, _ := strconv.ParseBool(os.Getenv("MOCK_AUTH")); mock {
			authMiddleware = auth.HttpMockMiddleware
		} else {
			authMiddleware = auth.HttpAuthMiddleware
		}
		server.RunHTTPServer(func(router chi.Router) http.Handler {
			return ports.HandlerFromMux(ports.NewHttpServer(app), router)
		}, authMiddleware)
	case "grpc":
		server.RunGRPCServer(func(server *grpc.Server) {
			srv := ports.NewGrpcServer(app)
			users.RegisterUsersServiceServer(server, srv)
		})
	}

}
