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
	fmt.Println(runtime.GOROOT())
	fmt.Println(runtime.NumCPU())
	n := runtime.GOMAXPROCS(2)
	fmt.Println(n)

	for i := 0; i < 10; i++ {
		go fmt.Println(0)
		fmt.Println(1)
	}
}
