package state

import "sync"

type NightState struct{}

var singleNightState *NightState

func getNightStateInstance() *NightState {
	var lock = &sync.Mutex{}
	if singleNightState == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleNightState == nil {
			singleNightState = &NightState{}
		}
	}
	return singleNightState
}

func (n *NightState) doClock(c Context, hour int) {
	if 9 <= hour && hour < 17 {
		dayStateInstance := getDayStateInstance()
		c.changeState(dayStateInstance)
	}
}

func (n *NightState) doUse(c Context) {
	c.setSecurityCenterLog("[Night]: Use Vault")
}

func (n *NightState) doAlarm(c Context) {
	c.setSecurityCenterLog("[Night]: Emergency Call")
}

func (n *NightState) doPhone(c Context) {
	c.setRecordLog("[Night]: Record voice")
}

func (n *NightState) toString() string {
	return "[Night]"
}
