package singleton

import "sync"

var lock = &sync.Mutex{}

type single struct{}

var singleInstance *single

func getInstance(isCreatedCh chan<- bool) *single {

	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			isCreatedCh <- false
			singleInstance = &single{}
		} else {
			isCreatedCh <- true
		}
	} else {
		isCreatedCh <- true
	}
	return singleInstance
}
