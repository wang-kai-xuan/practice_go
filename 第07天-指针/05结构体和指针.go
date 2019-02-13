package main

import "fmt"

type Person struct {
	name string
	age  int
	add  string
}

func main() {
	wkx := Person{"wkx", 24, "遵义"}
	cgr := Person{"cgr", 23, "遵义"}
	p := &wkx
	fmt.Println(*p)
	p = &cgr
	fmt.Println(*p)
	fmt.Println((*p).add)
}
