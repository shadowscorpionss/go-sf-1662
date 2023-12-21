package main

import (
	"fmt"
	"sync"
)

type bankClient struct {
	balance int
	rwLock  sync.RWMutex
}

func NewBankClient() BankClient {
	return &bankClient{
		rwLock: sync.RWMutex{},
	}
}

// Deposit deposits given amount to clients account
func (bc *bankClient) Deposit(amount int) {
	if amount < 0 {
		return
	}

	bc.rwLock.Lock()
	bc.balance += amount
	bc.rwLock.Unlock()
}

// Withdrawal withdraws given amount from clients account.
// return error if clients balance less the withdrawal amount
func (bc *bankClient) Withdrawal(amount int) error {
	var saldo int
	bc.rwLock.RLock()
	saldo = bc.balance
	bc.rwLock.RUnlock()

	if amount > saldo {
		return fmt.Errorf("balance (%d) is less then withdrawal amount (%d). Operation canceled", saldo, amount)
	}

	bc.rwLock.Lock()
	bc.balance -= amount
	bc.rwLock.Unlock()
	return nil
}

// Balance returns clients balance
func (bc *bankClient) Balance() int {
	bc.rwLock.RLock()
	defer bc.rwLock.RUnlock()
	return bc.balance
}
