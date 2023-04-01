package bridge

import "testing"

func TestBridge(t *testing.T) {
	tv := &tv{volume: 0, channel: 1, power: false}
	remote := &advancedRemote{}

	remote.setDevice(tv)

	remote.togglePower()

	if !tv.power {
		t.Fatalf("\nwant: %+v\ngot: %+v\n", true, tv.power)
	}

	remote.volumeDown()

	if !(tv.volume == 0) {
		t.Fatalf("\nwant: %+v\ngot: %+v\n", 0, tv.volume)
	}

	remote.channelDown()

	if !(tv.channel == 12) {
		t.Fatalf("\nwant: %+v\ngot: %+v\n", 12, tv.channel)
	}
}
