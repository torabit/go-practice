package bridge

type Remote interface {
	setDevice(*Device)
	togglePower()
	volumeDown()
	volumeUp()
	channelDown()
	channelUp()
}
