package server

import (
	"crypto-wallet-api/storage"
	"crypto-wallet-api/transactions"
	"crypto-wallet-api/wallet"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var globalWallet *wallet.Wallet

func StartServer() {
	r := gin.Default()

	// Wallet Endpoints
	r.POST("/wallet/new", createWallet)
	r.GET("/wallet/balance", getWalletBalance)

	// Transaction Endpoint
	r.POST("/transaction", createTransaction)

	// Start Server
	fmt.Println("Starting API on port 8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server")
	}
}

// Create a new wallet
func createWallet(c *gin.Context) {
	globalWallet = wallet.NewWallet()
	err := wallet.SaveWallet(globalWallet, "wallet.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save wallet"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"address": globalWallet.Address})
}

// Get wallet balance
func getWalletBalance(c *gin.Context) {
	if globalWallet == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No wallet found"})
		return
	}
	balance := storage.GetBalance(globalWallet.Address)
	c.JSON(http.StatusOK, gin.H{"address": globalWallet.Address, "balance": balance})
}

// Create a transaction
func createTransaction(c *gin.Context) {
	var txRequest struct {
		To     string  `json:"to"`
		Amount float64 `json:"amount"`
	}

	if err := c.BindJSON(&txRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if globalWallet == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No wallet found"})
		return
	}

	tx := transactions.NewTransaction(globalWallet.Address, txRequest.To, txRequest.Amount, globalWallet.PrivateKey)
	storage.UpdateBalance(txRequest.To, txRequest.Amount)
	c.JSON(http.StatusOK, gin.H{"message": "Transaction successful", "tx_hash": tx.Hash})
}
