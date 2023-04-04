package mediator

type Mediator interface {
	canArrive(ITrain) bool
	notifyAboutDeparture()
	setArrivedTrain(string)
	setDepartedTrain(string)
}
