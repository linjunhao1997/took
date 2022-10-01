package grpc

import (
	"context"
	"fmt"
	"took/pkg/account/util"
	fileconsole "took/pkg/fileconsole/api/v1/grpc/proto"
	"took/pkg/fileconsole/endpoint"
)

func decodeLoadFileRequestByGrpc(_ context.Context, req interface{}) (interface{}, error) {
	pb, ok := req.(*fileconsole.LoadFileRequest)
	if !ok {
		return nil, fmt.Errorf("grpc server decode request error")
	}
	request := endpoint.LoadFileRequest{
		Id: int(pb.Id),
	}
	return request, nil
}

func encodeLoadFileResponseByGrpc(_ context.Context, response interface{}) (interface{}, error) {
	data, ok := response.(endpoint.LoadFileResponse)
	if !ok {
		return nil, fmt.Errorf("grpc server encode response error (%T)", data)
	}

	if data.Err != nil {
		return nil, data.Err
	}

	resp := &fileconsole.FileResponse{
		Data: &fileconsole.File{
			Id:      int32(data.File.Id),
			Name:    data.File.Name,
			Size:    int32(data.File.Size),
			Type:    data.File.Type,
			Bucket:  data.File.Bucket,
			Creator: util.ConvertFromUser(data.File.Creator),
		},
	}
	return resp, nil
}
