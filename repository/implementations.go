package repository

import (
	"context"
)

func (r *Repository) SaveUser(ctx context.Context, user *User) (err error) {
	err = r.Db.QueryRowContext(ctx, `
		INSERT INTO user_ (full_name, phone, password) 
		VALUES ($1, $2, $3) RETURNING id
		`, user.Name, user.Phone, user.Password).Scan(&user.Id)
	if err != nil {
		return err
	}

	return
}

func (r *Repository) GetUserByPhone(ctx context.Context, phone string) (user User, err error) {
	err = r.Db.QueryRowContext(ctx, `
		SELECT id, full_name, phone, password FROM user_
		WHERE phone = $1
		`, phone).Scan(&user.Id, &user.Name, &user.Phone, &user.Password)
	if err != nil {
		return
	}

	return
}
