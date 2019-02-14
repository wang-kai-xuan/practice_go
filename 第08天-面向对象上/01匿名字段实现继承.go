package main

import "fmt"

type Person struct {
	id   int
	name string
	age  int
}
type Student struct {
	Person
	class int
	score int
}

func main() {
	stu := Student{Person{1, "wangkaixuan", 24}, 1, 100}
	fmt.Println(stu)
	stu.name = "chengangrong"
	fmt.Println(stu)
}
