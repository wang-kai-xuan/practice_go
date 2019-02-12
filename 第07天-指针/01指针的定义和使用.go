package main

import "fmt"

func main0701() {
	a := 10
	ap := &a
	fmt.Println(a)
	*ap = 20
	fmt.Println(a)
}
func main() {
	var a int = 20
	var b *int = &a
	fmt.Println(a)
	*b = 30
	fmt.Println(a)
}
