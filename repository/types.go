// This file contains types that are used in the repository layer.
package repository

import (
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       int64
	Name     string
	Phone    string
	Password string
}

func (u *User) HashPassword() error {
	if hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost); err != nil {
		return err
	} else {
		u.Password = string(hashed)
		return nil
	}
}

func (u *User) ComparePassword(passAttempt string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(passAttempt))
}

func (u *User) GenerateToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": u,
	})

	return token.SignedString([]byte("hmacSampleSecret"))
}
