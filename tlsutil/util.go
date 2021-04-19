package tlsutil

import (
	"bufio"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
)

func InspectChainFunc(rawCerts [][]byte, verifiedChains [][]*x509.Certificate) error {

	if len(verifiedChains) > 0 {
		fmt.Println("Valid certificate chain")

		for _, chain := range verifiedChains {
			for i, cert := range chain {
				fmt.Printf("Certificate %d:\n", i)
				fmt.Println(dumpCertInfo(cert))
			}
		}
	}

	return nil
}

func dumpCertInfo(cert *x509.Certificate) string {

	if cert.Subject.CommonName == cert.Issuer.CommonName {
		return fmt.Sprintf("Self-signed certificate:  %v\n", cert.Issuer.CommonName)
	}

	s := fmt.Sprintf("   Subject:   %v\n", cert.DNSNames)
	s += fmt.Sprintf("   Issuer:    %v\n", cert.Issuer.CommonName)

	return s
}

func OutputPEMFile(pf string) error {

	data, err := ioutil.ReadFile(pf)
	if err != nil {
		return err
	}

	for len(data) > 0 {
		var blk *pem.Block
		blk, data = pem.Decode(data)

		fmt.Printf("Type: %v\n", blk.Type)

		switch blk.Type {
		case "CERTIFICATE":
			cert, err := x509.ParseCertificate(blk.Bytes)
			if err != nil {
				return err
			}
			fmt.Println(dumpCertInfo(cert))
		}

	}
	return nil
}

func wait() {
	fmt.Println("[Press ENTER to continue]")
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanRunes)
	s.Scan()
	fmt.Println()
}
