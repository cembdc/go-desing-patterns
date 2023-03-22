package main

import "fmt"

type PayPal struct{}

func (p *PayPal) MakePayment(amount float32) bool {
	// connect to PayPal and process payment
	fmt.Println("PayPal payment done")
	return true
}

type Amazon struct{}

func (a *Amazon) PayAmazon(amount float32) bool {
	// connect to Amazon and process payment
	fmt.Println("Amazon payment done")
	return true
}

type PaymentGateway interface {
	ProcessPayment(amount float32) bool
}

type PayPalAdapter struct {
	PayPal *PayPal
}

func (p *PayPalAdapter) ProcessPayment(amount float32) bool {
	return p.PayPal.MakePayment(amount)
}

type AmazonAdapter struct {
	Amazon *Amazon
}

func (a *AmazonAdapter) ProcessPayment(amount float32) bool {
	return a.Amazon.PayAmazon(amount)
}

// main.go

func main() {
	paymentGateway := &PayPalAdapter{
		PayPal: &PayPal{},
	}
	paymentGateway2 := &AmazonAdapter{
		Amazon: &Amazon{},
	}
	// triggers PayPal
	paymentGateway.ProcessPayment(100)

	// triggers Amazon
	paymentGateway2.ProcessPayment(100)
}
