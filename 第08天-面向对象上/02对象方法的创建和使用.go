package main

import "fmt"

type MyInt int

type Student struct {
	id   int
	name string
	age  int
}

func (s Student) print() {
	fmt.Printf("学号:%d\n", s.id)
	fmt.Printf("名字:%s\n", s.name)
	fmt.Printf("年龄:%d\n", s.age)
}
func main01() {
	var a MyInt = 2
	fmt.Println(a)
}
func main() {
	stu := Student{1, "wangkaixuan", 24}
	stu.print()
}
