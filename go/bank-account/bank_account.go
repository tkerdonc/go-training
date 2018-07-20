// This package implements a bank account
package account

import "sync"

type Account struct {
	amount int64
	closed bool
	mutex  *sync.Mutex
}

// Open creates a new account with the inital balance specified as a
// parameter.
func Open(initialBalance int64) *Account {
	if initialBalance < 0 {
		return nil
	}

	var mutex = &sync.Mutex{}
	var a = Account{initialBalance, false, mutex}
	return &a
}

// Close closes the account
func (account *Account) Close() (int64, bool) {
	account.mutex.Lock()
	defer account.mutex.Unlock()
	if account.closed {
		return 0, false
	}
	account.closed = true
	return account.amount, true
}

// Balance returns the current balance of the account
func (account *Account) Balance() (int64, bool) {
	account.mutex.Lock()
	defer account.mutex.Unlock()
	if account.closed {
		return 0, false
	}
	return account.amount, true
}

// Deposit adds the value passed as a parameter to the account's balance
func (account *Account) Deposit(amount int64) (int64, bool) {
	account.mutex.Lock()
	defer account.mutex.Unlock()
	if account.closed || account.amount+amount < 0 {
		return 0, false
	}
	account.amount += amount
	return account.amount, true
}
