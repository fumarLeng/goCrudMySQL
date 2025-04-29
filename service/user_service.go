package service

import (
	"goCrudMySQL/model"
	"goCrudMySQL/repository"
)

type UserService interface {
	Create(u model.User) (*model.User, error)
	List() ([]model.User, error)
	Get(id uint) (*model.User, error)
	Update(id uint, data model.User) (*model.User, error)
	Delete(id uint) error
}

type userService struct{ repo repository.UserRepository }

func NewUserService(r repository.UserRepository) UserService {
	return &userService{r}
}

func (s *userService) Create(u model.User) (*model.User, error) {
	if err := s.repo.Create(&u); err != nil {
		return nil, err
	}
	return &u, nil
}
func (s *userService) List() ([]model.User, error)      { return s.repo.FindAll() }
func (s *userService) Get(id uint) (*model.User, error) { return s.repo.FindByID(id) }
func (s *userService) Update(id uint, data model.User) (*model.User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	user.Name = data.Name
	user.Email = data.Email
	if err := s.repo.Update(user); err != nil {
		return nil, err
	}
	return user, nil
}
func (s *userService) Delete(id uint) error { return s.repo.Delete(id) }
