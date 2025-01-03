package main

import (
	"fmt"
	"sync"
	"time"
)

func countdown(seconds *int, mutex *sync.Mutex) {
	for {
		mutex.Lock()
		if *seconds <= 0 {
			mutex.Unlock()
			break
		}
		time.Sleep(1 * time.Second)
		*seconds -= 1
		mutex.Unlock()
	}
}

func main() {
	mutex := sync.Mutex{}
	count := 5
	go countdown(&count, &mutex)
	for {
		mutex.Lock()
		if count <= 0 {
			break
		}

		time.Sleep(500 * time.Millisecond)
		fmt.Println(count)
		mutex.Unlock()
	}
}
