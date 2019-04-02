package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Printf("ba%s\n", bytes.Repeat([]byte("na"), 2))
	a := bytes.Repeat([]byte("na"), 2)
	fmt.Printf("%s\n", a)
	fmt.Printf("%d\n", a)
}
