package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"os"
	"tls-example/tlsutil"
)

var (
	certFile   = "./ca/local.client/cert.pem"
	keyFile    = "./ca/local.client/key.pem"
	caCertFile = "./ca/minica.pem"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Need a server URL")
		os.Exit(1)
	}

	srvr := os.Args[1]

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: getTlsConf(),
		},
	}

	resp, err := client.Get(srvr)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}

func getTlsConf() *tls.Config {

	return &tls.Config{
		RootCAs:               tlsutil.CAPool(caCertFile),
		GetClientCertificate:  tlsutil.GetClientCert(certFile, keyFile),
		VerifyPeerCertificate: tlsutil.InspectChainFunc,
	}
}
