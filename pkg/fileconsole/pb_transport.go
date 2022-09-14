package fileconsole

import (
	"context"
	"fmt"
	"took/pkg/account/util"
	"took/pkg/fileconsole/fileconsolepb"
)

func decodeLoadFileRequestByGrpc(_ context.Context, req interface{}) (interface{}, error) {
	pb, ok := req.(*fileconsolepb.LoadFileRequest)
	if !ok {
		return nil, fmt.Errorf("grpc server decode request error")
	}
	request := loadFileRequest{
		Id: int(pb.Id),
	}
	return request, nil
}

func encodeLoadFileResponseByGrpc(_ context.Context, response interface{}) (interface{}, error) {
	data, ok := response.(loadFileResponse)
	if !ok {
		return nil, fmt.Errorf("grpc server encode response error (%T)", data)
	}

	if data.Err != nil {
		return nil, data.Err
	}

	resp := &fileconsolepb.FileResponse{
		Data: &fileconsolepb.File{
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
