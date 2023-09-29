package main

import "fmt"

type PaymentStrategy interface {
	Pay(amount int) bool
}

type CreditCardStrategy struct{}

func (s *CreditCardStrategy) Pay(amount int) bool {
	fmt.Printf("payment successful by card\n")
	return true
}

type CashStrategy struct{}

func (s *CashStrategy) Pay(amount int) bool {
	fmt.Printf("payment successful by cash\n")
	return true
}

type Order struct {
	strategy PaymentStrategy
}

func NewOrder(strategy PaymentStrategy) *Order {
	return &Order{strategy: strategy}
}

func (o *Order) Pay(amount int) bool {
	return o.strategy.Pay(amount)
}

func main() {
	order := NewOrder(&CreditCardStrategy{})
	order.Pay(100)

	order.strategy = &CashStrategy{}
	order.Pay(100)
}
