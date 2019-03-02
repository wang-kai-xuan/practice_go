package main

import (
	"fmt"
	"runtime"
)

func song() {
	for i := 0; i < 100; i++ {
		fmt.Println("---song---")
	}
}
func dance() {
	for i := 0; i < 100; i++ {
		fmt.Println("---dance---")
	}
}
func main() {
	go func() {
		defer fmt.Println("---in defer---")
		fmt.Println("---goroutine 0---")
		fmt.Println("---goroutine 1---")
		runtime.Goexit()
		fmt.Println("---goroutine 2---")
		fmt.Println("---goroutine 3---")
	}()
	for {

	}
}
