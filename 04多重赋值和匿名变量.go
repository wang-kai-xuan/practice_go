package main

import "fmt"

func main0401() {
	a,b,c,d := 10,3.14,"王凯旋",true
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
func main() {
	main0401()
	_,_,c,d := 10,3.14,"王凯旋",true
	//fmt.Println(a)
	//fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}