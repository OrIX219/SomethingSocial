package ports

import (
	"context"

	"github.com/OrIX219/SomethingSocial/internal/common/genproto/posts"
	"github.com/OrIX219/SomethingSocial/internal/posts/app"
	"github.com/OrIX219/SomethingSocial/internal/posts/app/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GrpcServer struct {
	app app.Application
}

func NewGrpcServer(app app.Application) GrpcServer {
	return GrpcServer{
		app: app,
	}
}

func (g GrpcServer) GetUserPostsCount(ctx context.Context,
	request *posts.GetUserPostsCountRequest) (*posts.GetUserPostsCountResponse, error) {
	count, err := g.app.Queries.GetPostsCount.Handle(ctx, query.GetPostsCount{
		UserId: request.UserId,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &posts.GetUserPostsCountResponse{
		Count: count,
	}, nil
}
