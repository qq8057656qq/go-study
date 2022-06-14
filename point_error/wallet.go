package point_error

import (
	"errors"
	"fmt"
)

var InsufficientFoundsError = errors.New("cannot withdraw, insufficient funds")

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (wallet *Wallet) Deposit(amount Bitcoin) {
	wallet.balance += amount
}

func (wallet *Wallet) Balance() Bitcoin {
	return wallet.balance
}

func (wallet *Wallet) Withdraw(amount Bitcoin) error {
	if wallet.balance < amount {
		return InsufficientFoundsError
	}
	wallet.balance -= amount
	return nil
}
