package client

import (
	"crypto/tls"
	"io"
	"log"
	"net/http"
	"time"
)

type Client struct {
	httpClient    *http.Client
	ClientTimeout time.Duration
	EtcdEndpoint  string
	Certs         certificates
}

func (c *Client) New() *Client {
	cert, caCertPool, ok := c.Certs.Check()
	if !ok {
		log.Fatal("pizdec")
	}
	// Setup HTTPS
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCertPool,
	}

	//tlsConfig.BuildNameToCertificate()

	c.httpClient = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
		Timeout: time.Duration(c.ClientTimeout) * time.Second,
	}

	return &Client{}
}

func (c *Client) Do() ([]byte, error) {
	resp, err := c.httpClient.Get(c.EtcdEndpoint)
	if err != nil {
		log.Printf("Failed send GET request: %s", err.Error())
		return nil, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Failed get body from response: %s", err.Error())
		return nil, err
	}
	return data, nil
}
