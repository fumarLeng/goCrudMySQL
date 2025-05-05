package user

import "context"

type Service interface {
	Create(ctx context.Context, in User) (*User, error)
	List(ctx context.Context) ([]User, error)
	Get(ctx context.Context, id uint) (*User, error)
	Update(ctx context.Context, id uint, in User) (*User, error)
	Delete(ctx context.Context, id uint) error
}

type svc struct{ repo Repository }

func NewService(r Repository) Service { return &svc{r} }

func (s *svc) Create(ctx context.Context, in User) (*User, error) {
	if err := s.repo.Create(ctx, &in); err != nil {
		return nil, err
	}
	return &in, nil
}
func (s *svc) List(ctx context.Context) ([]User, error)        { return s.repo.List(ctx) }
func (s *svc) Get(ctx context.Context, id uint) (*User, error) { return s.repo.Get(ctx, id) }
func (s *svc) Update(ctx context.Context, id uint, in User) (*User, error) {
	u, err := s.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	u.Name, u.Email = in.Name, in.Email
	return u, s.repo.Update(ctx, u)
}
func (s *svc) Delete(ctx context.Context, id uint) error { return s.repo.Delete(ctx, id) }
