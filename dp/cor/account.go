package cor

type IAccount interface {
	pay(*PaymentInfo)
	canPay(int) bool
	setNext(IAccount)
}
