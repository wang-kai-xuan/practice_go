package main

import "fmt"

func add05(a int, b int) {
	fmt.Println(a + b)
}
func add04(arr ...int) {
	add05(arr[0], arr[1])
}
func main() {
	add04(1, 2, 3, 4, 5, 5)
}
