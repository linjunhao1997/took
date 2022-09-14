package fileconsole

import (
	"context"
	"took/pkg/account"
	"took/pkg/fileconsole/file"
)

type Service interface {
	LoadFile(ctx context.Context, id int) (*file.File, error)
	LoadFiles(ctx context.Context, id ...int) ([]*file.File, error)
	//UploadFile(id int)
}

type service struct {
	fileRepo       file.Repository
	accountService account.Service
}

func NewFileConsoleService(fileRepo file.Repository, accountService account.Service) Service {
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
