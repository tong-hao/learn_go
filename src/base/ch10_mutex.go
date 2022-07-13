package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var v int

func add(m *sync.Mutex) {
	for i := 0; i < 1000; i++ {
		m.Lock()
		v++
		m.Unlock()
	}
	wg.Done()
}

func add2(m *sync.RWMutex) {
	for i := 0; i < 1000; i++ {
		m.Lock()
		v++
		m.Unlock()
	}
	wg.Done()
}

func read(m *sync.RWMutex)  {
	defer wg.Done()
	m.RLock()
	fmt.Println("read, v=", v)
	m.RUnlock()

}

func main() {
	// mutex
	{
		var m sync.Mutex

		wg.Add(2)
		go add(&m)
		go add(&m)
		wg.Wait()

		fmt.Println("v=",v)
	}

	//
	{
		var m sync.RWMutex
		wg.Add(2)
		go add2(&m)
		go add2(&m)

		for i:=0;i<10;i++ {
			wg.Add(1)
			go read(&m)
			time.Sleep(time.Microsecond)
		}
		wg.Wait()
		fmt.Println("v=",v)
	}

}
