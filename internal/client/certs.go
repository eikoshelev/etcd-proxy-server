package client

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"os"
)

type certificates struct {
	Cert string
	Ca   string
	Key  string
}

func (certs *certificates) Check() (tls.Certificate, *x509.CertPool, bool) {
	log.Printf("Checking cert files..\n")

	// load client cert
	cert, err := tls.LoadX509KeyPair(certs.Cert, certs.Key)
	if err != nil {
		log.Fatalf("Failed read 'Cert' or/and 'Key' file(s): %s\n", err)
	}

	log.Printf("%s — OK\n", certs.Cert)
	log.Printf("%s — OK\n", certs.Ca)

	// load CA cert
	caCert, err := os.ReadFile(certs.Ca)
	if err != nil {
		log.Fatalf("Failed read 'Ca' file: %s", err)
	}

	log.Printf("%s — OK\n", certs.Key)

	caCertPool := x509.NewCertPool()

	if ok := caCertPool.AppendCertsFromPEM(caCert); ok {
		return cert, caCertPool, ok
	}

	return tls.Certificate{}, nil, false
}
