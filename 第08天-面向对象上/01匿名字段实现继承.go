package main

import "fmt"

type Person struct {
	id   int
	name string
	age  int
}
type Student struct {
	Person
	name  string
	class int
	score int
}

func main() {
	stu := Student{Person{1, "wangkaixuan", 24}, "wkx", 1, 100}
	fmt.Println(stu)
	stu.name = "chengangrong"
	fmt.Println(stu)
}
