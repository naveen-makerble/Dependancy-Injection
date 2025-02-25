package main

import "fmt"

type PaymentProcessor interface {
	ProcessPayment(amount float64) string
}

type PaypalProcessor struct{}

func (p *PaypalProcessor) ProcessPayment(amount float64) string {
	return fmt.Sprintf("Payment of %f via paypal succesful!", amount)
}

type StripeProcessor struct{}

func (p *StripeProcessor) ProcessPayment(amount float64) string {
	return fmt.Sprintf("Payment of %f via paypal succesful!", amount)
}

type Shoppingcart struct {
	paymentProcessor PaymentProcessor
}

func NewShoppingCart(paymentProcessor PaymentProcessor) *Shoppingcart {
	return &Shoppingcart{
		paymentProcessor: paymentProcessor,
	}
}

func (cart *Shoppingcart) Checkout(amount float64) {
	result := cart.paymentProcessor.ProcessPayment(amount)
	fmt.Println(result)
}

func main() {
	paypalProcessor := &PaypalProcessor{}
	shoppingCartPaypal := NewShoppingCart(paypalProcessor)
	shoppingCartPaypal.Checkout(100.0)
}
