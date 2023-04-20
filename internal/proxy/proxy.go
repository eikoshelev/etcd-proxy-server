package proxy

import (
	"net/http"

	"github.com/eikoshelev/etcd-proxy-server/internal/client"
)

type Proxy struct {
	Client *client.Client
}

func (p *Proxy) Handler(w http.ResponseWriter, r *http.Request) {}
