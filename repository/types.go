// This file contains types that are used in the repository layer.
package repository

type UpdateUserInput struct {
	WhereId  int64
	Phone    string
	FullName string
}
