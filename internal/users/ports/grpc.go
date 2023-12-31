package ports

import (
	"context"

	"github.com/OrIX219/SomethingSocial/internal/common/genproto/users"
	"github.com/OrIX219/SomethingSocial/internal/users/app"
	"github.com/OrIX219/SomethingSocial/internal/users/app/command"
	"github.com/OrIX219/SomethingSocial/internal/users/app/query"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GrpcServer struct {
	app app.Application
}

func NewGrpcServer(app app.Application) GrpcServer {
	return GrpcServer{app}
}

func (g GrpcServer) AddUser(ctx context.Context,
	request *users.AddUserRequest) (*empty.Empty, error) {
	err := g.app.Commands.AddUser.Handle(ctx, command.AddUser{
		UserId: request.UserId,
		Name:   request.Name,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &empty.Empty{}, nil
}

func (g GrpcServer) GetKarma(ctx context.Context,
	request *users.GetKarmaRequest) (*users.GetKarmaResponse, error) {
	karma, err := g.app.Queries.GetKarma.Handle(ctx, query.GetKarma{
		UserId: request.UserId,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &users.GetKarmaResponse{
		Amount: karma,
	}, nil
}

func (g GrpcServer) UpdateKarma(ctx context.Context,
	request *users.UpdateKarmaRequest) (*empty.Empty, error) {
	err := g.app.Commands.UpdateKarma.Handle(ctx, command.UpdateKarma{
		UserId: request.UserId,
		Delta:  request.Delta,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &empty.Empty{}, nil
}

func (g GrpcServer) UpdateLastLogIn(ctx context.Context,
	request *users.UpdateLastLogInRequest) (*empty.Empty, error) {
	err := g.app.Commands.UpdateLastLogIn.Handle(ctx, command.UpdateLastLogIn{
		UserId: request.UserId,
		Time:   request.Time.AsTime(),
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &empty.Empty{}, nil
}

func (g GrpcServer) UpdatePostsCount(ctx context.Context,
	request *users.UpdatePostsCountRequest) (*empty.Empty, error) {
	err := g.app.Commands.UpdatePostsCount.Handle(ctx, command.UpdatePostsCount{
		UserId: request.UserId,
		Delta:  request.Delta,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &empty.Empty{}, nil
}
