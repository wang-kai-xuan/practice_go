package main

import "fmt"

func add0401(a int, b int) (sum int) {
	sum = a + b
	return
}
func main() {
	fmt.Println(add0401(7, 8))
}
