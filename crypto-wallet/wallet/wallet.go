package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"os"
)

// Wallet holds the private and public key pair
type Wallet struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  string
}

// GenerateWallet creates a new wallet with ECDSA key pair
func GenerateWallet() (*Wallet, error) {
	privKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}

	// Convert public key to a readable format
	pubKeyBytes := append(privKey.PublicKey.X.Bytes(), privKey.PublicKey.Y.Bytes()...)
	pubKeyHex := hex.EncodeToString(pubKeyBytes)

	return &Wallet{
		PrivateKey: privKey,
		PublicKey:  pubKeyHex,
	}, nil
}

// SavePrivateKey stores the private key in a file
func SavePrivateKey(priv *ecdsa.PrivateKey, filename string) error {
	keyBytes, err := x509.MarshalECPrivateKey(priv)
	if err != nil {
		return err
	}

	pemBlock := &pem.Block{Type: "EC PRIVATE KEY", Bytes: keyBytes}
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return pem.Encode(file, pemBlock)
}
