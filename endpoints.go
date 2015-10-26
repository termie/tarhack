package main

import (
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"
)

func CreateGetFileEndpoint(svc TarService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		args, ok := request.(GetFileArgs)
		if !ok {
			return nil, endpoint.ErrBadCast
		}
		file, err := svc.GetFile(ctx, args)
		return file, err
	}
}
