package main

import (
	"fmt"
	"time"
)

// 这个demo的主要目的是为了说明：
// 1、两个goroutine使用无缓冲的通道进行通信
//  将导致两个goroutine执行流程中的某一个节点（时间、事件）同步化
func main() {
	chan1 := make(chan int)
	go func() {
		num := 0
		for {
			// 向chan1写入的数据被读取/接受后，才能继续执行
			// chan1 <- num语句执行过程分为3个阶段: 发送数据->等待读取->执行下一个语句
			chan1 <- num
			fmt.Println("child send:", num)
			num++
			fmt.Println("child wait")
			time.Sleep(time.Millisecond * 2)
		}
	}()

	for {
		fmt.Println("main wait")
		time.Sleep(time.Millisecond * 2)
		// chan1的数据读取后，才能继续执行
		num := <-chan1
		if num > 10 {
			return
		}
		fmt.Println("main recv:", num)
	}
}
