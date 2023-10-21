package adapters

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	posts "github.com/OrIX219/SomethingSocial/internal/posts/domain/post"
	"github.com/jmoiron/sqlx"
)

const (
	postsTable     = "posts"
	upvotesTable   = "upvotes"
	downvotesTable = "downvotes"
)

type PostModel struct {
	Id       string     `db:"id"`
	Content  string     `db:"content"`
	PostDate time.Time  `db:"post_date"`
	EditDate *time.Time `db:"edit_date"`
	Karma    int64      `db:"karma"`
	Author   int64      `db:"author"`
}

type PostsPostgresRepository struct {
	db *sqlx.DB
}

func NewPostsPostgresRepository(db *sqlx.DB) *PostsPostgresRepository {
	if db == nil {
		panic("PostsPostgresRepository nil db")
	}

	return &PostsPostgresRepository{
		db: db,
	}
}

func (r *PostsPostgresRepository) AddPost(post *posts.Post) error {
	if post == nil {
		return errors.New("nil post")
	}

	postModel := r.marshalPost(post)

	query := fmt.Sprintf(`INSERT INTO %s (id, content, post_date, karma, author)
		VALUES ($1, $2, $3, $4, $5)`, postsTable)

	_, err := r.db.Exec(query, postModel.Id, postModel.Content,
		postModel.PostDate, postModel.Karma, postModel.Author)

	return err
}

func (r *PostsPostgresRepository) GetPost(postId string) (*posts.Post, error) {
	var postModel PostModel
	query := fmt.Sprintf(`SELECT * FROM %s WHERE id = $1`, postsTable)
	err := r.db.Get(&postModel, query, postId)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, posts.PostNotFoundError{
				Id: postId,
			}
		default:
			return nil, err
		}
	}

	return r.unmarshalPost(postModel)
}

