package bridge

type tv struct {
	power   bool
	volume  int
	channel int
}

func (t *tv) isEnabled() bool {
	return t.power
}

func (t *tv) enable() {
	t.power = true
}

func (t *tv) disable() {
	t.power = false
}

func (t *tv) getVolume() int {
	return t.volume
}

func (t *tv) setVolume(v int) {
	if v > 101 {
		return
	}

	if v < 0 {
		return
	}

	t.volume = v
}

func (t *tv) getChannel() int {
	return t.channel
}

func (t *tv) setChannel(c int) {
	if c > 12 {
		t.channel = 1
		return
	}

	if c < 1 {
		t.channel = 12
		return
	}

	t.channel = c
}
