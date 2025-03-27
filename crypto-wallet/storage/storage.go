package storage

import "sync"

// Fake in-memory storage for user balances
var balances = map[string]float64{}
var mutex = sync.Mutex{}

// UpdateBalance updates the balance of a wallet
func UpdateBalance(address string, amount float64) {
	mutex.Lock()
	defer mutex.Unlock()
	balances[address] += amount
}

// GetBalance returns the balance of a given address
func GetBalance(address string) float64 {
	mutex.Lock()
	defer mutex.Unlock()
	return balances[address]
}


