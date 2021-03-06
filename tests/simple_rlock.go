package tests

import (
	"sync"
)

type another struct {
	m sync.RWMutex
}

func (a *another) Test() {
	a.m.RLock()
	defer a.m.RUnlock()

	a.m.Lock() // want "Mutex lock is acquired on this line"
	a.m.Unlock()
}
