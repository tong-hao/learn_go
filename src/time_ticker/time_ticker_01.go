package main

import (
	"fmt"
	"time"
)

// 创建一个计时器
var timeTickerChan = time.Tick(time.Second * 2)

func doTask() {
	for {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
		<-timeTickerChan
	}
}

func main() {
	go doTask()
	time.Sleep(10 * time.Second)
}
