package main

import (
	"fmt"
	"sync"
)
import "time"

func hello() {
	fmt.Println("hello2")
}

var wg sync.WaitGroup

func hello2() {
	fmt.Println("hello2")
	defer wg.Done() // -1
}

func main() {
	// 1. sleep
	{
		go hello()
		fmt.Println("01")
		time.Sleep(1 * time.Second)
	}

	// 2. sync.WaitGroup
	{
		wg.Add(1) // +1
		go hello2()
		wg.Wait() // 阻塞，直到为0
	}

	// 3. channel
	{
		var a chan int
		a = make(chan int, 10) // 缓冲大小为10
		a <- 10                // 把10发送给a通道
		x := <-a               // 从a通道取一个值
		fmt.Printf("x=%v\n", x)

		a <- 11
		a <- 12
		close(a) // 关闭通道(先关闭再range)
		for i := range a {
			fmt.Printf("i=%v\n", i)
		}

	}

	// 4. select
	{
		a := make(chan int, 10)
		for i := 0; i < 10; i++ {
			select {
			case x := <-a: // 接收
				fmt.Println(x)
			case a <- i: // 发送
			}
		}
	}


}
