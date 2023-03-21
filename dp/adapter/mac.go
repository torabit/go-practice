package adapter

type Mac struct{}

func (m *Mac) InsertIntoLightningPort() *ComStatus {
	stats := &ComStatus{}

	stats.os = MacOS
	stats.connector = Lightning

	return stats
}
