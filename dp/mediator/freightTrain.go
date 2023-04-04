package mediator

type FreightTrain struct {
	mediator Mediator
}

func (p *FreightTrain) arrive() {
	if !p.mediator.canArrive(p) {
		return
	}
	p.mediator.setArrivedTrain("FreightTrain")
}

func (p *FreightTrain) depart() {
	p.mediator.notifyAboutDeparture()
}

func (p *FreightTrain) permitArrival() {
	p.arrive()
	p.mediator.setDepartedTrain("FreightTrain")
}
