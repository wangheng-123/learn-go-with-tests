package pointers_and_errors

import (
	"fmt"
	"testing"
)

type Wallet struct {
	balance int
}

func (w *Wallet) Deposit(amount int) {
	fmt.Println("address of balance in Deposit is", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() int {
	return w.balance
}

func TestWalllet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(10)

	got := wallet.Balance()
	fmt.Println("address of balance in test is", &wallet.balance)
	want := 10
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
