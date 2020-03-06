package main

import "fmt"

func main() {
	a := 10
	b := 30
	defer func(a int, b int) {
		fmt.Println(a + b)
	}(a, b)
	a = 1
	b = 2
	defer func() {
		fmt.Println(a + b)
	}()
	fmt.Println(a - b)
}
