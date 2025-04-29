package repository

import (
	"goCrudMySQL/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *model.User) error
	FindAll() ([]model.User, error)
	FindByID(id uint) (*model.User, error)
	Update(user *model.User) error
	Delete(id uint) error
}

type userRepository struct{ db *gorm.DB }

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(u *model.User) error { return r.db.Create(u).Error }
func (r *userRepository) FindAll() ([]model.User, error) {
	var users []model.User
	err := r.db.Find(&users).Error
	return users, err
}
func (r *userRepository) FindByID(id uint) (*model.User, error) {
	var user model.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *userRepository) Update(u *model.User) error { return r.db.Save(u).Error }
func (r *userRepository) Delete(id uint) error       { return r.db.Delete(&model.User{}, id).Error }
