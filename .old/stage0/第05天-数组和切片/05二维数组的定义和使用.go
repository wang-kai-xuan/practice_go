package main

import "fmt"

func main() {
	var arr [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(arr[0][0])
	fmt.Println(arr[0][1])
	fmt.Println(arr[1][0])
	fmt.Println(arr)
}
