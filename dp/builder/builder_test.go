package builder

import "testing"

func TestBuilder(t *testing.T) {
	normalBuilder := getBuilder(normal)
	iglooBuilder := getBuilder(igloo)

	director := newDirector(normalBuilder)

	t.Run("case buildHouse method", func(t *testing.T) {
		normalHouse := director.buildHouse()

		if normalHouse.doorType != "Wooden Door" {
			t.Fatal("This houses door type is not Wooden\ndoor type is:", normalHouse.doorType)
		}
		if normalHouse.windowType != "Wooden Window" {
			t.Fatal("This houses window type is not Wooden\nwinodw type is:", normalHouse.windowType)
		}
		if normalHouse.floor != 2 {
			t.Fatal("This house is not 2 floor")
		}
	})

	t.Run("case setBuilder method", func(t *testing.T) {
		director.setBuilder(iglooBuilder)
		iglooHouse := director.buildHouse()

		if iglooHouse.doorType != "Snow Door" {
			t.Fatal("This houses door type is not Snow\ndoor type is:", iglooHouse.doorType)
		}
		if iglooHouse.windowType != "Snow Window" {
			t.Fatal("This houses window type is not Snow\nwinodw type is:", iglooHouse.windowType)
		}
		if iglooHouse.floor != 1 {
			t.Fatal("This house is not 1 floor")
		}
	})
}
