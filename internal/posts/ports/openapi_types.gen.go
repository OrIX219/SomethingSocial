// Package ports provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.15.0 DO NOT EDIT.
package ports

import (
	"time"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

const (
	BearerAuthScopes = "bearerAuth.Scopes"
)

// Error defines model for Error.
type Error struct {
	Message string `json:"message"`
	Slug    string `json:"slug"`
}

// Post defines model for Post.
type Post struct {
	Author   int64              `json:"author"`
	Content  string             `json:"content"`
	EditDate *time.Time         `json:"editDate,omitempty"`
	Id       openapi_types.UUID `json:"id"`
	Karma    int64              `json:"karma"`
	PostDate time.Time          `json:"postDate"`
}

// PostArray defines model for PostArray.
type PostArray = []Post

// PostContent defines model for PostContent.
type PostContent struct {
	Content string `json:"content"`
}

// Status defines model for Status.
type Status struct {
	Status string `json:"status"`
}

// Posts defines model for Posts.
type Posts = PostArray

// UnexpectedError defines model for UnexpectedError.
type UnexpectedError = Error

// GetPostsParams defines parameters for GetPosts.
type GetPostsParams struct {
	Author   *string    `form:"author,omitempty" json:"author,omitempty"`
	Limit    *int       `form:"limit,omitempty" json:"limit,omitempty"`
	DateFrom *time.Time `form:"dateFrom,omitempty" json:"dateFrom,omitempty"`
	DateTo   *time.Time `form:"dateTo,omitempty" json:"dateTo,omitempty"`
	Sort     *string    `form:"sort,omitempty" json:"sort,omitempty"`

	// Vote Vote from current user
	Vote *string `form:"vote,omitempty" json:"vote,omitempty"`
}

// CreatePostJSONRequestBody defines body for CreatePost for application/json ContentType.
type CreatePostJSONRequestBody = PostContent

// EditPostJSONRequestBody defines body for EditPost for application/json ContentType.
type EditPostJSONRequestBody = PostContent
