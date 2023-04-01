package strategy

type HandType int

const (
	ROCK HandType = iota
	SCISSORS
	PAPER
)

type Hand struct {
	handType HandType
}

func (h *Hand) fight(e *Hand) int {
	handType := e.handType

	if h.handType == handType {
		return 0
	} else if (h.handType+1)%3 == e.handType {
		return 1
	} else {
		return -1
	}
}

func (h *Hand) isStrongerThan(e *Hand) bool {
	return h.fight(e) == 1
}

func (h *Hand) isWeakerThan(e *Hand) bool {
	return h.fight(e) == -1
}

func (h *Hand) getHand(ht HandType) *Hand {
	h.handType = ht
	return h
}
