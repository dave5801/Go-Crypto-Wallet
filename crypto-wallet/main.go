package main

import (
	"fmt"
	"log"

	"crypto-wallet/wallet"
)

func main() {
	// Create a new wallet
	w := wallet.NewWallet()

	// Display wallet details
	fmt.Println("New Wallet Address:", w.Address)

	// Save private key
	err := wallet.SaveWallet(w, "wallet.json")
	if err != nil {
		log.Fatal("Error saving wallet:", err)
	}

	// Load the wallet
	loadedWallet, err := wallet.LoadWallet("wallet.json")
	if err != nil {
		log.Fatal("Error loading wallet:", err)
	}

	fmt.Println("Loaded Wallet Address:", loadedWallet.Address)
}
