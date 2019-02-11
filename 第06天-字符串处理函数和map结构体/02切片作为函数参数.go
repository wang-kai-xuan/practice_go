package main

import "fmt"

func test0601(s []int) []int {
	s = append(s, 9)
	fmt.Printf("%p\n", s)
	return s
}
func main() {
	arr := []int{12, 34, 44}
	fmt.Println(arr)
	fmt.Printf("%p\n", arr)
	res := test0601(arr)
	fmt.Println(res)
	fmt.Printf("%p\n", res)
}
