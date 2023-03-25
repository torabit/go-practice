package abstractfactory

import (
	"errors"
	"testing"
)

func TestAbstractFactory(t *testing.T) {
	t.Run("case: Does get a luxury brand factory", func(t *testing.T) {

		_, got := GetLuxuryBrandFactory("gucci")
		want := FactoryError

		if !errors.Is(got, want) {
			t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), want, got)
		}
	})

	t.Run("case: Shoe method test", func(t *testing.T) {

		gucciFactory, _ := GetLuxuryBrandFactory("Gucci")

		gucciShoe := gucciFactory.makeShoe()

		gucciShoe.setLogo("gucci")
		gucciShoe.setSize(17)
		gucciShoe.setPrice(3000)

		if gucciShoe.getLogo() != "gucci" {
			t.Fatalf("want: %+v\ngot: %+v\n", "gucci", gucciShoe.getLogo())
		}

		if gucciShoe.getSize() != 17 {
			t.Fatalf("want: %+v\ngot: %+v\n", 17, gucciShoe.getSize())
		}

		if gucciShoe.getPrice() != 3000 {
			t.Fatalf("want: %+v\ngot: %+v\n", 3000, gucciShoe.getPrice())
		}
	})
}
