package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 5, 6, 7}
	p := &slice

	for i := 0; i < len(*p); i++ {
		fmt.Println((*p)[i])
	}
}
