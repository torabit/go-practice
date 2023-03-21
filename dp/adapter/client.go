package adapter

type Client struct{}

func (c *Client) InsertLightningConnectorIntoComputer(com Computer) *ComStatus {
	stats := com.InsertIntoLightningPort()

	return stats
}
