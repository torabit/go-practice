package singleton

import "testing"

func TestSingleton(t *testing.T) {
	t.Run("case The same instance", func(t *testing.T) {
		isCreatedCh := make(chan bool)
		defer close(isCreatedCh)

		for i := 0; i < 30; i++ {
			go getInstance(isCreatedCh)
			b := <-isCreatedCh
			if i == 0 {
				if b == true {
					t.Fatal("The instance already created")
				}
			} else {
				if b == false {
					t.Fatal("The instance does not created")
				}
			}
		}
	})
}
