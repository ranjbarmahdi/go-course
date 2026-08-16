package main

import "fmt"

type BankAccount struct {
	Owner   string
	Balance float64
}

func (b *BankAccount) Deposit(amount float64) {
	b.Balance += amount
}

func (b *BankAccount) Withdraw(amount float64) bool {
	if b.Balance < amount {
		return false
	}

	b.Balance -= amount
	return true
}

func main() {
	account := BankAccount{
		Owner:   "Mahdi",
		Balance: 100,
	}

	account.Deposit(50)

	fmt.Println(account.Balance)
}
