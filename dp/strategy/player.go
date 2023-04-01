package strategy

type Player struct {
	name      string
	strategy  Strategy
	wincount  int
	losecount int
	gamecount int
}

func (p *Player) nextHand() *Hand {
	return p.strategy.nextHand()
}

func (p *Player) win() {
	p.strategy.study(true)
	p.wincount++
	p.gamecount++
}

func (p *Player) lose() {
	p.strategy.study(false)
	p.losecount++
	p.gamecount++
}

func (p *Player) even() {
	p.gamecount++
}
