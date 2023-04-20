package proxy

import (
	"net/http"

	"github.com/eikoshelev/etcd-proxy-server/internal/client"
)

type Proxy struct {
	Client *client.Client
}

func Setup(client *client.Client) *Proxy {
	return &Proxy{
		Client: client,
	}
}

func (p *Proxy) Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "only GET is allowed", http.StatusMethodNotAllowed)
		return
	}
}
