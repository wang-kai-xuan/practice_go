package main

import "fmt"

func swap(a *int, b *int) {
	*a, *b = *b, *a
}
func main() {
	a := 10
	b := 20
	fmt.Println(a, b)
	swap(&a, &b)
	fmt.Println(a, b)
}
