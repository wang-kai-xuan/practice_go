package main

import "fmt"

func add0401(a int, b int) (sum int) {
	sum1 := a + b
	fmt.Println(sum1)
	return
}
func main() {
	fmt.Println(add0401(7, 8))
}
