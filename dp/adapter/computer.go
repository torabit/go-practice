package adapter

type OS int
type Connector int

const (
	MacOS OS = iota
	WindowsOS
)

const (
	Lightning Connector = iota
	USB
)

type ComStatus struct {
	os        OS
	connector Connector
}

type Computer interface {
	InsertIntoLightningPort() *ComStatus
}
