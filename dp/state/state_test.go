package state

import "testing"

func TestState(t *testing.T) {
	isCreatedCh := make(chan bool)
	defer close(isCreatedCh)
	go getDayStateInstance(isCreatedCh)

	vaultFrame := &VaultFrame{currentState: singleDayState}

	t.Run("case: use vault while at noon", func(t *testing.T) {
		vaultFrame.use()
		expectedLog := "record...: [Noon]: Use Vault"
		if !(vaultFrame.recordLog == expectedLog) {
			t.Fatalf("want: %+v\ngot: %+v", expectedLog, vaultFrame.recordLog)
		}
	})

	t.Run("case: use phone while at night", func(t *testing.T) {
		vaultFrame.setClock(24)
		vaultFrame.phone()
		expectedLog := "record...: [Night]: Record voice"
		if !(vaultFrame.recordLog == expectedLog) {
			t.Fatalf("want: %+v\ngot: %+v", expectedLog, vaultFrame.recordLog)
		}
	})

	t.Run("case: use vault while at night", func(t *testing.T) {
		vaultFrame.use()
		expectedLog := "call!: [Night]: Use Vault"
		if !(vaultFrame.securityCenterLog == expectedLog) {
			t.Fatalf("want: %+v\ngot: %+v", expectedLog, vaultFrame.securityCenterLog)
		}
	})

	t.Run("case: alart while at noon", func(t *testing.T) {
		vaultFrame.setClock(10)
		vaultFrame.alarm()
		expectedLog := "call!: [Noon]: Emergency Call"
		if !(vaultFrame.securityCenterLog == expectedLog) {
			t.Fatalf("want: %+v\ngot: %+v", expectedLog, vaultFrame.securityCenterLog)
		}
	})
}
