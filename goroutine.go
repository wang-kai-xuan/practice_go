package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	num := 0

	go func() {
		for {
			ch <- 1
		}
	}()

	go func() {
		for {
			num += <-ch
		}
	}()

	time.Sleep(time.Second)
	fmt.Println("num:", num)
}
