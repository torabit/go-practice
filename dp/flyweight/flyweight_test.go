package flyweight

import "testing"

func TestFlyweight(t *testing.T) {
	game := newGame()

	for i := 0; i < 5; i++ {
		game.addPlayer(T)
	}
	for i := 0; i < 5; i++ {
		game.addPlayer(CT)
	}

	dressFactoryInstance := getDressFactorySingleInstance()
	t.Run("case: are there any duplicate instances", func(t *testing.T) {
		dressMap := dressFactoryInstance.dressMap
		if !(len(dressMap) == 2) {
			t.Fatalf("want: %+v\ngot: %+v", 2, len(dressMap))
		}
	})

	t.Run("case: get dress color", func(t *testing.T) {
		tDress := dressFactoryInstance.dressMap[T]
		ctDress := dressFactoryInstance.dressMap[CT]
		if !(tDress.getColor() == "red") {
			t.Fatalf("want: %+v\ngot: %+v", "red", tDress.getColor())
		}

		if !(ctDress.getColor() == "blue") {
			t.Fatalf("want: %+v\ngot: %+v", "blue", ctDress.getColor())
		}

	})
}
