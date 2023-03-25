package abstractfactory

type Gucci struct{}

func (g *Gucci) makeShoe() IShoe {
	return &GucciShoe{
		Shoe{
			logo:  "gucci",
			size:  14,
			price: 1000,
		},
	}
}

func (g *Gucci) makeJacket() IJacket {
	return &GucciJacket{
		Jacket{
			logo:  "gucci",
			size:  15,
			price: 2500,
		},
	}
}
