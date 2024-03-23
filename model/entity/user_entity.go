package entity

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rulyadhika/my_gram_final_asgmt/app/config"
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

func (u *User) signToken(payload jwt.Claims) (any, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	stringToken, err := token.SignedString([]byte(config.GetAppConfig().JWT_SECRET_KEY))

	if err != nil {
		log.Printf("[SignToken - UserEntity] err:%s\n", err.Error())
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return stringToken, nil
}

func (u *User) GenerateToken() (any, error) {
	payload := u.getJwtClaims()
	fmt.Println(payload)

	stringToken, err := u.signToken(payload)

	if err != nil {
		return nil, err
	}

	return stringToken, nil
}

func (u *User) ParseToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errs.NewUnauthorizedError("invalid token")
		}

		return []byte(config.GetAppConfig().JWT_SECRET_KEY), nil
	})

	if err != nil {
		return err
	}

	return u.BindTokenToUserEntity(token)
}

func (u *User) BindTokenToUserEntity(jwtToken *jwt.Token) error {
	var claims jwt.MapClaims

	if mapClaims, ok := jwtToken.Claims.(jwt.MapClaims); !ok {
		return errs.NewUnauthorizedError("invalid token")
	} else {
		claims = mapClaims
	}

	if id, ok := claims["id"].(float64); !ok {
		return errs.NewUnauthorizedError("invalid token")
	} else {
		u.Id = uint(id)
	}

	if email, ok := claims["email"].(string); !ok {
		return errs.NewUnauthorizedError("invalid token")
	} else {
		u.Email = email
	}

	if username, ok := claims["username"].(string); !ok {
		return errs.NewUnauthorizedError("invalid token")
	} else {
		u.Username = username
	}

	return nil
}
