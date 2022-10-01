package service

import (
	"context"
	accountService "took/pkg/account/service"
	"took/pkg/fileconsole/domain/file"
)

type Service interface {
	LoadFile(ctx context.Context, id int) (*file.File, error)
	LoadFiles(ctx context.Context, id ...int) ([]*file.File, error)
	//UploadFile(id int)
}

type service struct {
	fileRepo       file.Repository
	accountService accountService.AccountService
}

func NewFileConsoleService(fileRepo file.Repository, accountService accountService.AccountService) Service {
	return &service{
		fileRepo,
		accountService,
	}
}

func (s *service) LoadFile(ctx context.Context, id int) (*file.File, error) {
	file, err := s.fileRepo.FindOne(id)
	if err != nil {
		return nil, err
	}
	user, err := s.accountService.LoadUser(ctx, file.CreatorId)
	if err != nil {
		return nil, err
	}
	file.Creator = user
	return file, nil
}

func (s *service) LoadFiles(ctx context.Context, id ...int) ([]*file.File, error) {
	return s.fileRepo.Find(id...)
}
