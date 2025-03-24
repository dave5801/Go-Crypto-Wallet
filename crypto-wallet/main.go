package main

import (
	"crypto-wallet/transaction"
	"crypto-wallet/wallet"
	"fmt"
)

func main() {
	// Generate a new wallet
	myWallet, err := wallet.GenerateWallet()
	if err != nil {
		fmt.Println("Error creating wallet:", err)
		return
	}

	// Save the private key
	err = wallet.SavePrivateKey(myWallet.PrivateKey, "wallet.pem")
	if err != nil {
		fmt.Println("Error saving private key:", err)
		return
	}

	fmt.Println("Wallet created! Public Key:", myWallet.PublicKey)

	// Simulate sending a transaction
	tx := transaction.Transaction{
		From:   myWallet.PublicKey,
		To:     "receiver_public_key",
		Amount: 0.01,
	}

	err = transaction.SendTransaction(tx, myWallet.PrivateKey)
	if err != nil {
		fmt.Println("Transaction failed:", err)
		return
	}
}
