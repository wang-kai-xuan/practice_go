package main

import "fmt"

func main0301() {
	ch := 'a'
	str := "a"

	fmt.Printf("%c\n", ch)
	fmt.Printf("%s\n", str)
}
func main0302() {
	a := "王凯旋"
	var count int
	count = len(a)
	fmt.Println(count)
}
func main() {
	a := "王凯旋"
	b := "正在学习go语言"
	str3 := a + b
	fmt.Println(str3)
}
