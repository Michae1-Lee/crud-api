package usecases

import (
	"crud-api/domain"
)

type UserServiceInterface interface {
	GetUser(id int) (domain.User, error)
	CreateUser(user domain.User) error
	DeleteUser(id int) error
}
