package cor

import "testing"

func TestChainOfResponsibility(t *testing.T) {
	bank := &Bank{balance: 1000}
	bitcoin := &Bitcoin{balance: 2000}
	paypal := &Paypal{balance: 3000}

	paymentInfo := &PaymentInfo{
		price:       3000,
		productname: "Book",
		method:      "TBD",
	}
	bitcoin.setNext(paypal)
	bank.setNext(bitcoin)
	bank.pay(paymentInfo)

	if !(paymentInfo.method == "Paypal") {
		t.Fatalf("%+v", paymentInfo)
	}

}
