package user

import (
	"context"
	"gorm.io/gorm"
)

// Repository 定義 User 資料存取介面
type Repository interface {
	Create(ctx context.Context, u *User) error
	List(ctx context.Context) ([]User, error)
	Get(ctx context.Context, id uint) (*User, error)
	Update(ctx context.Context, u *User) error
	Delete(ctx context.Context, id uint) error
}

// repo 是實作 Repository 接口的結構體（相當於實作類別）
type repo struct{ db *gorm.DB }

// NewRepository 是 constructor，回傳 Repository 實作
func NewRepository(db *gorm.DB) Repository {
	return &repo{db}
}

// Create 新增一筆使用者資料
func (r *repo) Create(ctx context.Context, u *User) error {
	return r.db.WithContext(ctx).Create(u).Error
}

// List 取得所有使用者資料
func (r *repo) List(ctx context.Context) ([]User, error) {
	var us []User
	err := r.db.WithContext(ctx).Find(&us).Error
	return us, err
}

// Get 根據 ID 取得單一使用者
func (r *repo) Get(ctx context.Context, id uint) (*User, error) {
	var u User
	if err := r.db.WithContext(ctx).First(&u, id).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

// Update 更新使用者資料（整筆覆蓋）
func (r *repo) Update(ctx context.Context, u *User) error {
	res := r.db.WithContext(ctx).Save(u)
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return res.Error
}

// Delete 根據 ID 刪除使用者
func (r *repo) Delete(ctx context.Context, id uint) error {
	res := r.db.WithContext(ctx).Delete(&User{}, id)
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return res.Error
}
