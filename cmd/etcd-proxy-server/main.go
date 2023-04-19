package main

import (
	"flag"
	"net/http"
	"os"
	"time"
)

type network struct {
	Addr           string
	ServerRTimeout time.Duration
	ServerWTimeout time.Duration
	ClientTimeout  time.Duration
	HostIP         string
	Client         *http.Client
	EtcdEndpoint   string
}

func main() {

	var net network

	flag.StringVar(&net.Addr, "addr", ":8888", "Server port")
	flag.DurationVar(&net.ServerRTimeout, "serverRTimeout", 10, "ReadTimeout for server")
	flag.DurationVar(&net.ServerWTimeout, "serverWTimeout", 10, "WriteTimeout for server")
	flag.DurationVar(&net.ClientTimeout, "clientTimeout", 10, "Timeout for client")
	flag.StringVar(&net.HostIP, "hostIP", os.Getenv("HOST_IP"), "Host machine IP")

	cert, caCertPool, ok := certsCheck()
	if ok {
		net.proxy(cert, caCertPool)
	}
}
