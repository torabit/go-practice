package observer

type Customer struct {
	id         string
	lastNotify string
}

func (c *Customer) update(itemName string) {
	c.lastNotify = c.id + ":" + itemName + " now in stock"
}

func (c *Customer) getID() string {
	return c.id
}
