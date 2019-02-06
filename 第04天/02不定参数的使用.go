package main

import "fmt"

func add02(arr ...int) {
	var sum int
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	fmt.Println(sum)
}

func add03(arr ...int) {
	sum := 0
	for index, value := range arr {
		fmt.Println(index, value)
		sum += value
	}
	fmt.Println(sum)
}
func main() {
	add02(1, 2, 3, 4)
	add02(11, 2, 3, 4)
	add02(1, 22, 3, 4)
	add03(1, 22, 3, 4)
}
