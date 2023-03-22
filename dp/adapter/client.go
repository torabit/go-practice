package adapter

type Client struct{}

func (c *Client) InsertLightningConnectorIntoComputer(com IComputer) *ComStatus {
	stats := com.InsertIntoLightningPort()

	return stats
}
