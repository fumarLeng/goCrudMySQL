package user

import (
	"context"
	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, u *User) error
	List(ctx context.Context) ([]User, error)
	Get(ctx context.Context, id uint) (*User, error)
	Update(ctx context.Context, u *User) error
	Delete(ctx context.Context, id uint) error
}

type repo struct{ db *gorm.DB }

func NewRepository(db *gorm.DB) Repository { return &repo{db} }

func (r *repo) Create(ctx context.Context, u *User) error {
	return r.db.WithContext(ctx).Create(u).Error
}
func (r *repo) List(ctx context.Context) ([]User, error) {
	var us []User
	err := r.db.WithContext(ctx).Find(&us).Error
	return us, err
}
func (r *repo) Get(ctx context.Context, id uint) (*User, error) {
	var u User
	if err := r.db.WithContext(ctx).First(&u, id).Error; err != nil {
		return nil, err
	}
	return &u, nil
}
func (r *repo) Update(ctx context.Context, u *User) error {
	res := r.db.WithContext(ctx).Save(u)
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return res.Error
}
func (r *repo) Delete(ctx context.Context, id uint) error {
	res := r.db.WithContext(ctx).Delete(&User{}, id)
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return res.Error
}
