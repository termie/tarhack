package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"golang.org/x/net/context"
)

func CreateGetFileBinding(ctx context.Context, e endpoint.Endpoint) http.Handler {
	decode := func(r *http.Request) (interface{}, error) {
		q := r.URL.Query()
		request := GetFileArgs{
			TarURL:      q.Get("tarUrl"),
			Filename:    q.Get("filename"),
			ContentType: q.Get("contentType"),
		}
		return request, nil
	}
	encode := func(w http.ResponseWriter, response interface{}) error {
		resp, ok := response.(File)
		if !ok {
			return endpoint.ErrBadCast
		}

		// Stream the request

		proxyreq, err := http.NewRequest("GET", resp.TarURL, nil)
		if err != nil {
			return err
		}
		proxyreq.Header.Add("Range", fmt.Sprintf("bytes:%d-%d", resp.Range.First, resp.Range.Last))

		proxyresp, err := http.DefaultClient.Do(proxyreq)
		if err != nil {
			return err
		}

		w.Header().Add("Content-Type", resp.ContentType)
		w.WriteHeader(200)

		defer proxyresp.Body.Close()

		io.Copy(w, proxyresp.Body)
		return nil
	}
	return httptransport.NewServer(ctx, e, decode, encode)
}
