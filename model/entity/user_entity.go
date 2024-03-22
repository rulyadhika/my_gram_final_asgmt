package entity

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rulyadhika/my_gram_final_asgmt/pkg/errs"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        uint
	Username  string
	Email     string
	Password  string
	Age       uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) HashPassword() error {
	salt := 8

	bs, err := bcrypt.GenerateFromPassword([]byte(u.Password), salt)

	if err != nil {
		log.Printf("[HashPassword - UserEntity] err:%s\n", err.Error())
		return errs.NewInternalServerError("something went wrong")
	}

	u.Password = string(bs)

	return nil
}

func (u *User) ValidatePassword(hashPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(u.Password))

	return err == nil
}

func (u *User) getJwtClaims() jwt.Claims {
	return jwt.MapClaims{
		"id":       u.Id,
		"username": u.Username,
		"email":    u.Email,
	}
}

func (u *User) signToken(payload jwt.Claims, secret string) (any, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	stringToken, err := token.SignedString([]byte(secret))

	if err != nil {
		log.Printf("[SignToken - UserEntity] err:%s\n", err.Error())
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return stringToken, nil
}

func (u *User) GenerateToken(secret string) (any, error) {
	payload := u.getJwtClaims()

	stringToken, err := u.signToken(payload, secret)

	if err != nil {
		return nil, err
	}

	return stringToken, nil
}
