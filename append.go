package main

import "fmt"

func main() {
	x := []int{12, 2, 3}
	y := []int{2, 12, 31}
	x = append(x, y...)
	fmt.Println(x)
}
