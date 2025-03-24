package wallet

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"os"
)

// LoadPrivateKey loads a private key from a file
func LoadPrivateKey(filename string) (*ecdsa.PrivateKey, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(file)
	if block == nil {
		return nil, err
	}

	privKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return privKey, nil
}
