package strategy

import "testing"

func TestHandSimulator(t *testing.T) {
	handA := &Hand{handType: ROCK}
	handB := &Hand{handType: SCISSORS}
	handC := &Hand{handType: PAPER}
	handD := &Hand{handType: ROCK}

	t.Run("case: fight method", func(t *testing.T) {
		var result int

		result = handA.fight(handB)
		if !(result == 1) {
			t.Fatalf("\nwant: %+v\ngot: %+v", 1, result)
		}

		result = handA.fight(handD)
		if !(result == 0) {
			t.Fatalf("\nwant: %+v\ngot: %+v", 0, result)
		}

		result = handB.fight(handA)
		if !(result == -1) {
			t.Fatalf("\nwant: %+v\ngot: %+v", -1, result)
		}
	})

	t.Run("case: rock versus scissors", func(t *testing.T) {
		var result bool

		result = handA.isStrongerThan(handB)
		if !(result == true) {
			t.Fatalf("\nwant: %+v\ngot: %+v", true, result)
		}

		result = handA.isWeakerThan(handB)
		if !(result == false) {
			t.Fatalf("\nwant: %+v\ngot: %+v", false, result)
		}

		result = handB.isStrongerThan(handA)
		if !(result == false) {
			t.Fatalf("\nwant: %+v\ngot: %+v", true, result)
		}

		result = handB.isWeakerThan(handA)
		if !(result == true) {
			t.Fatalf("\nwant: %+v\ngot: %+v", false, result)
		}
	})

	t.Run("case: rock versus paper", func(t *testing.T) {
		var result bool

		result = handC.isStrongerThan(handA)
		if !(result == true) {
			t.Fatalf("\nwant: %+v\ngot: %+v", true, result)
		}

		result = handC.isWeakerThan(handA)
		if !(result == false) {
			t.Fatalf("\nwant: %+v\ngot: %+v", false, result)
		}

		result = handA.isStrongerThan(handC)
		if !(result == false) {
			t.Fatalf("\nwant: %+v\ngot: %+v", true, result)
		}

		result = handA.isWeakerThan(handC)
		if !(result == true) {
			t.Fatalf("\nwant: %+v\ngot: %+v", false, result)
		}
	})

	t.Run("case: paper versus scissors", func(t *testing.T) {
		var result bool

		result = handB.isStrongerThan(handC)
		if !(result == true) {
			t.Fatalf("\nwant: %+v\ngot: %+v", true, result)
		}

		result = handB.isWeakerThan(handC)
		if !(result == false) {
			t.Fatalf("\nwant: %+v\ngot: %+v", false, result)
		}

		result = handC.isStrongerThan(handB)
		if !(result == false) {
			t.Fatalf("\nwant: %+v\ngot: %+v", true, result)
		}

		result = handC.isWeakerThan(handB)
		if !(result == true) {
			t.Fatalf("\nwant: %+v\ngot: %+v", false, result)
		}

	})
}

func TestStrategy(t *testing.T) {
	player1 := &Player{
		name:     "player1",
		strategy: &OnlyPaperStrategy{},
	}

	player2 := &Player{
		name:     "player2",
		strategy: &OnlyRockStrategy{},
	}

	for i := 0; i < 10; i++ {
		nextHand1 := player1.nextHand()
		nextHand2 := player2.nextHand()

		if nextHand1.isStrongerThan(nextHand2) {
			player1.win()
			player2.lose()
		} else if nextHand2.isStrongerThan(nextHand1) {
			player2.win()
			player1.lose()
		} else {
			player1.even()
			player2.even()
		}
	}

	if !(player1.wincount == 10) {
		t.Fatalf("\nwant: %+v\ngot: %+v", 10, player1.wincount)
	}
}
