package client

import (
	"crypto/tls"
	"io"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

type Client struct {
	httpClient    *http.Client
	ClientTimeout time.Duration
	EtcdEndpoint  string
	Certs         certificates
}

func Setup(c *Client) (*Client, error) {
	cert, caCertPool, err := c.Certs.Verify()
	if err != nil {
		return nil, errors.Wrap(err, "certificate verification failed")
	}
	// Setup HTTPS
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCertPool,
	}
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
		Timeout: time.Duration(c.ClientTimeout) * time.Second,
	}
	return &Client{
		httpClient: client,
	}, nil
}

func (c *Client) Do() ([]byte, error) {
	resp, err := c.httpClient.Get(c.EtcdEndpoint)
	if err != nil {
		return nil, errors.Wrap(err, "failed to complete the request")
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read response body")
	}
	return data, nil
}
