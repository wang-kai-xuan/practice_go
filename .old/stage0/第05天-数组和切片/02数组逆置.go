package main

import "fmt"

func main() {
	var arr [7]int = [7]int{23, 3, 45, 6, 56, 7, 34}
	left := 0
	right := len(arr) - 1
	fmt.Println(arr)
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
	fmt.Println(arr)
}
