package main

import "fmt"

type Humaner interface {
	SayHello()
}
type Student struct {
	name string
	age  int
	sex  string
}
type Teacher struct {
	name string
	age  int
	sex  string
	sub  string
}

func (h *Teacher) SayHello() {
	fmt.Println(*h)
}
func (h *Student) SayHello() {
	fmt.Println(*h)
}

func main() {
	s := Student{"王凯旋", 24, "男"}
	t := Teacher{"王凯旋", 24, "男", "计算机"}
	var h Humaner
	h = &s
	h.SayHello()
	h = &t
	h.SayHello()
}
