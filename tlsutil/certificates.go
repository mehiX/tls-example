package tlsutil

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
)

func GetCertificate(cert, key string) func(*tls.ClientHelloInfo) (*tls.Certificate, error) {

	c, err := tls.LoadX509KeyPair(cert, key)

	return func(hello *tls.ClientHelloInfo) (*tls.Certificate, error) {
		fmt.Printf("Client is asking for server: %v\n", hello.ServerName)

		if err != nil {
			fmt.Println("Don't have a certificate")
		} else {
			err := OutputPEMFile(cert)
			if err != nil {
				fmt.Println(err)
			}
		}

		wait()

		return &c, err
	}
}

func GetClientCert(cert, key string) func(*tls.CertificateRequestInfo) (*tls.Certificate, error) {

	c, err := tls.LoadX509KeyPair(cert, key)

	return func(*tls.CertificateRequestInfo) (*tls.Certificate, error) {
		fmt.Println("Server is asking for a certificate")

		if err != nil {
			fmt.Println("Don't have a certificate")
		} else {
			err := OutputPEMFile(cert)
			if err != nil {
				fmt.Println(err)
			}
		}

		wait()

		return &c, err
	}

}

func CAPool(cert string) *x509.CertPool {
	d, err := ioutil.ReadFile(cert)
	if err != nil {
		fmt.Println("No root CA certificate")
		return nil
	}
	p := x509.NewCertPool()
	p.AppendCertsFromPEM(d)
	return p
}
