package client

import (
	"crypto/tls"
	"crypto/x509"
	"os"
	"strconv"

	"github.com/OrIX219/SomethingSocial/internal/common/genproto/posts"
	"github.com/OrIX219/SomethingSocial/internal/common/genproto/users"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

func NewUsersClient() (client users.UsersServiceClient,
	close func() error, err error) {
	grpcAddr := os.Getenv("USERS_GRPC_ADDR")
	if grpcAddr == "" {
		return nil, func() error { return nil },
			errors.New("empty env USERS_GRPC_ADDR")
	}

	opts, err := grpcDialOpts(grpcAddr)
	if err != nil {
		return nil, func() error { return nil }, err
	}

	conn, err := grpc.Dial(grpcAddr, opts...)
	if err != nil {
		return nil, func() error { return nil }, err
	}

	return users.NewUsersServiceClient(conn), conn.Close, nil
}

func NewPostsClient() (client posts.PostsServiceClient,
	close func() error, err error) {
	grpcAddr := os.Getenv("POSTS_GRPC_ADDR")
	if grpcAddr == "" {
		return nil, func() error { return nil },
			errors.New("empty env POSTS_GRPC_ADDR")
	}

	opts, err := grpcDialOpts(grpcAddr)
	if err != nil {
		return nil, func() error { return nil }, err
	}

	conn, err := grpc.Dial(grpcAddr, opts...)
	if err != nil {
		return nil, func() error { return nil }, err
	}

	return posts.NewPostsServiceClient(conn), conn.Close, nil
}

func grpcDialOpts(grpcAddr string) ([]grpc.DialOption, error) {
	if noTLS, _ := strconv.ParseBool(os.Getenv("GRPC_NO_TLS")); noTLS {
		return []grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		}, nil
	}

	systemRoots, err := x509.SystemCertPool()
	if err != nil {
		return nil, errors.Wrap(err, "Cannot load root CA cert")
	}
	creds := credentials.NewTLS(&tls.Config{
		RootCAs:    systemRoots,
		MinVersion: tls.VersionTLS12,
	})

	return []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
	}, nil
}
