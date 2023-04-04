package observer

type IObserver interface {
	update(string)
	getID() string
}
