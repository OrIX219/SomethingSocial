package server

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net"
	"os"
	"strconv"

	"github.com/OrIX219/SomethingSocial/internal/common/logs"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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
	opts, err := grpcServerOptions()
	if err != nil {
		logrus.Fatal(err)
	}

	grpcServer := grpc.NewServer(opts...)
	registerServer(grpcServer)

	listen, err := net.Listen("tcp", addr)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.WithField("grpcEndpoint", addr).Info("Starting gRPC listener")
	logrus.Fatal(grpcServer.Serve(listen))
}

func grpcServerOptions() ([]grpc.ServerOption, error) {
	logrusEntry := logrus.NewEntry(logrus.StandardLogger())
	interceptorLogger := logs.NewInterceptorLogger(logrusEntry)

	opts := []logging.Option{}
	serverOptions := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			logging.UnaryServerInterceptor(interceptorLogger, opts...),
		),
		grpc.ChainStreamInterceptor(
			logging.StreamServerInterceptor(interceptorLogger, opts...),
		),
	}

	if noTLS, _ := strconv.ParseBool(os.Getenv("GRPC_NO_TLS")); noTLS {
		return serverOptions, nil
	}

	systemRoots, err := x509.SystemCertPool()
	if err != nil {
		return nil, errors.Wrap(err, "Cannot load root CA cert")
	}

	creds := grpc.Creds(credentials.NewTLS(&tls.Config{
		RootCAs:    systemRoots,
		MinVersion: tls.VersionTLS12,
	}))
	serverOptions = append(serverOptions, creds)

	return serverOptions, nil
}
