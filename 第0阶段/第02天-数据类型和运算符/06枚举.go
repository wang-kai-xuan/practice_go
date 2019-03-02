package main

import "fmt"

func main() {
	const (
		a = iota
		b
		c
		d
		e
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}
