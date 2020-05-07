package main

import "testing"

func TestWallet(t *testing.T) {
	myWallet := Wallet{}

	myWallet.Deposit(10)

	got := myWallet.Balance()
	want := 10

	if got != want {
		t.Errorf("got %q expected %q", got, want)
	}
}
