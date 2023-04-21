package server

import (
	"net/http"
	"time"

	"github.com/eikoshelev/etcd-proxy-server/internal/proxy"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type Server struct {
	httpServer   *http.Server
	Port         string
	RTimeout     time.Duration
	WTimeout     time.Duration
	HostIP       string
	MetricsRoute string
	Proxy        *proxy.Proxy
	Logger       *zap.Logger
}

func Setup(s *Server, proxy *proxy.Proxy, logger *zap.Logger) *Server {
	mux := http.NewServeMux()
	mux.HandleFunc(s.MetricsRoute, s.Proxy.Handler)
	server := &http.Server{
		Addr:         s.Port,
		ReadTimeout:  time.Duration(s.RTimeout) * time.Second,
		WriteTimeout: time.Duration(s.WTimeout) * time.Second,
		Handler:      mux,
	}
	return &Server{
		httpServer: server,
		Proxy:      proxy,
		Logger:     logger,
	}
}

func (s *Server) Run() {
	if err := s.httpServer.ListenAndServe(); errors.Is(err, http.ErrServerClosed) {
		s.Logger.Fatal("server closed", zap.String("error", err.Error()))
	} else if err != nil {
		s.Logger.Fatal("error listening for server", zap.String("reason", err.Error()))
	}
}
