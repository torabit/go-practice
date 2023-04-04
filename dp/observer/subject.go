package observer

type ISubject interface {
	register(IObserver)
	deregister(IObserver)
	notifyAll()
}
