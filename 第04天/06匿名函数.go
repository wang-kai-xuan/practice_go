package main

import "fmt"

func main() {
	f := func(a int, b int) int {
		return a + b
	}
	fmt.Println(f(9, 19))
}
