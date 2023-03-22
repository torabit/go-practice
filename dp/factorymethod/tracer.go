package factorymethod

type Tracer struct {
	Hero
}

func newTracer() IHero {
	return &Tracer{
		Hero{
			name: "Tracer",
			role: Damage,
		},
	}
}
