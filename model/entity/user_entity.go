package entity

import (
	"log"
	"time"

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
