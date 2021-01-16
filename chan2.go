package main

import "fmt"

// 验证：写已经关闭的通道会导致程序崩溃
func main() {
	chan1 := make(chan int)
	go func() {
		chan1 <- 8
	}()
	fmt.Println("value1:", <-chan1)
	close(chan1)
	chan1 <- 9
}
