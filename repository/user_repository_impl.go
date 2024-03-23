package repository

import (
	"database/sql"
	"errors"
	"log"
	"time"

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
	var sqlQuery string
	var rowsEmail, rowsUsername *sql.Rows
	var err error

	if user.Id == 0 {
		// this block is used for checking email is unique when on registering a new user

		sqlQuery = `SELECT id FROM users WHERE email=$1`
		rowsEmail, err = db.QueryContext(ctx, sqlQuery, user.Email)
	} else {
		// this block is used for checking email is unique when on updating an existing user

		sqlQuery = `SELECT id FROM users WHERE email=$1 AND NOT id=$2`
		rowsEmail, err = db.QueryContext(ctx, sqlQuery, user.Email, user.Id)
	}

	if err != nil {
		log.Printf("[CheckEmailAndUsernameUnique(email) - Repo] err:%s\n", err.Error())
		return errs.NewInternalServerError("something went wrong")
	}
	defer rowsEmail.Close()

	if rowsEmail.Next() {
		return errs.NewConflictError("email has already been taken")
	}

	if user.Id == 0 {
		// this block is used for checking username is unique when on registering a new user

		sqlQuery = `SELECT id FROM users WHERE username=$1`
		rowsUsername, err = db.QueryContext(ctx, sqlQuery, user.Username)
	} else {
		// this block is used for checking username is unique when on updating an existing user

		sqlQuery = `SELECT id FROM users WHERE username=$1 AND NOT id=$2`
		rowsUsername, err = db.QueryContext(ctx, sqlQuery, user.Username, user.Id)
	}

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

func (u *UserRepositoryImpl) GetUserByEmail(ctx *gin.Context, db *sql.DB, user *entity.User) error {
	sqlQuery := `SELECT id, username, password FROM users WHERE email=$1`

	err := db.QueryRowContext(ctx, sqlQuery, user.Email).Scan(&user.Id, &user.Username, &user.Password)

	if err != nil {
		log.Printf("[GetUserByEmail - Repo] err:%s\n", err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return errs.NewNotFoundError("user not found")
		}

		return errs.NewInternalServerError("something went wrong")
	}

	return nil
}

func (u *UserRepositoryImpl) Update(ctx *gin.Context, db *sql.DB, user entity.User) (entity.User, error) {
	sqlQuery := `UPDATE users SET username=$1, email=$2, age=$3, updated_at=$4 WHERE id=$5 RETURNING updated_at`

	err := db.QueryRowContext(ctx, sqlQuery, user.Username, user.Email, user.Age, time.Now(), user.Id).Scan(&user.UpdatedAt)

	if err != nil {
		log.Printf("[UpdateUser - Repo] err:%s\n", err.Error())
		return user, errs.NewInternalServerError("something went wrong")
	}

	return user, nil
}

func (u *UserRepositoryImpl) GetUserById(ctx *gin.Context, db *sql.DB, userId int) (entity.User, error) {
	sqlQuery := `SELECT id, email, username, password FROM users WHERE id=$1`

	user := entity.User{}

	err := db.QueryRowContext(ctx, sqlQuery, userId).Scan(&user.Id, &user.Email, &user.Username, &user.Password)

	if err != nil {
		log.Printf("[GetUserById - Repo] err:%s\n", err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return user, errs.NewNotFoundError("user not found")
		}

		return user, errs.NewInternalServerError("something went wrong")
	}

	return user, nil
}

func (u *UserRepositoryImpl) Delete(ctx *gin.Context, db *sql.DB, userId int) error {
	var affectedRow int

	sqlQuery := `DELETE FROM users WHERE id=$1 RETURNING id`
	err := db.QueryRowContext(ctx, sqlQuery, userId).Scan(&affectedRow)

	if err != nil {
		log.Printf("[DeleteUser - Repo] err:%s\n", err.Error())
		return errs.NewInternalServerError("something went wrong")
	}

	return nil
}
