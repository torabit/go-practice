package state

import "sync"

type DayState struct{}

var singleDayState *DayState

func getDayStateInstance() *DayState {
	var lock = &sync.Mutex{}
	if singleDayState == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleDayState == nil {
			singleDayState = &DayState{}
		}
	}
	return singleDayState
}

func (d *DayState) doClock(c Context, hour int) {
	if hour < 9 || 17 <= hour {
		nightStateInstance := getNightStateInstance()
		c.changeState(nightStateInstance)
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
