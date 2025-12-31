package main

import "testing"

func TestWallet(t *testing.T) {
	// Create new instance of wallet
	wallet := Wallet{}

	//Deposit money into wallet
	wallet.Deposit(Bitcoin(10))

	// Expected and actual
	act := wallet.Balance
	exp := Bitcoin(10)

	if act != exp {
		t.Errorf("Got %d when expecting %d", act, exp)
	}
}
