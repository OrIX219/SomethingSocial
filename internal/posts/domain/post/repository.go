package posts

import (
	"fmt"
)

type Repository interface {
	AddPost(post *Post) error
	GetPost(postId string) (*Post, error)
	DeletePost(postId string, userId int64) error
	UpdatePost(postId string, updateFn func(post *Post) (*Post, error)) error
	UpvotePost(postId string, userId int64) (int, error)
	RemoveUpvote(postId string, userId int64) (int, error)
	DownvotePost(postId string, userId int64) (int, error)
	RemoveDownvote(postId string, userId int64) (int, error)
	GetUpvoters(postId string) ([]int64, error)
	GetDownvoters(postId string) ([]int64, error)
	GetAuthor(postId string) (int64, error)
	GetPostsCount(userId int64) (int64, error)
	GetFeed(authors []int64) ([]*Post, error)
	GetPosts(userId int64, filter PostFilter) ([]*Post, error)
}

type PostNotFoundError struct {
	Id string
}

func (e PostNotFoundError) Error() string {
	return fmt.Sprintf("Post %s not found", e.Id)
}
