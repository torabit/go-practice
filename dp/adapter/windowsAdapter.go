package adapter

type WindowsAdapter struct {
	windowMachine *Windows
}

func (w *WindowsAdapter) InsertIntoLightningPort() *ComStatus {
	stats := &ComStatus{}
	windowsStats := w.windowMachine.insertIntoUSBPort()

	stats.os = windowsStats.os
	stats.connector = Lightning

	return stats
}
