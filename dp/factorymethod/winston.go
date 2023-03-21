package factorymethod

type Winston struct {
	Hero
}

func newWinston() iHero {
	return &Winston{
		Hero{
			name: "Winston",
			role: Tank,
		},
	}
}
