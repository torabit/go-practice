package command

type OffCommand struct {
	device IDevice
}

func (o *OffCommand) execute() {
	o.device.off()
}
