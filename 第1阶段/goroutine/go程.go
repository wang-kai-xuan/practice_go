package main

import "fmt"

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
	go dance()
	go song()
	for {

	}
}
