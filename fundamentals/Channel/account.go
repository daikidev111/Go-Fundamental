package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

var pl = fmt.Println

type Account struct {
	balance int
	lock sync.Mutex // for obtaining a lock for mutual exclusion
}

func (a *Account) getBalance() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.balance
}

func (a *Account) withdraw(v int) error {
	a.lock.Lock()
	defer a.lock.Unlock()

	if (a.balance < v) {
		pl("Not enough balance")
		return errors.New("Error")
	} else {
		a.balance = a.balance - v
		fmt.Printf("%d withdrawn : Balance: %d\n", v, a.balance)
		return nil
	}
}

func main() {
	var acc Account
	acc.balance = 100
	pl("Balance:", acc.getBalance())
	for i := 0; i<12; i++ {
		go acc.withdraw(10)
	}
	time.Sleep(2 * time.Second)
}
