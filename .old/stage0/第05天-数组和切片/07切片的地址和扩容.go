package main

import "fmt"

func main() {
	slice := make([]int, 10, 20)
	fmt.Printf("%T", slice)
	fmt.Println(slice)
	slice = append(slice, 78)
	fmt.Println(len(slice))
	fmt.Println(slice)
}
