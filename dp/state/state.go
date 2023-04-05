package state

type IState interface {
	doClock(Context, int)
	doUse(Context)
	doAlarm(Context)
	doPhone(Context)
}
