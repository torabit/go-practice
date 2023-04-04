package mediator

import (
	"reflect"
	"testing"
)

func TestMediator(t *testing.T) {
	stationManager := newStationManager()

	passengerTrain := &PassengerTrain{
		mediator: stationManager,
	}
	freightTrain := &FreightTrain{
		mediator: stationManager,
	}

	passengerTrain.arrive()
	freightTrain.arrive()

	expected := []string{"PassengerTrain"}

	if !(reflect.DeepEqual(stationManager.arrivedTrain, expected)) {
		t.Fatalf("\nwant: %+v\ngot: %+v", expected, stationManager.arrivedTrain)
	}
}
