package adapter

import "testing"

func TestAdapter(t *testing.T) {

	client := &Client{}

	t.Run("case for mac", func(t *testing.T) {
		mac := &Mac{}
		c := client.InsertLightningConnectorIntoComputer(mac)

		if c.os != MacOS {
			t.Fatal("OS is not Mac")
		}

		if c.connector != Lightning {
			t.Fatal("Lightning connector does not plugged into mac machine")
		}
	})

	t.Run("case for windows", func(t *testing.T) {

		windowsMachine := &Windows{}
		windowsMachineAdapter := &WindowsAdapter{
			windowMachine: windowsMachine,
		}

		c := client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)

		if c.os != WindowsOS {
			t.Fatal("OS is not Windows")
		}

		if c.connector != Lightning {
			t.Fatal("Adapter has not converts USB")
		}

	})
}
