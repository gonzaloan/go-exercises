package main

import "fmt"

type Payment interface {
	Pay()
}

type CashPayment struct{}

func (CashPayment) Pay() {
	fmt.Println("Payment using Cash")
}
func ProcessPayment(p Payment) {
	p.Pay()
}

type BankPayment struct{}

// Pay : This doesn't fullfil the Payment interface, bc has a parameter
func (BankPayment) Pay(bankAccount int) {
	fmt.Println("Payment using Bank Account ", bankAccount)
}

// BankPaymentAdapter : We create this Adapter to do so
type BankPaymentAdapter struct {
	BankPayment *BankPayment
	bankAccount int
}

// Pay : and implement what we need.
func (bpa *BankPaymentAdapter) Pay() {
	bpa.BankPayment.Pay(bpa.bankAccount)
}
func main() {
	cash := &CashPayment{}
	ProcessPayment(cash)

	bpa := &BankPaymentAdapter{
		bankAccount: 5,
		BankPayment: &BankPayment{},
	}
	ProcessPayment(bpa)
}
