package flyweight

type game struct {
	terrorists        []*Player
	counterTerrorists []*Player
}

func newGame() *game {
	return &game{
		terrorists:        make([]*Player, 1),
		counterTerrorists: make([]*Player, 1),
	}
}

func (c *game) addPlayer(playerType PlayerType) {
	if playerType == T {
		player := newPlayer(T)
		c.terrorists = append(c.terrorists, player)
		return
	}
	if playerType == CT {
		player := newPlayer(CT)
		c.counterTerrorists = append(c.counterTerrorists, player)
		return
	}
	return
}
