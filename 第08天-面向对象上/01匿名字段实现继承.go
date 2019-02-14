package main

import "fmt"

type Person struct {
	id   int
	name string
	age  int
}
type Student struct {
	*Person
	name  string
	class int
	score int
}

func main() {

	stu := Student{&Person{1, "wangkaixuan", 24}, "wkx", 1, 100}
	fmt.Println(stu)
	stu.name = "chengangrong"
	fmt.Println(stu)

	fmt.Printf("%d\n", stu.id)
	fmt.Printf("%s\n", stu.name)
	fmt.Printf("%s\n", stu.Person.name)
	fmt.Printf("%d\n", stu.age)
	fmt.Printf("%d\n", stu.class)
	fmt.Printf("%d\n", stu.score)
}
