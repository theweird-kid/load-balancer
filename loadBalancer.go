package main

import (
	"net/http"
)

var N = 10

type Server interface {
	Address() string
	isAlive() bool
	Serve(http.ResponseWriter, *http.Request)
}

type LoadBalancer struct {
	port            string
	roundRobinCount int32
	servers         []Server
}

func NewLoadBalancer(port string, rrCnt int32, servers []Server) *LoadBalancer {
	return &LoadBalancer{
		port:            port,
		roundRobinCount: rrCnt,
		servers:         servers,
	}
}
