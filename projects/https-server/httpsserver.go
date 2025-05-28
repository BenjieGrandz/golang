package main

import (
	"crypto/rand" // random secure number generator
	"crypto/rsa"  // rsa key generation
	"crypto/tls"	// TLS support in go
	"crypto/x509"	// certificate creation and parsing
	"crypto/x509/pkix" // certificate name fields (subject, issuer, etc.)
	"encoding/pem"	// PEM encoding (text format for keys/certs)
	"fmt"
	"math/big"	// generates serial numbers
	"net/http"
	"time"
)

func main() {

}