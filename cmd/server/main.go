package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	"tls-example/tlsutil"
)

var (
	certFile   = "./ca/local.host/cert.pem"
	keyFile    = "./ca/local.host/key.pem"
	caCertFile = "./ca/minica.pem"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("need listen address")
		os.Exit(1)
	}

	addr := os.Args[1]

	server := &http.Server{
		Addr:      addr,
		Handler:   getHandler(),
		TLSConfig: getTlsConf(certFile, keyFile),
	}

	if err := server.ListenAndServeTLS("", ""); err != nil {
		log.Println(err)
	}

}

func getHandler() http.Handler {
	m := http.NewServeMux()
	m.HandleFunc("/", func(wr http.ResponseWriter, r *http.Request) {
		wr.Write([]byte("Ciao!!"))
	})
	return m
}

func getTlsConf(cert, key string) *tls.Config {

	return &tls.Config{
		ClientCAs:             tlsutil.CAPool(caCertFile),
		ClientAuth:            tls.RequireAndVerifyClientCert,
		GetCertificate:        tlsutil.GetCertificate(cert, key),
		VerifyPeerCertificate: tlsutil.InspectChainFunc,
	}
}
