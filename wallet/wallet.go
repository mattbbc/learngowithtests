package wallet

import (
	"errors"
	"fmt"
)

// Why do this? Why not just use int?
// It means we can create methods on Bitcoin
type Bitcoin int

// For example, implement Stringer interface
// so we can use %s when printing
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

var ErrInsuffucientFunds = errors.New("insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if w.balance >= amount {
		w.balance -= amount
		return nil
	}

	return ErrInsuffucientFunds
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
