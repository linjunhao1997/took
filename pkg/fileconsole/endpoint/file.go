package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"took/pkg/fileconsole/domain/file"
	"took/pkg/fileconsole/service"
)

type LoadFileRequest struct {
	Id int
}

type LoadFileResponse struct {
	File *file.File `json:"file,omitempty"`
	Err  error      `json:"error,omitempty"`
}

func MakeLoadFileEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(LoadFileRequest)
		file, err := s.LoadFile(ctx, req.Id)
		return LoadFileResponse{File: file, Err: err}, nil
	}
}
