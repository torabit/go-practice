package adapter

type Windows struct{}

func (w *Windows) insertIntoUSBPort() *ComStatus {
	stats := &ComStatus{}

	stats.os = WindowsOS
	stats.connector = USB

	return stats
}
