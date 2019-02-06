package main

import "fmt"

func main0301() {
	var score int
	fmt.Scan(&score)

	switch score / 10 {
	case 70:
		fallthrough
	case 60:
		fmt.Println("可以上复旦")
	case 50:
		fmt.Println("可以上复旦吗")
	}
}
func main() {
	var score int
	fmt.Scan(&score)
	switch score / 10 {
	case 10:
		fallthrough
		// TODO 需要深入理解其用法
	case 9:
		fmt.Println("A")
	case 8:
		fmt.Println("B")
	case 7:
		fmt.Println("C")
	case 6:
		fmt.Println("D")
	default:
		fmt.Println("E")
	}
}
