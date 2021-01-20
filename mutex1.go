package main

import (
	"fmt"
	"time"
)

func main() {
	// ch := make(chan struct{})
	var x, y int
	// 因为并发执行，打印出来的x,y的值，是不定的
	/*
	   x:0y:1x:1y:1x:1y:1x:1y:1x:1y:1x:1y:1x:1y:1x:1y:1x:1y:1y:1x:1
	   x:0y:1x:1y:1x:1y:1x:1y:1x:1y:1x:1y:1x:1y:1x:1y:1x:1y:1x:1y:1
	   x:0y:1x:1y:1x:1y:1x:1y:1x:1y:1x:1y:1x:1y:1x:1y:1x:1y:1x:1y:1
	   x:1x:1y:0x:1x:1y:1y:1x:1y:0y:1y:1x:1x:1x:1x:1x:1y:1y:1y:1y:1
	   x:0y:1x:1y:1x:1y:1x:1y:1x:1y:1x:1y:1x:1y:1x:1y:1x:1y:1x:1y:1
	   x:0y:1x:1y:1x:1y:1x:1y:1x:1y:1x:1y:1x:1y:1x:1y:1x:1y:1x:1y:1
	   x:1x:1y:1x:1x:1y:1x:1y:1x:1x:0y:1y:1y:1x:0y:1y:1y:1x:0y:1x:1
	   x:0y:1x:1y:1x:1y:1x:1y:1x:1y:1x:1y:1x:1y:1x:1y:1x:1y:1y:1x:1
	   x:0y:1x:1y:1x:1y:1x:1y:1x:1y:1x:1y:1x:1y:1x:1y:1x:1y:1x:1y:1
	   x:0y:1x:1y:1x:1y:1x:1y:1x:1y:1x:1y:1x:1y:1x:1y:1x:1y:1x:1y:1
	   x:0y:1x:1y:1x:1y:1x:1y:1x:1y:1x:1y:1x:1y:1x:1y:1x:1y:1x:1y:1
	*/
	for i := 0; i < 10; i++ {
		go func() {
			x = 1                  // Al
			fmt.Print("y:", y, "") // A2
			// ch <- struct{}{}
		}()
		go func() {
			y = 1                  // Bl
			fmt.Print("x:", x, "") // B2
			// ch <- struct{}{}
		}()

	}
	// <-ch
	// <-ch
	time.Sleep(time.Second)
	fmt.Print("\n")
}
