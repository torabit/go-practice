package state

import "sync"

type DayState struct{}

var singleDayState *DayState

func getDayStateInstance(isCreatedCh chan<- bool) *DayState {
	var lock = &sync.Mutex{}
	if singleDayState == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleDayState == nil {
			isCreatedCh <- false
			singleDayState = &DayState{}
		} else {
			isCreatedCh <- true
		}
	} else {
		isCreatedCh <- true
	}
	return singleDayState
}

func (d *DayState) doClock(c Context, hour int) {
	if hour < 9 || 17 <= hour {
		isCreatedCh := make(chan bool)
		defer close(isCreatedCh)
		go getNightStateInstance(isCreatedCh)
		c.changeState(singleNightState)
	}
}

func (d *DayState) doUse(c Context) {
	c.setRecordLog("[Noon]: Use Vault")
}

func (d *DayState) doAlarm(c Context) {
	c.setSecurityCenterLog("[Noon]: Emergency Call")
}

func (d *DayState) doPhone(c Context) {
	c.setSecurityCenterLog("[Noon]: Call")
}

func (d *DayState) toString() string {
	return "[Noon]"
}
