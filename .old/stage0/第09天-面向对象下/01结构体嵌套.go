package main

import "fmt"

type student struct {
	*student
	name string
	age  int
	sex  string
}

func main() {
	s := student{&student{nil, "wkx", 24, "男"}, "cgr", 23, "女"}
	fmt.Println(s)
}
