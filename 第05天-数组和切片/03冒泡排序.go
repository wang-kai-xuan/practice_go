package main

import "fmt"

func main() {
	var arr [10]int = [10]int{9, 23, 20, 33, 11, 8, 12, 213, 98, 88}
	fmt.Println(arr)
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println(arr)
}
