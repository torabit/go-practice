package factorymethod

type Kiriko struct {
	Hero
}

func newKiriko() IHero {
	return &Kiriko{
		Hero{
			name: "Kiriko",
			role: Support,
		},
	}
}
