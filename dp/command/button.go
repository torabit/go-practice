package command

type Button struct {
	command ICommand
}

func (b *Button) press() {
	b.command.execute()
}
