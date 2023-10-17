package server

import (
	"net/http"
	"os"

	"github.com/OrIX219/SomethingSocial/internal/common/logs"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sirupsen/logrus"
)

func RunHTTPServer(createHandler func(router chi.Router) http.Handler,
	authMiddleware func(http.Handler) http.Handler) {
	RunHTTPServerOnAddr(":"+os.Getenv("PORT"), createHandler, authMiddleware)
}

func RunHTTPServerOnAddr(addr string,
	createHandler func(router chi.Router) http.Handler,
	authMiddleware func(http.Handler) http.Handler) {
	apiRouter := chi.NewRouter()
	setMiddlewares(apiRouter, authMiddleware)

	rootRouter := chi.NewRouter()
	rootRouter.Mount("/api", createHandler(apiRouter))

	logrus.WithField("port", addr).Info("Starting HTTP server")

	err := http.ListenAndServe(addr, rootRouter)
	if err != nil {
		logrus.WithError(err).Panic("Unable to start HTTP server")
	}
}

func setMiddlewares(router *chi.Mux,
	authMiddleware func(http.Handler) http.Handler) {
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(logs.NewStructuredLogger(logrus.StandardLogger()))
	router.Use(middleware.Recoverer)

	router.Use(authMiddleware)

	router.Use(
		middleware.SetHeader("X-Content-Type-Options", "nosniff"),
		middleware.SetHeader("X-Frame-Options", "deny"),
	)
	router.Use(middleware.NoCache)
}
