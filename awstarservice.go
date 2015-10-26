package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"golang.org/x/net/context"
)

type AwsTarService struct {
	// TODO(termie): add a cache
}

// GetFileRange for a tarball uploaded with appropriate metadata
func (s *AwsTarService) GetFile(ctx context.Context, args GetFileArgs) (File, error) {
	resp, err := http.Head(args.TarURL)
	if err != nil {
		return File{}, err
	}

	filename := url.QueryEscape(args.Filename)

	filerange := resp.Header.Get(fmt.Sprintf("x-amz-meta-%s", filename))

	parts := strings.Split(filerange, "-")
	first, _ := strconv.ParseInt(parts[0], 10, 64)
	last, _ := strconv.ParseInt(parts[1], 10, 64)
	response := File{
		Range:       FileRange{first, last},
		TarURL:      args.TarURL,
		Filename:    args.Filename,
		ContentType: args.ContentType,
	}
	return response, nil
}
