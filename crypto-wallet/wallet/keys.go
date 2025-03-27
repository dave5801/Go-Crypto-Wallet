package wallet

import (
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"os"
)

// WalletFile represents a struct for storing wallet data
type WalletFile struct {
	PrivateKey string `json:"private_key"`
	Address    string `json:"address"`
}

// SaveWallet stores the private key and address in a JSON file
func SaveWallet(w *Wallet, filename string) error {
	privKeyBytes, err := x509.MarshalECPrivateKey(w.PrivateKey)
	if err != nil {
		return err
	}

	privKeyPEM := pem.EncodeToMemory(&pem.Block{Type: "EC PRIVATE KEY", Bytes: privKeyBytes})
	walletFile := WalletFile{PrivateKey: string(privKeyPEM), Address: w.Address}

	data, err := json.Marshal(walletFile)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0600)
}

// LoadWallet loads the private key and address from a JSON file
func LoadWallet(filename string) (*Wallet, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var walletFile WalletFile
	err = json.Unmarshal(data, &walletFile)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode([]byte(walletFile.PrivateKey))
	if block == nil {
		return nil, errors.New("failed to decode PEM block")
	}

	privKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return &Wallet{PrivateKey: privKey, Address: walletFile.Address}, nil
}
