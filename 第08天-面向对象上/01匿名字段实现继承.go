package main

import "fmt"

type Person struct {
	id   int
	name string
	age  int
}
type Addr struct {
	country string
	city    string
}
type Student struct {
	*Person
	Addr
	name  string
	class int
	score int
}

func main() {

	stu := Student{&Person{1, "wangkaixuan", 24}, Addr{"中国", "遵义"}, "wkx", 1, 100}
	fmt.Println(stu)
	stu.name = "chengangrong"
	fmt.Println(stu)

	fmt.Printf("%d\n", stu.id)
	fmt.Printf("%s\n", stu.name)
	fmt.Printf("%s\n", stu.Person.name)
	fmt.Printf("%d\n", stu.age)
	fmt.Printf("%d\n", stu.class)
	fmt.Printf("%d\n", stu.score)
	fmt.Printf("%s\n", stu.country)
	fmt.Printf("%s\n", stu.city)
}
