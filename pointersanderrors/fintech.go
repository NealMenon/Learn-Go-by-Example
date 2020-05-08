package main

import (
	"errors"
	"fmt"
)

// ErrInsufficientFunds error value
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Wallet struct
type Wallet struct {
	balance Bitcoin
}

// Bitcoin struct of type int
type Bitcoin int

// Stringer gives string value
type Stringer interface {
	String() string
}

// String toString() in go
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Balance in wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// Deposit amt into wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Withdraw amt into wallet
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.Balance() {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil

}

func main() {

}
