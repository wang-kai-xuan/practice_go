package main

import "fmt"

func main0701() {
	var a int
	fmt.Scan(&a)
	a += 100
	fmt.Println(a)
}
func main0702() {
	var a, b int
	fmt.Scan(&a, &b)
	a += 10
	b += 10
	fmt.Println(a)
	fmt.Println(b)
}
func main0703() {
	var a int
	var b string

	fmt.Scanf("%d%s", &a, &b)
	fmt.Println(a)
	fmt.Printf("==%s==", b)
}
func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	sum := a + b + c
	fmt.Println("总成绩是:", sum)
	fmt.Println("平均成绩是：", sum/3)
}
