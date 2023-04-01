package strategy

type OnlyPaperStrategy struct {
	won bool
}

func (o *OnlyPaperStrategy) nextHand() *Hand {
	hand := &Hand{handType: PAPER}
	return hand
}

func (o *OnlyPaperStrategy) study(w bool) {
	o.won = w
}
