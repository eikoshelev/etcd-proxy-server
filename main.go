package main

import (
	"flag"
	"log"
	"os"
)

// Network — server and client settings
type Network struct {
	Addr           string
	ServerRTimeout int
	ServerWTimeout int
	ClientTimeout  int
	HostIP         string
}

func main() {

	var net Network

	flag.StringVar(&net.Addr, "addr", ":8888", "Server port")
	flag.IntVar(&net.ServerRTimeout, "serverRTimeout", 10, "ReadTimeout for server")
	flag.IntVar(&net.ServerWTimeout, "serverWTimeout", 10, "WriteTimeout for server")
	flag.IntVar(&net.ClientTimeout, "clientTimeout", 10, "Timeout for client")
	flag.StringVar(&net.HostIP, "hostIP", os.Getenv("HOST_IP"), "Host machine IP")

	cert, caCertPool, err := certsCheck()
	if err != nil {
		log.Fatalf("Failed check certificates: %s", err)
	} else {
		proxy(cert, caCertPool, net)
	}
}
