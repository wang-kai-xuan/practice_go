package main

import "fmt"

func main() {
	a := 10.0
	b := 123.456
	c := "王凯旋"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	a, b = b, a

	fmt.Println(a)
	fmt.Println(b)
}
