package cor

type Paypal struct {
	next    IAccount
	balance int
}

func (p *Paypal) canPay(price int) bool {
	return p.balance >= price
}

func (p *Paypal) pay(pi *PaymentInfo) {
	if p.canPay(pi.price) {
		pi.method = "Paypal"
		pi.succeed = true
		p.balance -= pi.price
	} else if p.next != nil {
		p.next.pay(pi)
	} else {
		pi.succeed = false
	}
}

func (p *Paypal) setNext(next IAccount) {
	p.next = next
}
