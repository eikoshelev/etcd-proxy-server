package config

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/eikoshelev/etcd-proxy-server/internal/client"
	"github.com/eikoshelev/etcd-proxy-server/internal/server"
)

type ProxyConfig struct {
	Server server.Server
	Client client.Client
}

func Configure() (*ProxyConfig, error) {
	var conf ProxyConfig

	serverRT, err := time.ParseDuration("10s")
	if err != nil {
		return nil, err
	}
	serverWR, err := time.ParseDuration("10s")
	if err != nil {
		return nil, err
	}
	clientT, err := time.ParseDuration("10s")
	if err != nil {
		return nil, err
	}

	// server
	flag.StringVar(&conf.Server.Port, "serverPort", ":8888", "Server port")
	flag.DurationVar(&conf.Server.RTimeout, "serverRTimeout", serverRT, "ReadTimeout for server")
	flag.DurationVar(&conf.Server.WTimeout, "serverWTimeout", serverWR, "WriteTimeout for server")
	flag.StringVar(&conf.Server.HostIP, "hostIP", os.Getenv("HOST_IP"), "Host machine IP")
	flag.StringVar(&conf.Server.MetricsRoute, "metricsRoute", "/metrics", "Route for collecting metrics")
	// client
	flag.DurationVar(&conf.Client.ClientTimeout, "clientTimeout", clientT, "Timeout for client")
	// client certs
	flag.StringVar(&conf.Client.Certs.Cert, "certFile", "/etc/kubernetes/pki/etcd/healthcheck-client.crt", "A PEM eoncoded certificate file")
	flag.StringVar(&conf.Client.Certs.Ca, "caFile", "/etc/kubernetes/pki/etcd/ca.crt", "A PEM eoncoded CA's certificate file")
	flag.StringVar(&conf.Client.Certs.Key, "keyFile", "/etc/kubernetes/pki/etcd/healthcheck-client.key", "A PEM encoded private key file")

	flag.Parse()

	conf.Client.EtcdEndpoint = fmt.Sprintf("https://%s:2379/metrics", conf.Server.HostIP)

	return &conf, nil
}
