package state

type VaultFrame struct {
	currentState      IState
	recordLog         string
	securityCenterLog string
}

func (v *VaultFrame) setClock(h int) {
	v.currentState.doClock(v, h)
}

func (v *VaultFrame) changeState(s IState) {
	v.currentState = s
}

func (v *VaultFrame) setSecurityCenterLog(m string) {
	v.securityCenterLog = "call!: " + m
}

func (v *VaultFrame) setRecordLog(m string) {
	v.recordLog = "record...: " + m
}

func (v *VaultFrame) use() {
	v.currentState.doUse(v)
}

func (v *VaultFrame) alarm() {
	v.currentState.doAlarm(v)
}

func (v *VaultFrame) phone() {
	v.currentState.doPhone(v)
}
