package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

type Server interface {
	Address() string
	isAlive() bool
	Serve(rw http.ResponseWriter, req *http.Request)
}

func (s *SimpleServer) Address() string {
	return s.addr
}

func (s *SimpleServer) isAlive() bool {
	return true
}

func (s *SimpleServer) Serve(rw http.ResponseWriter, req *http.Request) {
	s.proxy.ServeHTTP(rw, req)
}

type SimpleServer struct {
	addr  string
	proxy *httputil.ReverseProxy
}

func newSimpleServer(addr string) *SimpleServer {
	serverUrl, err := url.Parse(addr)
	handleErr(err)
	return &SimpleServer{
		addr:  addr,
		proxy: httputil.NewSingleHostReverseProxy(serverUrl),
	}
}

type LoadBalancer struct {
	port            string
	roundRobinCount int
	servers         []Server
}

func NewLoadBalancer(port string, servers []Server) *LoadBalancer {
	return &LoadBalancer{
		port:            port,
		roundRobinCount: 0,
		servers:         servers,
	}
}

func (lb *LoadBalancer) getNextAvailServer() Server {
	server := lb.servers[lb.roundRobinCount%len(lb.servers)]
	for !server.isAlive() {
		lb.roundRobinCount++
		server = lb.servers[lb.roundRobinCount%len(lb.servers)]
	}
	lb.roundRobinCount++
	return server
}

func (lb *LoadBalancer) serveProxy(rw http.ResponseWriter, req *http.Request) {
	targetServer := lb.getNextAvailServer()
	fmt.Printf("forwarding reqest to addr %q\n", targetServer.Address())
	targetServer.Serve(rw, req)
}

func handleErr(err error) {
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
}

func main() {
	servers := []Server{
		newSimpleServer("https://fast.com/"),
		newSimpleServer("https://x.com/"),
		newSimpleServer("https://www.instagram.com/"),
	}

	lb := NewLoadBalancer("8080", servers)
	handleRedirect := func(rw http.ResponseWriter, req *http.Request) {
		lb.serveProxy(rw, req)

	}

	http.HandleFunc("/", handleRedirect)
	fmt.Printf("serving requests at localhost:%s\n", lb.port)
	http.ListenAndServe(":"+lb.port, nil)
}
