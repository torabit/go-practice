package observer

type Item struct {
	observerList []IObserver
	name         string
	inStock      bool
}

func newItem(name string) *Item {
	return &Item{
		name: name,
	}
}

func (i *Item) updateAvailability() {
	i.inStock = true
	i.notifyAll()
}

func (i *Item) register(o IObserver) {
	i.observerList = append(i.observerList, o)
}

func (i *Item) deregister(o IObserver) {
	i.observerList = removeFromslice(i.observerList, o)
}

func (i *Item) notifyAll() {
	for _, observer := range i.observerList {
		observer.update(i.name)
	}
}

func removeFromslice(observerList []IObserver, observerToRemove IObserver) []IObserver {
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if observerToRemove.getID() == observer.getID() {
			// 削除したいovserverをSliceの最後と入れ替える
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			// Sliceの最後以外を返す
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}
