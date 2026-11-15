package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func NewReverseProxy(target string) http.Handler {
	targetURL, err := url.Parse(target)
	if err != nil {
		panic(err)
	}

	return httputil.NewSingleHostReverseProxy(targetURL)
}
