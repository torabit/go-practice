package flyweight

type PlayerType int

const (
	T PlayerType = iota
	CT
)

type Player struct {
	dress      IDress
	playerType PlayerType
	lat        int
	long       int
}

func newPlayer(playerType PlayerType) *Player {
	dress, _ := getDressFactorySingleInstance().getDressByType(playerType)
	return &Player{
		playerType: playerType,
		dress:      dress,
	}
}

func (p *Player) nowLocation(lat, long int) {
	p.lat = lat
	p.long = long
}
