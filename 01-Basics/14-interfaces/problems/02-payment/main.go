package main

import "fmt"

type PaymentMethod interface {
	Pay(amount float64)
}

type CreditCard struct {
}

func (c CreditCard) Pay(amount float64) {
	fmt.Printf("Payed %.2f By Credit Card\n", amount)
}

type PayPal struct {
}

func (p PayPal) Pay(amount float64) {
	fmt.Printf("Payed %.2f By Pay Pal\n", amount)
}

func pay(pay PaymentMethod, amount float64) {
	pay.Pay(amount)
}

func checkout(p PaymentMethod, amount float64) {

}

func main() {
	creditCard := CreditCard{}
	payPal := PayPal{}

	pay(creditCard, 100)
	pay(payPal, 100)
}