func (r *PostsPostgresRepository) DeletePost(postId string, userId int64) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1 AND author = $2`,
		postsTable)
	res, err := r.db.Exec(query, postId, userId)
	if err != nil {
		return err
	}
	if rows, _ := res.RowsAffected(); rows == 0 {
		return posts.PostNotFoundError{
			Id: postId,
		}
	}

	return nil
}

func (r *PostsPostgresRepository) EditPost(userId int64,
	editedPost *posts.Post) error {
	query := fmt.Sprintf(`UPDATE %s SET content=$1, edit_date=$2
		WHERE id=$3 AND author=$4`, postsTable)
	res, err := r.db.Exec(query,
		editedPost.Content(), editedPost.PostDate(),
		editedPost.Id(), editedPost.Author())
	if err != nil {
		return err
	}
	if rows, _ := res.RowsAffected(); rows == 0 {
		return posts.PostNotFoundError{
			Id: editedPost.Id(),
		}
	}

	return nil
}

func (r *PostsPostgresRepository) UpvotePost(postId string, userId int64) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	karmaDelta := 0

	removeDownvoterQuery := fmt.Sprintf(`DELETE FROM %s
		WHERE post_id=$1 AND user_id=$2`, downvotesTable)
	res, err := tx.Exec(removeDownvoterQuery, postId, userId)
	if err != nil {
		return 0, err
	}
	if rows, _ := res.RowsAffected(); rows > 0 {
		karmaDelta++
	}

	addUpvoterQuery := fmt.Sprintf(`INSERT INTO %s (post_id, user_id) 
		VALUES ($1::uuid, $2::int) EXCEPT SELECT post_id, user_id FROM %[1]s`,
		upvotesTable)
	res, err = tx.Exec(addUpvoterQuery, postId, userId)
	if err != nil {
		return 0, err
	}
	if rows, _ := res.RowsAffected(); rows > 0 {
		karmaDelta++
	}

	return karmaDelta, tx.Commit()
}

func (r *PostsPostgresRepository) RemoveUpvote(postId string, userId int64) (int, error) {
	karmaDelta := 0

	removeUpvoterQuery := fmt.Sprintf(`DELETE FROM %s
		WHERE post_id=$1 AND user_id=$2`, upvotesTable)
	res, err := r.db.Exec(removeUpvoterQuery, postId, userId)
	if err != nil {
		return 0, err
	}
	if rows, _ := res.RowsAffected(); rows > 0 {
		karmaDelta--
	}

	return karmaDelta, nil
}

func (r *PostsPostgresRepository) DownvotePost(postId string, userId int64) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	karmaDelta := 0

	removeUpvoterQuery := fmt.Sprintf(`DELETE FROM %s
		WHERE post_id=$1 AND user_id=$2`, upvotesTable)
	res, err := tx.Exec(removeUpvoterQuery, postId, userId)
	if err != nil {
		return 0, err
	}
	if rows, _ := res.RowsAffected(); rows > 0 {
		karmaDelta--
	}

	addDownvoterQuery := fmt.Sprintf(`INSERT INTO %s (post_id, user_id) 
		VALUES ($1::uuid, $2::int) EXCEPT SELECT post_id, user_id FROM %[1]s`,
		downvotesTable)
	res, err = tx.Exec(addDownvoterQuery, postId, userId)
	if err != nil {
		return 0, err
	}
	if rows, _ := res.RowsAffected(); rows > 0 {
		karmaDelta--
	}

	return karmaDelta, tx.Commit()
}

func (r *PostsPostgresRepository) RemoveDownvote(postId string, userId int64) (int, error) {
	karmaDelta := 0

	removeDownvoterQuery := fmt.Sprintf(`DELETE FROM %s
		WHERE post_id=$1 AND user_id=$2`, downvotesTable)
	res, err := r.db.Exec(removeDownvoterQuery, postId, userId)
	if err != nil {
		return 0, err
	}
	if rows, _ := res.RowsAffected(); rows > 0 {
		karmaDelta++
	}

	return karmaDelta, nil
}

func (r *PostsPostgresRepository) GetUpvoters(postId string) ([]int64, error) {
	var upvoters []int64
	query := fmt.Sprintf(`SELECT u.user_id FROM %s u
		INNER JOIN %s p ON u.post_id = p.id WHERE p.id = $1`,
		upvotesTable, postsTable)
	err := r.db.Select(&upvoters, query, postId)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, posts.PostNotFoundError{
				Id: postId,
			}
		default:
			return nil, err
		}
	}

	return upvoters, nil
}

func (r *PostsPostgresRepository) GetDownvoters(postId string) ([]int64, error) {
	var downvoters []int64
	query := fmt.Sprintf(`SELECT d.user_id FROM %s d
		INNER JOIN %s p ON d.post_id = p.id WHERE p.id = $1`,
		downvotesTable, postsTable)
	err := r.db.Select(&downvoters, query, postId)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, posts.PostNotFoundError{
				Id: postId,
			}
		default:
			return nil, err
		}
	}

	return downvoters, nil
}

func (r *PostsPostgresRepository) GetAuthor(postId string) (int64, error) {
	var author int64
	query := fmt.Sprintf(`SELECT author FROM %s WHERE id = $1`, postsTable)
	err := r.db.Get(&author, query, postId)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return -1, posts.PostNotFoundError{
				Id: postId,
			}
		default:
			return -1, err
		}
	}

	return author, nil
}

func (r *PostsPostgresRepository) GetPostsCount(userId int64) (int64, error) {
	var count int64
	query := fmt.Sprintf(`SELECT COUNT(*) FROM %s WHERE author = $1`, postsTable)
	err := r.db.Get(&count, query, userId)

	return count, err
}

func (r *PostsPostgresRepository) GetFeed(authors []int64) ([]*posts.Post, error) {
	if len(authors) == 0 {
		return []*posts.Post{}, nil
	}

	setValues := make([]string, 0, len(authors))
	for i := range authors {
		setValues = append(setValues, fmt.Sprintf("(%d)", authors[i]))
	}
	tempTable := fmt.Sprintf("(values %s)", strings.Join(setValues, ","))

	var postModels []PostModel
	query := fmt.Sprintf(`SELECT p.* FROM %s p
		JOIN %s AS t(id) ON p.author = t.id`, postsTable, tempTable)
	err := r.db.Select(&postModels, query)
	if err != nil {
		return nil, err
	}

	feed := make([]*posts.Post, 0, len(postModels))
	for i := range postModels {
		post, err := r.unmarshalPost(postModels[i])
		if err == nil {
			feed = append(feed, post)
		}
	}

	return feed, nil
}

func (r *PostsPostgresRepository) GetPosts(userId int64,
	filter posts.PostFilter) ([]*posts.Post, error) {
	filters := []string{}
	args := []any{}
	argId := 1
	if filter.Author != nil {
		filters = append(filters, fmt.Sprintf("author = $%d", argId))
		args = append(args, *filter.Author)
		argId++
	}
	if filter.DateFrom != nil {
		filters = append(filters,
			fmt.Sprintf("post_date >= $%d", argId))
		args = append(args, filter.DateFrom)
		argId++
	}
	if filter.DateTo != nil {
		filters = append(filters,
			fmt.Sprintf("post_date <= $%d", argId))
		args = append(args, filter.DateTo)
		argId++
	}
	if filter.Vote != nil {
		switch *filter.Vote {
		case "up":
			filters = append(filters, fmt.Sprintf("u.user_id = $%d", argId))
		case "down":
			filters = append(filters, fmt.Sprintf("d.user_id = $%d", argId))
		}
		args = append(args, userId)
		argId++
	}

	strArray := make([]string, 0, 3) // 3 = filters + order + limit
	if len(filters) > 0 {
		strArray =
			append(strArray, fmt.Sprintf("WHERE %s", strings.Join(filters, " AND ")))
	}

	if filter.Sort != nil {
		switch *filter.Sort {
		case "newest":
			strArray = append(strArray, "ORDER BY post_date DESC")
		case "oldest":
			strArray = append(strArray, "ORDER BY post_date ASC")
		case "upvoted":
			strArray = append(strArray, "ORDER BY karma DESC")
		case "downvoted":
			strArray = append(strArray, "ORDER BY karma ASC")
		}
	}

	if filter.Limit != nil {
		strArray = append(strArray, fmt.Sprintf("LIMIT $%d", argId))
		args = append(args, *filter.Limit)
	}

	var postModels []PostModel
	query := fmt.Sprintf(`SELECT p.* FROM %s p
		LEFT JOIN %s u ON p.id = u.post_id
		LEFT JOIN %s d ON p.id = d.post_id
		%s`,
		postsTable, upvotesTable, downvotesTable, strings.Join(strArray, " "))
	err := r.db.Select(&postModels, query, args...)
	if err != nil {
		return nil, err
	}

	posts := make([]*posts.Post, 0, len(postModels))
	for i := range postModels {
		post, err := r.unmarshalPost(postModels[i])
		if err == nil {
			posts = append(posts, post)
		}
	}

	return posts, nil
}

func (r *PostsPostgresRepository) marshalPost(post *posts.Post) PostModel {
	return PostModel{
		Id:       post.Id(),
		Content:  post.Content(),
		PostDate: post.PostDate(),
		EditDate: post.EditDate(),
		Karma:    post.Karma(),
		Author:   post.Author(),
	}
}

func (r *PostsPostgresRepository) unmarshalPost(post PostModel) (*posts.Post, error) {
	return posts.UnmarshalFromRepository(post.Id, post.Content, post.PostDate,
		post.EditDate, post.Karma, post.Author)
}
