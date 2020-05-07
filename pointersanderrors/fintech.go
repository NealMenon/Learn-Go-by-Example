package main

// Wallet struct
type Wallet struct {
	balance Bitcoin
}

// Bitcoin struct of type int
type Bitcoin int

// Balance in wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// Deposit amt into wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func main() {

}
