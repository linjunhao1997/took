package service

import (
	"context"
	account "took/pkg/account/domain/user"
)

type AccountService interface {
	LoadUser(ctx context.Context, id int) (*account.User, error)
	LoadUsers(ctx context.Context, id ...int) ([]*account.User, error)
}

type service struct {
	userRepo account.Repository
}

func NewAccountService(userRepo account.Repository) AccountService {
	return &service{
		userRepo,
	}
}

func (s *service) LoadUser(ctx context.Context, id int) (*account.User, error) {
	return s.userRepo.FindOne(id)
}

func (s *service) LoadUsers(ctx context.Context, id ...int) ([]*account.User, error) {
	return s.userRepo.Find(id...)
}
