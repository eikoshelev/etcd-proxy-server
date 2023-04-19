package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func (net *network) proxy(cert tls.Certificate, caCert *x509.CertPool) {

	// Setup HTTPS
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCert,
	}

	tlsConfig.BuildNameToCertificate()

	net.Client = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
		Timeout: time.Duration(net.ClientTimeout) * time.Second,
	}

	net.EtcdEndpoint = fmt.Sprintf("https://%s:2379/metrics", net.HostIP)

	log.Printf("HOST_IP env: %s", net.HostIP)

	server := &http.Server{
		Addr:         net.Addr,
		ReadTimeout:  time.Duration(net.ServerRTimeout) * time.Second,
		WriteTimeout: time.Duration(net.ServerWTimeout) * time.Second,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodGet && r.RequestURI == "/metrics" {
				log.Printf("Request: %s %s —> %s%s", r.Host, r.Method, r.RemoteAddr, r.RequestURI)
				net.client(w)
			} else {
				log.Printf("Invalid requests: %s %s —> %s%s", r.Host, r.Method, r.RemoteAddr, r.RequestURI)
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("Invalid requests — only requests of the 'GET' method and on the url '/metrics' are allowed\n"))
			}
		}),
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed start server: %s", err.Error())
	}
}

func (net *network) client(w http.ResponseWriter) {

	// do GET metrics
	resp, err := net.Client.Get(net.EtcdEndpoint)
	if err != nil {
		log.Printf("Failed send GET request: %s", err.Error())
		w.Write([]byte("Failed send GET request\n"))
		return
	}

	defer resp.Body.Close()

	// dump response
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Failed get body from response: %s", err.Error())
		w.Write([]byte("Failed read body from response\n"))
		return
	}

	w.WriteHeader(resp.StatusCode)
	w.Write(data)
}
