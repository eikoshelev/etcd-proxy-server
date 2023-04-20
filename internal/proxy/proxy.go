package proxy

import (
	"net/http"

	"github.com/eikoshelev/etcd-proxy-server/internal/client"
	"go.uber.org/zap"
)

type Proxy struct {
	Client *client.Client
	Logger *zap.Logger
}

func Setup(client *client.Client, logger *zap.Logger) *Proxy {
	return &Proxy{
		Client: client,
		Logger: logger,
	}
}

func (p *Proxy) Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		p.Logger.Error("invalid request: only GET is allowed", zap.String("request method", r.Method), zap.String("from host", r.Host))
		http.Error(w, "invalid request: only GET is allowed", http.StatusMethodNotAllowed)
		return
	}
	resp, err := p.Client.GetMetrics()
	if err != nil {
		http.Error(w, "only GET is allowed", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
