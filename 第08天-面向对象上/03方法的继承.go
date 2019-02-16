package main

import "fmt"

type person7 struct {
	id   int
	name string
	age  int
	sex  string
}

type student7 struct {
	person7

	class int
	score int
	addr  string
}

func (p *person7) Print() {
	fmt.Printf("编号：%d\n", p.id)
	fmt.Printf("姓名：%s\n", p.name)
	fmt.Printf("年龄：%d\n", p.age)
	fmt.Printf("性别：%s\n", p.sex)

}
func (p *student7) Print1() {
	p.Print()
	fmt.Printf("班级：%d\n", p.class)
	fmt.Printf("分数：%d\n", p.score)
	fmt.Printf("地址：%s\n", p.addr)

}
func main() {
	p := person7{100, "wangkaixuan", 24, "男"}
	p.Print()
	fmt.Println()
	s := student7{p, 1, 100, "遵义"}
	s.Print1()
}
