package posts

import (
	"fmt"
)

type Repository interface {
	AddPost(post *Post) error
	GetPost(postId string) (*Post, error)
	DeletePost(postId string, userId int64) error
	UpdatePost(postId string, updateFn func(post *Post) (*Post, error)) error
	GetUpvoters(postId string) ([]int64, error)
	GetDownvoters(postId string) ([]int64, error)
	GetAuthor(postId string) (int64, error)
	GetPostsCount(userId int64) (int64, error)
	GetFeed(userId int64) ([]*Post, error)
	GetPosts(filter PostFilter) ([]*Post, error)
}

type PostNotFoundError struct {
	id string
}

func (e PostNotFoundError) Error() string {
	return fmt.Sprintf("Post %s not found", e.id)
}
