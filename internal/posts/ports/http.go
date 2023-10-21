package ports

import (
	"errors"
	"net/http"
	"time"

	"github.com/OrIX219/SomethingSocial/internal/common/auth"
	"github.com/OrIX219/SomethingSocial/internal/common/server/httperr"
	"github.com/OrIX219/SomethingSocial/internal/posts/app"
	"github.com/OrIX219/SomethingSocial/internal/posts/app/command"
	"github.com/OrIX219/SomethingSocial/internal/posts/app/query"
	posts "github.com/OrIX219/SomethingSocial/internal/posts/domain/post"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

type HttpServer struct {
	app app.Application
}

func NewHttpServer(app app.Application) HttpServer {
	return HttpServer{
		app: app,
	}
}

func (h HttpServer) GetPosts(w http.ResponseWriter, r *http.Request,
	params GetPostsParams) {
	currentUser, err := auth.UserFromCtx(r.Context())
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	filter := posts.PostFilter{
		Author:   params.Author,
		Limit:    params.Limit,
		DateFrom: params.DateFrom,
		DateTo:   params.DateTo,
		Sort:     params.Sort,
		Vote:     params.Vote,
	}
	posts, err := h.app.Queries.GetPosts.Handle(r.Context(), query.GetPosts{
		UserId: currentUser.Id,
		Filter: filter,
	})
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	render.Respond(w, r, responsePostArray(posts))
}

func (h HttpServer) CreatePost(w http.ResponseWriter, r *http.Request) {
	currentUser, err := auth.UserFromCtx(r.Context())
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	createPost := CreatePost{}
	if err := render.Decode(r, &createPost); err != nil {
		httperr.BadRequest("invalid-request", err, w, r)
		return
	}
	if err := validateCreatePostRequest(&createPost); err != nil {
		httperr.BadRequest("invalid-request", err, w, r)
		return
	}

	cmd := command.CreatePost{
		PostId:   uuid.New().String(),
		Content:  createPost.Content,
		PostDate: time.Now(),
		Author:   currentUser.Id,
	}
	err = h.app.Commands.CreatePost.Handle(r.Context(), cmd)
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	w.Header().Set("Location", "/posts/"+cmd.PostId)
	w.WriteHeader(http.StatusCreated)
}

func (h HttpServer) GetFeed(w http.ResponseWriter, r *http.Request) {
	currentUser, err := auth.UserFromCtx(r.Context())
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	posts, err := h.app.Queries.GetFeed.Handle(r.Context(), query.GetFeed{
		UserId: currentUser.Id,
	})
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	render.Respond(w, r, responsePostArray(posts))
}

func (h HttpServer) DeletePost(w http.ResponseWriter, r *http.Request,
	postId openapi_types.UUID) {
	currentUser, err := auth.UserFromCtx(r.Context())
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	cmd := command.DeletePost{
		PostId: postId.String(),
		UserId: currentUser.Id,
	}
	err = h.app.Commands.DeletePost.Handle(r.Context(), cmd)
	if err != nil {
		switch err.(type) {
		case posts.PostNotFoundError:
			httperr.NotFound("post-not-found", err, w, r)
		default:
			httperr.RespondWithSlugError(err, w, r)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h HttpServer) GetPost(w http.ResponseWriter, r *http.Request,
	postId openapi_types.UUID) {
	post, err := h.app.Queries.GetPost.Handle(r.Context(), query.GetPost{
		PostId: postId.String(),
	})
	if err != nil {
		switch err.(type) {
		case posts.PostNotFoundError:
			httperr.NotFound("post-not-found", err, w, r)
		default:
			httperr.RespondWithSlugError(err, w, r)
		}
		return
	}

	render.Respond(w, r, marshalPost(post))
}

func (h HttpServer) DownvotePost(w http.ResponseWriter, r *http.Request,
	postId openapi_types.UUID) {
	currentUser, err := auth.UserFromCtx(r.Context())
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	err = h.app.Commands.DownvotePost.Handle(r.Context(), command.DownvotePost{
		PostId: postId.String(),
		UserId: currentUser.Id,
	})
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	render.Respond(w, r, Status{
		Status: "ok",
	})
}

func (h HttpServer) UpvotePost(w http.ResponseWriter, r *http.Request,
	postId openapi_types.UUID) {
	currentUser, err := auth.UserFromCtx(r.Context())
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	err = h.app.Commands.UpvotePost.Handle(r.Context(), command.UpvotePost{
		PostId: postId.String(),
		UserId: currentUser.Id,
	})
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	render.Respond(w, r, Status{
		Status: "ok",
	})
}

func validateCreatePostRequest(createPost *CreatePost) error {
	if createPost.Content == "" {
		return errors.New("Empty post content")
	}

	return nil
}

func responsePostArray(posts []*posts.Post) PostArray {
	array := make([]Post, len(posts))
	for i := range posts {
		array[i] = marshalPost(posts[i])
	}

	return array
}

func marshalPost(post *posts.Post) Post {
	id, _ := uuid.Parse(post.Id())
	return Post{
		Id:       id,
		Content:  post.Content(),
		PostDate: post.PostDate(),
		Karma:    post.Karma(),
		Author:   post.Author(),
	}
}
