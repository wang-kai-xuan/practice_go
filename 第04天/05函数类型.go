package main

import "fmt"

type FUNC_ADD func(int, int) int

func add0501(a int, b int) (sum int) {
	sum = a + b
	return
}

func main() {
	var func_add FUNC_ADD
	func_add = add0501
	fmt.Println(func_add(9, 90))
}
