package main

import "github.com/eikoshelev/etcd-proxy-server/internal/config"

func main() {
	logger := logger.GetLogger()
	config := config.Configure()
}
