package main

import (
	"errors"
	"fmt"
)

/**
如何抛出异常呢？
如何捕获异常呢？
*/
func div(a int, b int) (value int, err error) {
	if b == 0 {
		err = errors.New("除数不能为0")
		return
	}
	value = a / b
	return
}
func main() {
	a := 10
	b := 5
	val, err := div(a, b)
	if err != nil {
		fmt.Println("捕获到错误了")
	} else {
		fmt.Println(val)
	}

}
