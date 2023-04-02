package visitor

import (
	"math"
	"testing"
)

func TestVisitor(t *testing.T) {
	square := &Square{side: 2}
	circle := &Circle{radius: 3}
	rectangle := &Rectangle{l: 2, b: 3}

	areaCalc := &AreaCalc{}

	t.Run("case: Square", func(t *testing.T) {
		square.accept(areaCalc)
		if !(areaCalc.area == 4) {
			t.Fatalf("\nwant: %+v\ngot: %+v", 4, areaCalc.area)
		}
	})

	t.Run("case: Circle", func(t *testing.T) {
		circle.accept(areaCalc)
		if !(areaCalc.area == 3*3*math.Pi) {
			t.Fatalf("\nwant: %+v\ngot: %+v", 3*3*math.Pi, areaCalc.area)
		}
	})

	t.Run("case: Rectangle", func(t *testing.T) {
		rectangle.accept(areaCalc)
		if !(areaCalc.area == 6) {
			t.Fatalf("\nwant: %+v\ngot: %+v", 6, areaCalc.area)
		}
	})
}
