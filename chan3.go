package main

import (
	"fmt"
	"time"
)

// 验证：缓冲通道在通道满后，对其写操作会阻塞，直到通道数据被读取一个
// 通道写入数据1ms一个元素
// 通道读取速度2s一个元素
// 一定时间后，写入数据速度和读取速度将"同步"
func main() {
	// 通道大小为5
	chan1 := make(chan int, 5)
	go func() {
		num := 0
		for {
			// 当程序退出时,通道内还存在5个元素,由此可得最后发送的数据是 6+5=11
			chan1 <- num
			fmt.Println("have been send:", num)
			time.Sleep(time.Millisecond * 1)
			num++
		}
	}()

	recvNum := 0
	for {
		recvNum = <-chan1
		fmt.Println("have been recv:", recvNum)
		time.Sleep(time.Second * 1)
		// 当接收到大于等于6时,程序退出

		if recvNum >= 6 {
			return
		}
	}
}
