package factorymethod

type Winston struct {
	Hero
}

func newWinston() IHero {
	return &Winston{
		Hero{
			name: "Winston",
			role: Tank,
		},
	}
}
