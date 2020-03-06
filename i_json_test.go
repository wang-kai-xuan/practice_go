package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

type person struct {
	Name string
	Age  int
}

// byte转字符串
func TestJson(t *testing.T) {
	p := person{Name: "wangkaixuan", Age: 25}
	var r interface{}
	r, e := json.Marshal(p)
	if e != nil {
		fmt.Println(e.Error())

	}

	fmt.Printf("%s\n", r)
	fmt.Printf("%v\n", r)
	r1 := fmt.Sprintf("%s\n", r)
	fmt.Printf("%s\n", r1)
}
