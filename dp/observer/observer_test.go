package observer

import "testing"

func TestObserver(t *testing.T) {
	snackItem := newItem("Wasa Beaf")
	observerFirst := &Customer{id: "taro"}
	observerSecond := &Customer{id: "kiki"}

	snackItem.register(observerFirst)
	snackItem.register(observerSecond)

	snackItem.notifyAll()

	expected_1 := "taro:Wasa Beaf now in stock"
	expected_2 := "kiki:Wasa Beaf now in stock"

	if !(observerFirst.lastNotify == expected_1) {
		t.Fatalf("want: %+v\ngot: %+v", expected_1, observerFirst.lastNotify)
	}

	if !(observerSecond.lastNotify == expected_2) {
		t.Fatalf("want: %+v\ngot: %+v", expected_2, observerSecond.lastNotify)
	}
}
