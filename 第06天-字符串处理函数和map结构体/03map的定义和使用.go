package main

import "fmt"

func main() {
	m := map[string]string{"name": "王凯旋", "age": "24", "gender": "男"}
	fmt.Println(m)
	m["name"] = "陈刚容"
	fmt.Println(m)
	delete(m, "name")
	fmt.Println(m)
}
