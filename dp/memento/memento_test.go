package memento

import "testing"

func TestMemento(t *testing.T) {
	caretaker := &Caretaker{
		mementoArray: make([]*Memento, 0),
	}

	originator := &Originator{
		state: "A",
	}

	caretaker.addMemento(originator.createMemento())

	t.Run("case: overwrite state", func(t *testing.T) {
		originator.setState("B")
		if !(originator.getState() == "B") {
			t.Fatalf("want: %+v\ngot: %+v", "B", originator.getState())
		}
	})

	t.Run("case: restore state", func(t *testing.T) {
		originator.restoreMemento(caretaker.getMemento(0))
		if !(originator.getState() == "A") {
			t.Fatalf("want: %+v\ngot: %+v", "A", originator.getState())
		}
	})
}
