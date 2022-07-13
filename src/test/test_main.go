package main

import (
	"fmt"
	"sync"
	"time"
)

func test(lock *sync.Mutex, idx int) {
	lock.Lock()
	defer lock.Unlock()
	fmt.Println("begin:", idx)
	time.Sleep(time.Second)
	fmt.Println("end:", idx)
}

func main() {

	locks := make([]sync.Mutex, 3)

	for i := 0; i < 1024; i++ {
		go test(&locks[i%3], i%3)
	}

	time.Sleep(10 * time.Second)

}
