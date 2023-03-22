package singleton

import "testing"

func TestSingleton(t *testing.T) {
	var myclass = &SingletonClass{}

	t.Run("case The same instance", func(t *testing.T) {

		one := myclass.getInstance(1)
		two := myclass.getInstance(2)

		if one != two {
			t.Fatal("instance does not same")
		}
	})
}
