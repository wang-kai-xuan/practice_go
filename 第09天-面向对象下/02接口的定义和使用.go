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

func SayHello(h Humaner) {
	h.SayHello()
}
func main() {
	s := Student{"王凯旋", 24, "男"}
	t := Teacher{"王凯旋", 24, "男", "计算机"}
	SayHello(&s)
	SayHello(&t)
}
