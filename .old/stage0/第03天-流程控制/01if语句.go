package main

import "fmt"

func main() {
	var score int
	fmt.Scan(&score)
	if score > 700 {
		fmt.Println("可以上清华")
	} else if score > 600 {
		fmt.Println("可以上复旦")
	} else if score > 300 {
		fmt.Println("可以上蓝翔")
	} else {
		fmt.Println("可以去新东方")
	}
}
