package main

import "fmt"

func main() {
	fmt.Println(1)
	fmt.Println(2)
	fmt.Println(3)
	panic("there")
	fmt.Println(4)
	fmt.Println(5)
}
