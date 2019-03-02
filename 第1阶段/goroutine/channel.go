package main

import (
	"fmt"
	"time"
)

var channel = make(chan int, 5)

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
	go func() {
		for i := 0; i < 15; i++ {
			channel <- i
			fmt.Println("-- in child---: i=", i, " len=", len(channel), " cap=", cap(channel))
		}
	}()
	for i := 0; i < 15; i++ {
		num := <-channel
		fmt.Println("-- in main---: num=", num)
	}

}
