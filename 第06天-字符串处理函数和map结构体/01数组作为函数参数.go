package main

import (
	"fmt"
)

func swap(arr [10]int) [10]int {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
func main() {
	arr := [10]int{10, 3, 2, 1, 4, 5, 6, 7, 8, 9}
	arr1 := [10]int{10, 3, 2, 1, 4, 5, 6, 7, 8, 9}
	fmt.Println(arr == arr1)
	fmt.Println(arr)
	fmt.Println(swap(arr))
	fmt.Println(arr)
	//var i int
	i := 1
	fmt.Println(i)
}
