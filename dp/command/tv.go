package command

type Tv struct {
	isRunning bool
}

func (t *Tv) on() {
	t.isRunning = true
}

func (t *Tv) off() {
	t.isRunning = false
}
