package main

import (
	"crypto/rand"      // random secure number generator
	"crypto/rsa"       // rsa key generation
	"crypto/tls"       // TLS support in go
	"crypto/x509"      // certificate creation and parsing
	"crypto/x509/pkix" // certificate name fields (subject, issuer, etc.)
	"encoding/pem"     // PEM encoding (text format for keys/certs)
	"fmt"
	"math/big" // generates serial numbers
	"net/http"
	"time"
)

// generates selfsigned cert in memory and returns tls.cert
func generateSelfSignedCert() (tls.Certificate, error) {
	// Generate 2048-bit RSA private key
	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return tls.Certificate{}, err
	}

	// Define Cert Properties
	template := x509.Certificate {
		SerialNumber: big.NewInt(1), // unique identifier for this cert
		Subject: pkix.Name{			// who the cert is for
			CommonName: "localhost",
			Organization: []string{"MyDevCert"}, //Optional
		},
		NotBefore: time.Now(),		// valid for now
		NotAfter: time.Now().Add(365 * 24 * time.Hour), // valid for one year


		// permissions this cert allows
		KeyUsage: x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth}, //use as server cert

		BasicConstraintsValid: true, // required in modern TLS
	}
	// create a s.s certificate (template signs itself)
	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)
	if err != nil {
		return tls.Certificate{}, err
	}

	// convert cert to PEM
	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: derBytes})

	// convert private key to PEM
	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(priv)})

	//combine cert + key to make tls.Certificate
	return tls.X509KeyPair(certPEM, keyPEM)
}

// helloHandler end point
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "🔒 Hello, Https world from Go!")
}

func main() {
	// Generate a self-signed TLS cetificate 
	cert, err := generateSelfSignedCert()
	if err != nil {
		panic("Failed to generate certificate: " + err.Error())
	}

	// configure the TLS settings
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	// create a http server that uses TLS
	server := &http.Server{
		Addr: ":8443",
		Handler: http.HandlerFunc(helloHandler),
		TLSConfig: tlsConfig,
	}

	fmt.Println("Starting HTTPS server on https://localshost:8443")

	// start serving the HTTPs - certs are in-memory so pass empty paths
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}