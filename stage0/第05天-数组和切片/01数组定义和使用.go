package main

import "fmt"

func main011() {
	var arr [10]int
	arr[0] = 999
	arr[1] = 99
	arr[2] = 919
	arr[3] = 9
	arr[4] = 9909
	fmt.Println(len(arr))
	fmt.Println(arr)
}
func main() {
	arr := [10]int{1, 2, 4, 5, 22, 23, 445, 452}
	fmt.Println(len(arr))
	fmt.Println(arr)
}
