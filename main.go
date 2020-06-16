package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"golang.org/x/crypto/ssh"
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
	rsaKey, _ := rsa.GenerateKey(rand.Reader, configBits)

	// id_rsa
	var privBuffer strings.Builder
	_ = pem.Encode(&privBuffer, &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(rsaKey),
	})
	print(privBuffer.String())

	// id_rsa.pub
	pubRsaKey, _ := ssh.NewPublicKey(&rsaKey.PublicKey)
	print(string(ssh.MarshalAuthorizedKey(pubRsaKey)))
}
