package main

import "fmt"

type Humaner0 interface {
	SayHello()
}
type Humaner1 interface {
	Humaner0
	SayHello1()
}

type Student0 struct {
	name string
	age  int
	sex  string
}
type Teacher0 struct {
	name string
	age  int
	sex  string
	sub  string
}

func (h *Teacher0) SayHello1() {
	fmt.Println("继承的", *h)
}
func (h *Student0) SayHello1() {
	fmt.Println("继承的", *h)
}
func (h *Teacher0) SayHello() {
	fmt.Println(*h)
}
func (h *Student0) SayHello() {
	fmt.Println(*h)
}

func SayHello1(h Humaner1) {
	h.SayHello()
	h.SayHello1()

}
func main() {
	s := Student0{"王凯旋", 24, "男"}
	t := Teacher0{"王凯旋", 24, "男", "计算机"}
	SayHello1(&s)
	SayHello1(&t)
}
