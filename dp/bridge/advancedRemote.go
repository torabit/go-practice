package bridge

type advancedRemote struct {
	device Device
}

func (a *advancedRemote) setDevice(d Device) {
	a.device = d
}

func (a *advancedRemote) togglePower() {
	if a.device.isEnabled() {
		a.device.disable()
		return
	}
	a.device.enable()
}

func (a *advancedRemote) volumeDown() {
	currentVolume := a.device.getVolume()
	a.device.setVolume(currentVolume - 1)
}

func (a *advancedRemote) volumeUp() {
	currentVolume := a.device.getVolume()
	a.device.setVolume(currentVolume + 1)
}

func (a *advancedRemote) channelDown() {
	currentChannel := a.device.getChannel()
	a.device.setChannel(currentChannel - 1)
}

func (a *advancedRemote) channelUp() {
	currentChannel := a.device.getChannel()
	a.device.setChannel(currentChannel + 1)
}
