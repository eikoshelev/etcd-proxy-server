package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"io/ioutil"
	"log"
)

// Certificates - struct for cert, ca and key files
type Certificates struct {
	Cert string
	Ca   string
	Key  string
}

func certsCheck() (tls.Certificate, *x509.CertPool, error) {

	var certs Certificates

	flag.StringVar(&certs.Cert, "certFile", "/etc/kubernetes/pki/etcd/healthcheck-client.crt", "A PEM eoncoded certificate file")
	flag.StringVar(&certs.Ca, "caFile", "/etc/kubernetes/pki/etcd/ca.crt", "A PEM eoncoded CA's certificate file")
	flag.StringVar(&certs.Key, "keyFile", "/etc/kubernetes/pki/etcd/healthcheck-client.key", "A PEM encoded private key file")

	flag.Parse()

	log.Printf("Checking cert files..\n")

	// Load client cert
	cert, err := tls.LoadX509KeyPair(certs.Cert, certs.Key)
	if err != nil {
		log.Fatalf("Failed read 'Cert' or/and 'Key' file(s): %s", err)
		return tls.Certificate{}, nil, err
	} else {
		log.Printf("%s — OK\n", certs.Cert)
		log.Printf("%s — OK\n", certs.Ca)
	}

	// Load CA cert
	caCert, err := ioutil.ReadFile(certs.Ca)
	if err != nil {
		log.Fatalf("Failed read 'Ca' file: %s", err)
		return tls.Certificate{}, nil, err
	} else {
		log.Printf("%s — OK\n", certs.Key)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	return cert, caCertPool, nil
}
