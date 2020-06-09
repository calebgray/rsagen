package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/asn1"
	"encoding/pem"
	"os"
	"strconv"
	"strings"
)

func main() {
	// do.do
	configBits := 2048
	switch len(os.Args) {
	case 2:
		var err error
		if configBits, err = strconv.Atoi(os.Args[1]); err != nil {
			configBits = 2048
		}
	}

	// generate keys
	keys, _ := rsa.GenerateKey(rand.Reader, configBits)

	// id_rsa
	var privBuffer strings.Builder
	_ = pem.Encode(&privBuffer, &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(keys),
	})
	print(privBuffer.String())

	// id_rsa.pub
	var pubBuffer strings.Builder
	asn1Bytes, _ := asn1.Marshal(keys.PublicKey)
	_ = pem.Encode(&pubBuffer, &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: asn1Bytes,
	})
	print(pubBuffer.String())
}
