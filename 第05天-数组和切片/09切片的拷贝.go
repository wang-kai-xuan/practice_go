package main

import "fmt"

func main() {
	var arr_slice []int = []int{1, 2, 3, 4, 5}
	arr_copy := make([]int, 5)
	copy(arr_copy, arr_slice)
	fmt.Println(arr_copy)
	fmt.Println(len(arr_copy))
	fmt.Printf("%p\n", &arr_copy)
	fmt.Println(arr_slice)
	fmt.Printf("%p\n", &arr_slice)

}
