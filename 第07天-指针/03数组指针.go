package main

import "fmt"

func main() {
	arr := [3]int{12, 32, 43}
	p := &arr
	fmt.Println(arr, *p)
	fmt.Println(arr, p)
	fmt.Println(len(*p))
	fmt.Println(len(p))
	fmt.Println((*p)[1])
	fmt.Println(p[1]) // TODO 为何访问到的数据是一样的呢, 因为go语言支持这种操作
}
