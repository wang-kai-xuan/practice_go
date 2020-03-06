package main

import "fmt"

func main() {
	var arr_slice []int = []int{1, 2, 3, 4, 5}
	fmt.Println(arr_slice[:])
	fmt.Println(arr_slice[:2])
	fmt.Println(arr_slice[2:])
}
