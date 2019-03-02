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
		for i := 0; i < 10; i++ {
			fmt.Println("---goroutine---")
		}
	}()
	runtime.Gosched()
	for i := 0; i < 10; i++ {
		fmt.Println("---main goroutine---")
	}
}
