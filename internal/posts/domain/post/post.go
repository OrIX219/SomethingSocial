package posts

import (
	"errors"
	"time"
)

type Post struct {
	id       string
	content  string
	postDate time.Time
	editDate *time.Time
	karma    int64
	author   int64
}

func NewPost(id, content string, postDate time.Time, author int64) (*Post, error) {
	if err := validatePostData(id, content, postDate, author); err != nil {
		return nil, err
	}

	return &Post{
		id:       id,
		content:  content,
		postDate: postDate,
		editDate: nil,
		karma:    0,
		author:   author,
	}, nil
}

func (p Post) Id() string {
	return p.id
}

func (p Post) Content() string {
	return p.content
}

func (p Post) PostDate() time.Time {
	return p.postDate
}

func (p Post) EditDate() *time.Time {
	return p.editDate
}

func (p Post) Karma() int64 {
	return p.karma
}

func (p Post) Author() int64 {
	return p.author
}

func (p *Post) Upvote() {
	p.karma++
}

func (p *Post) Downvote() {
	p.karma--
}

func UnmarshalFromRepository(id, content string, postDate time.Time,
	editDate *time.Time, karma, author int64) (*Post, error) {
	post, err := NewPost(id, content, postDate, author)
	if err != nil {
		return nil, err
	}
	post.editDate = editDate
	post.karma = karma

	return post, nil
}

func validatePostData(id, content string, postDate time.Time, author int64) error {
	if id == "" {
		return errors.New("Empty post id")
	}
	if content == "" {
		return errors.New("Empty post content")
	}
	if postDate.After(time.Now()) {
		return errors.New("Invalid post date")
	}
	if author < 1 {
		return errors.New("Invalid author id")
	}

	return nil
}
