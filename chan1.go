package main

import "fmt"

// 验证：读已经关闭的通道会读取到零值
func main() {
	chan1 := make(chan int)
	go func() {
		chan1 <- 8
	}()
	fmt.Println("value1:", <-chan1)
	close(chan1)
	fmt.Println("value2:", <-chan1)
}
