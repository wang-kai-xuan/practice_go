package main

import "fmt"

type student struct {
	id   int
	name string
	age  int
	sex  string
}

func f(s student) {
	s.id = 10
}
func main() {
	wkx := student{1, "王凯旋", 24, "男"}
	fmt.Println(wkx)
	wkx.name = "陈刚容"
	fmt.Println(wkx)
	f(wkx)
	fmt.Println(wkx)
}
