package strategy

type Strategy interface {
	nextHand() *Hand
	study(bool)
}
