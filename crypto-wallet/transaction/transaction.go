package transaction

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
)

// Transaction represents a simple crypto transaction
type Transaction struct {
	From   string
	To     string
	Amount float64
}

// SignTransaction signs a transaction with a private key
func SignTransaction(tx Transaction, privKey *ecdsa.PrivateKey) (string, error) {
	txData := fmt.Sprintf("%s%s%f", tx.From, tx.To, tx.Amount)
	hash := sha256.Sum256([]byte(txData))

	r, s, err := ecdsa.Sign(rand.Reader, privKey, hash[:])
	if err != nil {
		return "", err
	}

	signature := append(r.Bytes(), s.Bytes()...)
	return hex.EncodeToString(signature), nil
}

// SendTransaction submits a transaction to an API (e.g., blockchain network)
func SendTransaction(tx Transaction, privKey *ecdsa.PrivateKey) error {
	signature, err := SignTransaction(tx, privKey)
	if err != nil {
		return err
	}

	txData := fmt.Sprintf(`{"from": "%s", "to": "%s", "amount": %f, "signature": "%s"}`,
		tx.From, tx.To, tx.Amount, signature)

	resp, err := http.Post("https://api.crypto-network/send", "application/json", bytes.NewBuffer([]byte(txData)))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Println("Transaction sent! Status:", resp.Status)
	return nil
}
