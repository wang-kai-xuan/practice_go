package main

import "fmt"

/**
计算菜价格
*/
func main() {
	price := 3.2
	var weight float64

	fmt.Scan(&weight)
	sum := price * weight
	fmt.Println(sum)
}
