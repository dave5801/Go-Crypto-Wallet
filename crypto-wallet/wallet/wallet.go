package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

type Wallet struct {
	PrivateKey *ecdsa.PrivateKey
	Address    string
}

// NewWallet generates a new wallet with a private key and address
func NewWallet() *Wallet {
	privKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	address := generateAddress(privKey)
	return &Wallet{PrivateKey: privKey, Address: address}
}

// generateAddress creates a wallet address from the public key
func generateAddress(privKey *ecdsa.PrivateKey) string {
	pubKey := append(privKey.PublicKey.X.Bytes(), privKey.PublicKey.Y.Bytes()...)
	hash := sha256.Sum256(pubKey)
	return hex.EncodeToString(hash[:])
}
