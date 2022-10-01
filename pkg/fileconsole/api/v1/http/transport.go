package http

import (
	"context"
	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
	"took/pkg/fileconsole/endpoint"
	"took/pkg/fileconsole/service"
	"took/pkg/util/restful"
)

func MakeHandler(bs service.Service, logger kitlog.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(restful.EncErrResp),
	}

	loadFileHandler := kithttp.NewServer(
		endpoint.MakeLoadFileEndpoint(bs),
		decodeLoadFileRequestByHttp,
		restful.EncodeResp,
		opts...,
	)

	r := mux.NewRouter()

	r.Handle("/booking/v1/files/{id}", loadFileHandler).Methods("GET")

	return r
}

func decodeLoadFileRequestByHttp(_ context.Context, r *http.Request) (interface{}, error) {
	id, err := restful.DecodeQueryId(r)
	if err != nil {
		return nil, err
	}
	return endpoint.LoadFileRequest{Id: id}, nil
}
