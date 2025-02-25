package main

import "fmt"

type Notifier interface {
	SendNotificaiton(string)
}

type EmailNotifier struct{}

func (e *EmailNotifier) SendNotificaiton(msg string) {
	fmt.Println("Notification sent via email: ", msg)
}

type SMSNotifier struct{}

func (e *SMSNotifier) SendNotificaiton(msg string) {
	fmt.Println("Notification sent via email: ", msg)
}

type OrderService struct {
	notifier Notifier
}

func NewOrderService(notifier Notifier) *OrderService {
	return &OrderService{
		notifier: notifier,
	}
}

func (o *OrderService) PlaceOrder(orderID string) {
	fmt.Printf("Order %s placed successfully.\n", orderID)
	o.notifier.SendNotificaiton("Your order has been placed")
}

func main() {
	emailNotifier := &EmailNotifier{}
	orderServiceWithEmail := NewOrderService(emailNotifier)
	orderServiceWithEmail.PlaceOrder("12345")

	smsNotifier := &SMSNotifier{}
	orderServiceWithSMS := NewOrderService(smsNotifier)
	orderServiceWithSMS.PlaceOrder("67890")
}
