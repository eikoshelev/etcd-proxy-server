package main

import (
	"github.com/eikoshelev/etcd-proxy-server/internal/client"
	"github.com/eikoshelev/etcd-proxy-server/internal/config"
	"github.com/eikoshelev/etcd-proxy-server/internal/logger"
	"github.com/eikoshelev/etcd-proxy-server/internal/proxy"
	"github.com/eikoshelev/etcd-proxy-server/internal/server"

	"go.uber.org/zap"
)

func main() {
	logger := logger.GetLogger()
	config, err := config.Configure()
	if err != nil {
		logger.Fatal("configuration failed", zap.String("reason", err.Error()), zap.Any("config", config))
	}
	logger.Info("configuration info", zap.Any("config", config))

	client, err := client.Setup(&config.Client)
	if err != nil {
		logger.Fatal("client configuration failed", zap.String("reason", err.Error()), zap.Any("config", config))
	}
	proxy := proxy.Setup(client, logger)
	server := server.Setup(&config.Server, proxy, logger)
	logger.Info("start proxy server")
	server.Run()
}
