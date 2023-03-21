package factorymethod

import "testing"

func TestFactoryMethod(t *testing.T) {

	t.Run("case winston", func(t *testing.T) {
		winston, _ := getHeroByName("winston")

		if winston.getName() != "Winston" {
			t.Fatal(winston.getName(), "is not Winston")
		}

		winston, _ = getHeroByRole(Tank)

		if winston.getName() != "Winston" {
			t.Fatal(winston.getName(), "is not Winston")
		}
	})

	t.Run("case tracer", func(t *testing.T) {
		tracer, _ := getHeroByName("tracer")

		if tracer.getName() != "Tracer" {
			t.Fatal(tracer.getName(), "is not Tracer")
		}

		tracer, _ = getHeroByRole(Damage)

		if tracer.getName() != "Tracer" {
			t.Fatal(tracer.getName(), "is not Tracer")
		}
	})

	t.Run("case kiriko", func(t *testing.T) {
		kiriko, _ := getHeroByName("kiriko")

		if kiriko.getName() != "Kiriko" {
			t.Fatal(kiriko.getName(), "is not Kiriko")
		}

		kiriko, _ = getHeroByRole(Support)

		if kiriko.getName() != "Kiriko" {
			t.Fatal(kiriko.getName(), "is not Kiriko")
		}
	})
}
