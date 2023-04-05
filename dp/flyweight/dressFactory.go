package flyweight

import (
	"fmt"
	"sync"
)

var dressFactorySingleInstance *DressFactory

type DressFactory struct {
	dressMap map[PlayerType]IDress
}

func (d *DressFactory) getDressByType(playerType PlayerType) (IDress, error) {
	// すでにInstanceがある場合
	if d.dressMap[playerType] != nil {
		return d.dressMap[playerType], nil
	}

	// ない場合、新規作成
	if playerType == T {
		d.dressMap[playerType] = newTDress()
		return d.dressMap[T], nil
	}
	if playerType == CT {
		d.dressMap[playerType] = newCTDress()
		return d.dressMap[CT], nil
	}

	return nil, fmt.Errorf("Wrong player type passed")
}

func getDressFactorySingleInstance() *DressFactory {
	var lock = &sync.Mutex{}
	if dressFactorySingleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if dressFactorySingleInstance == nil {
			dressFactorySingleInstance = &DressFactory{
				dressMap: make(map[PlayerType]IDress),
			}
		}
	}
	return dressFactorySingleInstance
}
