package credentials

import (
	"crypto/tls"
	"crypto/x509"
	_ "embed"
	"fmt"

	grpccredentials "google.golang.org/grpc/credentials"
)

var (
	//go:embed server.pem
	serverPem []byte
)

func NewTLSCredentials() grpccredentials.TransportCredentials {
	creds, err := newCredentials()
	if err != nil {
		panic(err)
	}
	return creds
}

func newCredentials() (grpccredentials.TransportCredentials, error) {
	cp := x509.NewCertPool()
	if !cp.AppendCertsFromPEM(serverPem) {
		return nil, fmt.Errorf("credentials: failed to append certificates")
	}
	return grpccredentials.NewTLS(&tls.Config{ServerName: "", RootCAs: cp, InsecureSkipVerify: true}), nil
}
