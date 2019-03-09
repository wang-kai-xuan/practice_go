package main

import "fmt"

func demo12() (res []int) {
	res = []int{1111, 2, 3, 4, 5}
	return
}
func main() {
	var arr_slice []int = []int{1, 2, 3, 4, 5}
	fmt.Println(arr_slice)
	arr_slice[0] = 123
	arr_slice[2] = 456
	fmt.Println(arr_slice)
	arr_slice = append(arr_slice, 6)
	fmt.Println(arr_slice)
	fmt.Printf("%T", demo12())

}
