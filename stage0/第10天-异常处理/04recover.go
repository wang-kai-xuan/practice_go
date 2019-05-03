package main

import "fmt"

func demo() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	//var arr [10] int
	//arr[10] = 19

	var p *int
	*p = 123
}
func main() {
	demo()
}
