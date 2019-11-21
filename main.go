package main

import (
	"flag"
	"net/http"
	"os"
)

type network struct {
	Addr           string
	ServerRTimeout int
	ServerWTimeout int
	ClientTimeout  int
	HostIP         string
	Client         *http.Client
	EtcdEndpoint   string
}

func main() {

	var net network

	flag.StringVar(&net.Addr, "addr", ":8888", "Server port")
	flag.IntVar(&net.ServerRTimeout, "serverRTimeout", 3, "ReadTimeout for server")
	flag.IntVar(&net.ServerWTimeout, "serverWTimeout", 3, "WriteTimeout for server")
	flag.IntVar(&net.ClientTimeout, "clientTimeout", 10, "Timeout for client")
	flag.StringVar(&net.HostIP, "hostIP", os.Getenv("HOST_IP"), "Host machine IP")

	cert, caCertPool, ok := certsCheck()
	if ok {
		net.proxy(cert, caCertPool)
	}
}
