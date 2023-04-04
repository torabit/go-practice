package mediator

type StationManager struct {
	isPlatformFree bool
	trainQueue     []ITrain
	arrivedTrain   []string
	departedTrain  []string
}

func newStationManager() *StationManager {
	return &StationManager{
		isPlatformFree: true,
	}
}

func (s *StationManager) canArrive(t ITrain) bool {
	if s.isPlatformFree {
		s.isPlatformFree = false
		return true
	}
	s.trainQueue = append(s.trainQueue, t)
	return false
}

func (s *StationManager) notifyAboutDeparture() {
	if !s.isPlatformFree {
		s.isPlatformFree = true
	}
	if len(s.trainQueue) > 0 {
		firstTrainQueue := s.trainQueue[0]
		s.trainQueue = s.trainQueue[1:]
		firstTrainQueue.permitArrival()
	}
}

func (s *StationManager) setArrivedTrain(n string) {
	s.arrivedTrain = append(s.arrivedTrain, n)
}

func (s *StationManager) setDepartedTrain(n string) {
	s.departedTrain = append(s.departedTrain, n)
}
