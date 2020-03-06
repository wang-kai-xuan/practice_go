package main

import "fmt"

func main() {
	a, b, c := 0, 0, 0
	fmt.Scan(&a, &b, &c)

	sum := a + b + c
	fmt.Println(sum)
	fmt.Println(sum / 3)
	fmt.Printf("%.4f", float64(sum)/3)
}
