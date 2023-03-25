package abstractfactory

import (
	"errors"
)

type ILuxuryBrandFactory interface {
	makeShoe() IShoe
	makeJacket() IJacket
}

var FactoryError = errors.New("Wrong brand type passed")

func GetLuxuryBrandFactory(brand string) (ILuxuryBrandFactory, error) {

	if brand == "Gucci" {
		return &Gucci{}, nil
	}

	return nil, FactoryError
}
