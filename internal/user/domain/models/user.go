package models

import (
	"encoding/hex"
	"net/mail"
	"time"

	customErr "github.com/eli95/template/pkg/custom-errors"
	"github.com/eli95/template/pkg/utils"
	"golang.org/x/crypto/sha3"
)

type User struct {
	ID        int64     `json:"id" db:"id"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"password" db:"password"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

func (u *User) Validate() error {
	if u.Email == "" {
		return customErr.ErrEmailRequired
	}
	if u.Password == "" {
		return customErr.ErrPasswordRequired
	}
	if !utils.IsValid(u.Password) {
		return customErr.ErrPasswordFormat
	}
	if u.CreatedAt.IsZero() {
		return customErr.ErrCreateAtRequired
	}
	_, err := mail.ParseAddress(u.Email)
	if err != nil {
		return customErr.ErrInvalidEmail
	}
	return nil
}

func (u *User) Hash256Password(password string) string {
	buf := []byte(password)
	pwd := sha3.New256()
	pwd.Write(buf)
	return hex.EncodeToString(pwd.Sum(nil))
}
