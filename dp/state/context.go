package state

type Context interface {
	setClock(int)
	changeState(IState)
	setSecurityCenterLog(string)
	setRecordLog(string)
}
