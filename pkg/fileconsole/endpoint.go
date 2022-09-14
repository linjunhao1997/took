package fileconsole

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/sirupsen/logrus"
	"took/pkg/fileconsole/file"
)

type loadFileRequest struct {
	Id int
}

type loadFileResponse struct {
	File *file.File `json:"file,omitempty"`
	Err  error      `json:"error,omitempty"`
}

func makeLoadFileEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		logrus.Info("收到请求")
		req := request.(loadFileRequest)
		file, err := s.LoadFile(ctx, req.Id)
		return loadFileResponse{File: file, Err: err}, nil
	}
}
