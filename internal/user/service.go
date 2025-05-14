package user

import (
	"context"
	"goCrudMySQL/common"
	"strings"
)

type Service interface {
	Create(ctx context.Context, in CreateUserRq) (*User, error)
	List(ctx context.Context) ([]User, error)
	Get(ctx context.Context, id uint) (*User, error)
	Update(ctx context.Context, id uint, in UpdateUserRequest) (*User, error)
	Delete(ctx context.Context, id uint) error
}

type svc struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &svc{r}
}

// Create 建立使用者
func (s *svc) Create(ctx context.Context, in CreateUserRq) (*User, error) {
	u := User{
		Name:  in.Name,
		Email: in.Email,
	}
	if err := s.repo.Create(ctx, &u); err != nil {
		return nil, err
	}
	return &u, nil
}

// List 取得所有使用者
func (s *svc) List(ctx context.Context) ([]User, error) {
	return s.repo.List(ctx)
}

// Get 根據 ID 取得使用者
func (s *svc) Get(ctx context.Context, id uint) (*User, error) {
	u, err := s.repo.Get(ctx, id)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, common.ErrUserNotFound
		}
		return nil, err
	}
	return u, nil
}

// Update 更新使用者資料
func (s *svc) Update(ctx context.Context, id uint, in UpdateUserRequest) (*User, error) {
	u, err := s.repo.Get(ctx, id)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, common.ErrUserNotFound
		}
		return nil, err
	}

	if in.Name != nil {
		u.Name = *in.Name
	}
	if in.Email != nil {
		u.Email = *in.Email
	}
	if err := s.repo.Update(ctx, u); err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return nil, common.ErrEmailAlreadyExists
		}
		return nil, err
	}
	return u, nil
}

// Delete 刪除使用者
func (s *svc) Delete(ctx context.Context, id uint) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return common.ErrUserNotFound
		}
		return err
	}
	return nil
}
