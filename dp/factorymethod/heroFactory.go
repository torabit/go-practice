package factorymethod

import "fmt"

func getHeroByName(name string) (iHero, error) {

	if name == "winston" {
		return newWinston(), nil
	}

	if name == "tracer" {
		return newTracer(), nil
	}
	if name == "kiriko" {
		return newKiriko(), nil
	}

	return nil, fmt.Errorf("Wrong Hero name passed")
}

func getHeroByRole(role Role) (iHero, error) {

	if role == Tank {
		return newWinston(), nil
	}

	if role == Damage {
		return newTracer(), nil
	}

	if role == Support {
		return newKiriko(), nil
	}

	return nil, fmt.Errorf("Wrong Hero role passed")
}
