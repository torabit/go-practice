package factorymethod

type Tracer struct {
	Hero
}

func newTracer() iHero {
	return &Tracer{
		Hero{
			name: "Tracer",
			role: Damage,
		},
	}
}
