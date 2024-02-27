package domain

import (
	"context"
	"errors"
	"fmt"
	"reflect"

	"github.com/SawitProRecruitment/UserService/helper/jwt"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       int64
	Name     string
	Phone    string
	Password string
}

const UserIdCtxKey = "user_id"

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
	return jwt.GenerateUserToken(u.Id, u.Phone, u.Name)
}

func UserIdCtx(ctx context.Context) (id int64, err error) {
	fmt.Println(reflect.TypeOf(ctx.Value(UserIdCtxKey)))
	ids, ok := ctx.Value(UserIdCtxKey).(float64)
	if !ok {
		err = errors.New("invalid user ctx")
		return
	}

	id = int64(ids)

	return
}
