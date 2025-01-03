package main

import "sync"

type ReadWriteMutex struct {
	readersCounter int
	readersLock    sync.Mutex
	globalLock     sync.Mutex
}

func (rw *ReadWriteMutex) ReadLock() {
	rw.readersLock.Lock()
	rw.readersCounter++
	if rw.readersCounter == 1 {
		rw.globalLock.Lock()
	}

	rw.readersLock.Unlock()
}

func (rw *ReadWriteMutex) WriteLock() {
	rw.globalLock.Lock()
}

func (rw *ReadWriteMutex) ReadUnlock() {
	rw.readersLock.Lock()
	rw.readersCounter--
	if rw.readersCounter == 0 {
		rw.globalLock.Unlock()
	}
	rw.readersLock.Unlock()
}

func (rw *ReadWriteMutex) WriteUnlock() {
	rw.globalLock.Unlock()
}

func (rw *ReadWriteMutex) TryLock() bool {
	return rw.globalLock.TryLock()
}

func (rw *ReadWriteMutex) TryReadLock() bool {
	rw.readersLock.Lock()
	rw.readersCounter++
	if rw.readersCounter == 1 {
		return rw.globalLock.TryLock()
	}

	rw.readersLock.Unlock()
	return true
}
