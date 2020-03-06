package main

import "fmt"

func main() {
	var i []interface{}
	i = append(i, 10, 10.2, "wkx")
	fmt.Println(i)

	// 类型断言，输出字符串
	for _, value := range i {
		if data, ok := value.(string); ok {
			fmt.Println(data)
		}
	}
}
