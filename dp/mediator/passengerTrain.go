package mediator

type PassengerTrain struct {
	mediator Mediator
}

func (p *PassengerTrain) arrive() {
	if !p.mediator.canArrive(p) {
		return
	}
	p.mediator.setArrivedTrain("PassengerTrain")
}

func (p *PassengerTrain) depart() {
	p.mediator.notifyAboutDeparture()
}

func (p *PassengerTrain) permitArrival() {
	p.arrive()
	p.mediator.setDepartedTrain("PassengerTrain")
}
