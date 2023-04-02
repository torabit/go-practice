package decorator

import "testing"

func TestDecorator(t *testing.T) {
	pizza := &VeggeMania{}

	pizzaWithCheese := &CheeseTopping{pizza: pizza}
	pizzaWithCheeseAndTomato := &TomatoTopping{pizza: pizzaWithCheese}

	if !(pizzaWithCheeseAndTomato.getPrice() == 32) {
		t.Fatalf("\nwant: %+v\ngot: %+v", 32, pizzaWithCheeseAndTomato.getPrice())
	}
}
