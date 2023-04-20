package server

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/eikoshelev/etcd-proxy-server/internal/proxy"
)

type Server struct {
	httpServer   *http.Server
	Port         string
	RTimeout     time.Duration
	WTimeout     time.Duration
	HostIP       string
	MetricsRoute string
	Proxy        *proxy.Proxy
}

func (s *Server) Setup() {
	mux := http.NewServeMux()
	mux.HandleFunc(s.MetricsRoute, s.Proxy.Handler)
	s.httpServer = &http.Server{
		Addr:         s.Port,
		ReadTimeout:  time.Duration(s.RTimeout) * time.Second,
		WriteTimeout: time.Duration(s.WTimeout) * time.Second,
		Handler:      mux,
	}
}

func (s *Server) Run() {
	if err := s.httpServer.ListenAndServe(); errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error listening for server: %s\n", err)
	}
}
