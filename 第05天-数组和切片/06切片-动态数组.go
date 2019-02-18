package main

import "fmt"

func main() {
	var arr_slice []int = []int{1, 2, 3, 4, 5}
	fmt.Println(arr_slice)
	arr_slice[0] = 123
	arr_slice[2] = 456
	fmt.Println(arr_slice)
	arr_slice = append(arr_slice, 6)
	fmt.Println(arr_slice)
	delete(arr_slice, 123)

}
