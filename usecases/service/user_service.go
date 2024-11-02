package service

import (
	"crud-api/domain"
	"crud-api/repository"
)

type UserService struct {
	userRepo repository.UserRepositoryInterface
}

func NewUserService(userRepo repository.UserRepositoryInterface) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) GetUser(id int) (domain.User, error) {
	user, err := s.userRepo.GetUser(id)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *UserService) CreateUser(user domain.User) error {
	if err := s.userRepo.CreateUser(user); err != nil {
		return err
	}
	return nil
}

func (s *UserService) DeleteUser(id int) error {
	if err := s.userRepo.DeleteUser(id); err != nil {
		return err
	}
	return nil
}

func (s *UserService) FindByLogin(login string) (domain.User, error) {
	user, err := s.userRepo.FindByLogin(login)
	if err != nil {
		return user, err
	}
	return user, nil
}
func (s *UserService) Find(login string, password string) (domain.User, error) {
	user, err := s.userRepo.Find(login, password)
	if err != nil {
		return user, err
	}
	return user, nil
}
