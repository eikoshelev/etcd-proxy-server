package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func proxy(cert tls.Certificate, caCert *x509.CertPool, net Network) {

	// Setup HTTPS
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCert,
	}

	tlsConfig.BuildNameToCertificate()
	transport := &http.Transport{TLSClientConfig: tlsConfig}

	log.Printf("HOST_IP env: %s", net.HostIP)

	server := &http.Server{
		Addr:         net.Addr,
		ReadTimeout:  time.Duration(net.ServerRTimeout) * time.Second,
		WriteTimeout: time.Duration(net.ServerWTimeout) * time.Second,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodGet && r.RequestURI == "/metrics" {
				log.Printf("Request: %s %s —> %s%s", r.Host, r.Method, r.RemoteAddr, r.RequestURI)
				client(w, transport, net.ClientTimeout, net.HostIP)
			} else {
				log.Printf("Invalid requests: %s %s —> %s%s", r.Host, r.Method, r.RemoteAddr, r.RequestURI)
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("Invalid requests — only requests of the \"GET\" method and on the url \"/metrics\" are allowed\n"))
			}
		}),
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed start server: %s", err)
	}
}

func client(w http.ResponseWriter, transport *http.Transport, clientTimeout int, hostIP string) {

	// HTTPS client
	client := &http.Client{
		Transport: transport,
		Timeout:   time.Duration(clientTimeout) * time.Second,
	}

	// Do GET metrics
	resp, err := client.Get("https://" + hostIP + ":2379/metrics")
	if err != nil {
		log.Printf("Failed send GET request: %s", err)
		w.Write([]byte("Failed send GET request\n"))
		return
	}

	defer resp.Body.Close()

	// Dump response
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Failed get body from response: %s", err)
		w.Write([]byte("Failed read body from response\n"))
		return
	}

	w.WriteHeader(resp.StatusCode)
	w.Write(data)
}
