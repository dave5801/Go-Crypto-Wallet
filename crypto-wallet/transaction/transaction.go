package transactions

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
)

// Transaction represents a basic cryptocurrency transaction
type Transaction struct {
	From   string
	To     string
	Amount float64
	Hash   string
}

// NewTransaction creates a new signed transaction
func NewTransaction(from string, to string, amount float64, privKey *ecdsa.PrivateKey) *Transaction {
	tx := &Transaction{From: from, To: to, Amount: amount}
	tx.Hash = tx.signTransaction(privKey)
	return tx
}

// signTransaction hashes the transaction and mocks a digital signature
func (tx *Transaction) signTransaction(privKey *ecdsa.PrivateKey) string {
	data := fmt.Sprintf("%s%s%f%d", tx.From, tx.To, tx.Amount, rand.Intn(1000))
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

