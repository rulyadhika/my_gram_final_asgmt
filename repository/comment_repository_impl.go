package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rulyadhika/my_gram_final_asgmt/model/entity"
	"github.com/rulyadhika/my_gram_final_asgmt/pkg/errs"
)

type CommentRepositoryImpl struct{}

func NewCommentRepositoryImpl() CommentRepository {
	return &CommentRepositoryImpl{}
}

func (c *CommentRepositoryImpl) FindAll(ctx *gin.Context, db *sql.DB) ([]CommentPhotoUser, error) {
	sqlQuery := `SELECT comments.id, comments.message, comments.photo_id, comments.user_id, comments.created_at, comments.updated_at,
	photos.id, photos.title, photos.caption, photos.photo_url, photos.user_id, photos.created_at, photos.updated_at,
	users.id, users.email, users.username, photo_user.id, photo_user.email, photo_user.username
	FROM comments JOIN photos ON comments.photo_id = photos.id
	JOIN users ON comments.user_id=users.id JOIN users AS photo_user ON photos.user_id = photo_user.id`

	rows, err := db.QueryContext(ctx, sqlQuery)
	if err != nil {
		log.Printf("[FindAllComment - Repo] err:%s\n", err.Error())
		return []CommentPhotoUser{}, err
	}
	defer rows.Close()

	comments := []CommentPhotoUser{}

	for rows.Next() {
		comment := CommentPhotoUser{}
		err := rows.Scan(&comment.Comment.Id, &comment.Message, &comment.PhotoId, &comment.Comment.UserId, &comment.Comment.CreatedAt,
			&comment.Comment.UpdatedAt, &comment.PhotoUser.Photo.Id, &comment.PhotoUser.Photo.Title, &comment.PhotoUser.Photo.Caption,
			&comment.PhotoUser.Photo.PhotoUrl, &comment.PhotoUser.Photo.UserId, &comment.PhotoUser.Photo.CreatedAt, &comment.PhotoUser.Photo.UpdatedAt,
			&comment.User.Id, &comment.User.Email, &comment.User.Username, &comment.PhotoUser.User.Id, &comment.PhotoUser.User.Email, &comment.PhotoUser.User.Username)

		if err != nil {
			log.Printf("[FindAllComment - Repo] err:%s\n", err.Error())
			return []CommentPhotoUser{}, err
		}

		comments = append(comments, comment)
	}

	return comments, nil
}

func (c *CommentRepositoryImpl) FindById(ctx *gin.Context, db *sql.DB, commentId int) (entity.Comment, error) {
	sqlQuery := `SELECT id, user_id, photo_id, message, created_at, updated_at FROM comments WHERE id=$1`

	comment := entity.Comment{}

	err := db.QueryRowContext(ctx, sqlQuery, commentId).Scan(&comment.Id, &comment.UserId, &comment.PhotoId, &comment.Message, &comment.CreatedAt, &comment.UpdatedAt)

	if err != nil {
		log.Printf("[FindCommentById - Repo] err: %s \n", err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return comment, errs.NewNotFoundError(fmt.Sprintf("social media with id:%v not found", commentId))
		}

		return comment, errs.NewInternalServerError("something went wrong")
	}

	return comment, nil
}

func (c *CommentRepositoryImpl) Create(ctx *gin.Context, db *sql.DB, comment entity.Comment) (entity.Comment, error) {
	sqlQuery := `INSERT INTO comments(message, photo_id, user_id) VALUES($1, $2, $3) RETURNING id, created_at`

	err := db.QueryRowContext(ctx, sqlQuery, comment.Message, comment.PhotoId, comment.UserId).Scan(&comment.Id, &comment.CreatedAt)

	if err != nil {
		log.Printf("[CreateComment - repo] err: %s\n", err.Error())
		return comment, errs.NewInternalServerError("something went wrong")
	}

	return comment, nil
}

func (c *CommentRepositoryImpl) Update(ctx *gin.Context, db *sql.DB, comment entity.Comment) (entity.Comment, error) {
	sqlQuery := `UPDATE comments SET message=$1, updated_at=$2 WHERE id=$3 RETURNING photo_id, user_id, updated_at`

	err := db.QueryRowContext(ctx, sqlQuery, comment.Message, time.Now(), comment.Id).Scan(&comment.PhotoId, &comment.UserId, &comment.UpdatedAt)

	if err != nil {
		log.Printf("[UpdateComment - Repo] err: %s\n", err.Error())

		if errors.Is(err, sql.ErrNoRows) {
			return comment, errs.NewNotFoundError(fmt.Sprintf("comment with id:%v not found", comment.Id))
		}

		return comment, errs.NewInternalServerError("something went wrong")
	}

	return comment, nil
}

func (c *CommentRepositoryImpl) Delete(ctx *gin.Context, db *sql.DB, commentId int) error {
	sqlQuery := `DELETE FROM comments WHERE id=$1 RETURNING id`

	var affectedRow int
	err := db.QueryRowContext(ctx, sqlQuery, commentId).Scan(&affectedRow)

	if err != nil {
		log.Printf("[DeleteComment - Repo] err: %s\n", err.Error())

		if errors.Is(err, sql.ErrNoRows) {
			return errs.NewNotFoundError(fmt.Sprintf("comment with id:%v not found", commentId))
		}

		return errs.NewInternalServerError("something went wrong")
	}

	return nil
}
