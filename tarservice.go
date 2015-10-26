package main

import "golang.org/x/net/context"

// TarService proxies range requests into tarballs to serve files directly
type TarService interface {
	GetFile(ctx context.Context, args GetFileArgs) (File, error)
}

type FileRange struct {
	First int64
	Last  int64
}

type File struct {
	Range       FileRange
	TarURL      string
	Filename    string
	ContentType string
}

type GetFileArgs struct {
	TarURL      string `json:"tarUrl"`
	Filename    string `json:"filename"`
	ContentType string `json:"contentType"`
}
