package repository

import (
	"crud-api/domain"
)

type UserRepositoryInterface interface {
	CreateUser(user domain.User) error
	GetUser(id int) (domain.User, error)
	DeleteUser(id int) error
	FindByLogin(login string) (domain.User, error)
	Find(login string, password string) (domain.User, error)
}
