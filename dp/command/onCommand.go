package command

type OnCommand struct {
	device IDevice
}

func (o *OnCommand) execute() {
	o.device.on()
}
