package main

import "fmt"

type PaymentHandler interface {
	Pay()
}

type Cash struct {
	Amount float32
}

type Bank struct {
	Amount float32
}

type BankAdapter struct {
	Bank
	Account int
}

func (ba *BankAdapter) Pay() {
	ba.Bank.Pay(ba.Account)
}

func (c *Cash) Pay() {
	fmt.Printf("paying cash %.2f", c.Amount)
}

func (b *Bank) Pay(account int) {
	fmt.Printf("paying bank %.2f", b.Amount)
}

func main() {
	cash := Cash{
		Amount: 5,
	}
	cash.Pay()
	bank := BankAdapter{
		Bank: Bank{
			Amount: 10,
		},
		Account: 12345,
	}
	bank.Pay()
}
