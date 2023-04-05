package command

import "testing"

func TestCommand(t *testing.T) {
	tv := &Tv{}
	onCommand := &OnCommand{
		device: tv,
	}
	offCommand := &OffCommand{
		device: tv,
	}

	t.Run("case: turn on TV", func(t *testing.T) {
		onButton := &Button{
			command: onCommand,
		}
		onButton.press()
		if !(tv.isRunning == true) {
			t.Fatal("tv is running")
		}
	})

	t.Run("case: turn off TV", func(t *testing.T) {
		offButton := &Button{
			command: offCommand,
		}
		offButton.press()
		if !(tv.isRunning == false) {
			t.Fatal("tv is not running")
		}
	})

}
