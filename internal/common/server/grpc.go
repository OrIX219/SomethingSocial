package server

import (
	"fmt"
	"net"
	"os"

	"github.com/OrIX219/SomethingSocial/internal/common/logs"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func RunGRPCServer(registerServer func(server *grpc.Server)) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := fmt.Sprintf(":%s", port)
	RunGRPCServerOnAddr(addr, registerServer)
}

func RunGRPCServerOnAddr(addr string, registerServer func(server *grpc.Server)) {
	logrusEntry := logrus.NewEntry(logrus.StandardLogger())
	interceptorLogger := logs.NewInterceptorLogger(logrusEntry)

	opts := []logging.Option{}

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			logging.UnaryServerInterceptor(interceptorLogger, opts...),
		),
		grpc.ChainStreamInterceptor(
			logging.StreamServerInterceptor(interceptorLogger, opts...),
		),
	)
	registerServer(grpcServer)

	listen, err := net.Listen("tcp", addr)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.WithField("grpcEndpoint", addr).Info("Starting gRPC listener")
	logrus.Fatal(grpcServer.Serve(listen))
}
