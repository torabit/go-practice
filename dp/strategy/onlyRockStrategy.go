package strategy

type OnlyRockStrategy struct {
	won bool
}

func (o *OnlyRockStrategy) nextHand() *Hand {
	hand := &Hand{handType: ROCK}
	return hand
}

func (o *OnlyRockStrategy) study(w bool) {
	o.won = w
}
