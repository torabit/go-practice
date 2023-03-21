package factorymethod

type Kiriko struct {
	Hero
}

func newKiriko() iHero {
	return &Kiriko{
		Hero{
			name: "Kiriko",
			role: Support,
		},
	}
}
