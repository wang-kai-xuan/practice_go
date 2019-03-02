package main

import (
	"fmt"
	"runtime"
	"time"
)

var channel = make(chan int)

func printer(s string) {
	for _, ch := range s {
		fmt.Printf("%c", ch)
		time.Sleep(200 * time.Millisecond)
	}
}
func user2() {
	printer("--user2 use--")
	channel <- 8
}
func user1() {
	printer("--user1 use--")
}
func main() {

	// channel可以用来操作的同步
	go user2()
	<-channel // 等待数据，无数据时会阻塞
	go user1()
	runtime.Goexit()
}
