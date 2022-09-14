package account

import (
	"context"
	"took/pkg/account/user"
)

type Service interface {
	LoadUser(ctx context.Context, id int) (*user.User, error)
	LoadUsers(ctx context.Context, id ...int) ([]*user.User, error)
}

type service struct {
	userRepo user.Repository
}

func NewAccountService(userRepo user.Repository) Service {
	return &service{
		userRepo,
	}
}

func (s *service) LoadUser(ctx context.Context, id int) (*user.User, error) {
	return s.userRepo.FindOne(id)
}

func (s *service) LoadUsers(ctx context.Context, id ...int) ([]*user.User, error) {
	return s.userRepo.Find(id...)
}
