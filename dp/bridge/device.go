package bridge

type Device interface {
	isEnabled() bool
	enable()
	disable()
	getVolume() int
	setVolume(int)
	getChannel() int
	setChannel(int)
}
