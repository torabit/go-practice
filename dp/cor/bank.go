package cor

type Bank struct {
	next    IAccount
	balance int
}

func (b *Bank) canPay(price int) bool {
	return b.balance >= price
}

func (b *Bank) pay(pi *PaymentInfo) {
	if b.canPay(pi.price) {
		pi.method = "Bank"
		pi.succeed = true
		b.balance -= pi.price
	} else if b.next != nil {
		b.next.pay(pi)
	} else {
		pi.succeed = false
	}
}

func (b *Bank) setNext(next IAccount) {
	b.next = next
}
