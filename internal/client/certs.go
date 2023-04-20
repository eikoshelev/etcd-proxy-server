package client

import (
	"crypto/tls"
	"crypto/x509"
	"os"
)

type certificates struct {
	Cert string
	Ca   string
	Key  string
}

func (certs *certificates) Verify() (tls.Certificate, *x509.CertPool, error) {
	// load client cert
	cert, err := tls.LoadX509KeyPair(certs.Cert, certs.Key)
	if err != nil {
		return tls.Certificate{}, nil, err
	}
	// load CA cert
	caCert, err := os.ReadFile(certs.Ca)
	if err != nil {
		return tls.Certificate{}, nil, err
	}
	caCertPool := x509.NewCertPool()
	if ok := caCertPool.AppendCertsFromPEM(caCert); ok {
		return cert, caCertPool, nil
	}
	return tls.Certificate{}, nil, err
}
