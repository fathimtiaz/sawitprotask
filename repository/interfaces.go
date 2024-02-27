// This file contains the interfaces for the repository layer.
// The repository layer is responsible for interacting with the database.
// For testing purpose we will generate mock implementations of these
// interfaces using mockgen. See the Makefile for more information.
package repository

import (
	"context"

	"github.com/SawitProRecruitment/UserService/domain"
)

type RepositoryInterface interface {
	SaveUser(ctx context.Context, user *domain.User) error
	GetUserByPhone(ctx context.Context, phone string) (user domain.User, err error)
	GetUserById(ctx context.Context, id int64) (user domain.User, err error)
	UpdateUser(ctx context.Context, input UpdateUserInput) (user domain.User, err error)
}
