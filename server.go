package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type HTTPServer struct {
	addr  string
	proxy *httputil.ReverseProxy
}

func NewHTTPServer(addr string) *HTTPServer {
	serveURl, err := url.Parse(addr)
	if err != nil {
		log.Fatal(err)
	}
	return &HTTPServer{
		addr:  addr,
		proxy: httputil.NewSingleHostReverseProxy(serveURl),
	}
}

func (s *HTTPServer) Address() string {
	return s.addr
}

func (s *HTTPServer) isAlive() bool {
	return false
}

func (s *HTTPServer) Serve(w http.ResponseWriter, r *http.Request) {
	s.proxy.ServeHTTP(w, r)
}
