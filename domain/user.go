package domain

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/SawitProRecruitment/UserService/helper/jwt"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       int64
	Name     string
	Phone    string
	Password string
}

const userCtxKey = "user_data"

func newPublicUser(user User) User {
	return User{
		Id:    user.Id,
		Name:  user.Name,
		Phone: user.Phone,
	}
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

func (u *User) GenerateToken() (token string, err error) {
	userData, err := json.Marshal(newPublicUser(*u))
	if err != nil {
		return
	}

	return jwt.GenerateUserToken(userData)
}

func NewAuthenticatedUser(ctx context.Context) (result User, err error) {
	authdValue, ok := ctx.Value(userCtxKey).([]byte)
	if !ok {
		err = errors.New("invalid user ctx")
		return
	}

	if err = json.Unmarshal(authdValue, &result); err != nil {
		return
	}

	return
}
