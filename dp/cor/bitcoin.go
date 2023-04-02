package cor

type Bitcoin struct {
	next    IAccount
	balance int
}

func (b *Bitcoin) canPay(price int) bool {
	return b.balance >= price
}

func (b *Bitcoin) pay(pi *PaymentInfo) {
	if b.canPay(pi.price) {
		pi.method = "Bitcoin"
		pi.succeed = true
		b.balance -= pi.price
	} else if b.next != nil {
		b.next.pay(pi)
	} else {
		pi.succeed = false
	}
}

func (b *Bitcoin) setNext(next IAccount) {
	b.next = next
}
