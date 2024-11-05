package usecases

import (
	"crud-api/domain"
)

type UserServiceInterface interface {
	CreateUser(user domain.User) error
	GetUser(id string) (domain.User, error)
	DeleteUser(id string) error
	FindByLogin(login string) (domain.User, error)
	Find(login string, password string) (domain.User, error)
}
