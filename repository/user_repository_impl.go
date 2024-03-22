package repository

import (
	"database/sql"
	"errors"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rulyadhika/my_gram_final_asgmt/model/entity"
	"github.com/rulyadhika/my_gram_final_asgmt/pkg/errs"
)

type UserRepositoryImpl struct{}

func NewUserRepositoryImpl() UserRepository {
	return &UserRepositoryImpl{}
}

func (u *UserRepositoryImpl) Register(ctx *gin.Context, db *sql.DB, user entity.User) (entity.User, error) {
	sqlQuery := `INSERT INTO users(username, email, age, password) VALUES($1, $2, $3, $4)`

	_, err := db.ExecContext(ctx, sqlQuery, user.Username, user.Email, user.Age, user.Password)

	if err != nil {
		log.Printf("[RegisterUser - Repo] err:%s\n", err.Error())
		return user, errs.NewInternalServerError("something went wrong")
	}

	return user, nil
}

func (u *UserRepositoryImpl) CheckEmailAndUsernameUnique(ctx *gin.Context, db *sql.DB, user entity.User) error {
	sqlQuery := `SELECT id FROM users WHERE email=$1`

	rowsEmail, err := db.QueryContext(ctx, sqlQuery, user.Email)

	if err != nil {
		log.Printf("[CheckEmailAndUsernameUnique(email) - Repo] err:%s\n", err.Error())
		return errs.NewInternalServerError("something went wrong")
	}
	defer rowsEmail.Close()

	if rowsEmail.Next() {
		return errs.NewConflictError("email has already been taken")
	}

	sqlQuery = `SELECT id FROM users WHERE username=$1`
	rowsUsername, err := db.QueryContext(ctx, sqlQuery, user.Username)

	if err != nil {
		log.Printf("[CheckEmailAndUsernameUnique(username) - Repo] err:%s\n", err.Error())
		return errs.NewInternalServerError("something went wrong")
	}
	defer rowsUsername.Close()

	if rowsUsername.Next() {
		return errs.NewConflictError("username has already been taken")
	}

	return nil
}

func (u *UserRepositoryImpl) GetUserByEmail(ctx *gin.Context, db *sql.DB, user entity.User) (entity.User, error) {
	sqlQuery := `SELECT id, username, password FROM users WHERE email=$1`

	err := db.QueryRowContext(ctx, sqlQuery, user.Email).Scan(&user.Id, &user.Username, &user.Password)

	if err != nil {
		log.Printf("[GetUserByEmail - Repo] err:%s\n", err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return user, errs.NewNotFoundError("user not found")
		}

		return user, errs.NewInternalServerError("something went wrong")
	}

	return user, nil
}

func (u *UserRepositoryImpl) Update(ctx *gin.Context, db *sql.DB, user entity.User) (entity.User, error) {
	panic("not implemented") // TODO: Implement
}

func (u *UserRepositoryImpl) Delete(ctx *gin.Context, db *sql.DB, userId int) error {
	panic("not implemented") // TODO: Implement
}
