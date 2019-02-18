package main

import "fmt"

type Person struct {
	name string
	sex  byte
	int
}

func main() {
	array := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := array[:6:8]
	fmt.Println(len(s1))
	fmt.Println(cap(s1))
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(s)
}
