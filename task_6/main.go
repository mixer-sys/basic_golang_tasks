package main

import (
	"errors"
	"fmt"
)

type BankAccount struct {
	balance float64
}

func (ba *BankAccount) Deposit(amount float64) {
	if amount > 0 {
		ba.balance += amount
	}
}

func (ba *BankAccount) Withdraw(amount float64) error {
	if amount > ba.balance {
		return errors.New("Not enouf money")
	}
	ba.balance -= amount
	return nil

}

func (ba *BankAccount) Balance() float64 {
	return ba.balance
}

func main() {
	account := BankAccount{}
	account.Deposit(500)
	err := account.Withdraw(200)
	fmt.Printf("Balance %.2f\n", account.balance)
	fmt.Println(err)
	err = account.Withdraw(400)
	fmt.Printf("Balance %.2f\n", account.balance)
	fmt.Println(err)
}
