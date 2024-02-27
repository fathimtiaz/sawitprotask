package repository

import (
	"context"
	"fmt"

	"github.com/SawitProRecruitment/UserService/domain"
)

func (r *Repository) SaveUser(ctx context.Context, user *domain.User) (err error) {
	err = r.Db.QueryRowContext(ctx, `
		INSERT INTO user_ (full_name, phone, password) 
		VALUES ($1, $2, $3) RETURNING id
		`, user.Name, user.Phone, user.Password).Scan(&user.Id)
	if err != nil {
		return err
	}

	return
}

func (r *Repository) GetUserByPhone(ctx context.Context, phone string) (user domain.User, err error) {
	err = r.Db.QueryRowContext(ctx, `
		SELECT id, full_name, phone, password FROM user_
		WHERE phone = $1
		`, phone).Scan(&user.Id, &user.Name, &user.Phone, &user.Password)
	if err != nil {
		return
	}

	return
}

func (r *Repository) GetUserById(ctx context.Context, id int64) (user domain.User, err error) {
	err = r.Db.QueryRowContext(ctx, `
		SELECT id, full_name, phone, password FROM user_
		WHERE id = $1
		`, id).Scan(&user.Id, &user.Name, &user.Phone, &user.Password)
	if err != nil {
		return
	}

	return
}

func (r *Repository) UpdateUser(ctx context.Context, input UpdateUserInput) (err error) {
	var setQ string = "SET "
	var whereQ string = "WHERE phone = $"
	var argCount = 0
	var args []interface{}

	if input.Phone != "" {
		argCount++
		setQ += fmt.Sprintf("phone = $%d", argCount)
		args = append(args, input.Phone)
	}

	if input.FullName != "" {
		argCount++
		setQ += fmt.Sprintf("full_name = $%d", argCount)
		args = append(args, input.FullName)
	}

	setQ += "\n"

	argCount++
	whereQ += fmt.Sprintf("%d", argCount)
	args = append(args, input.WhereId)

	return r.Db.QueryRowContext(ctx, `UPDATE user_`+setQ+whereQ, args...).Scan()
}
