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

func TestMap(t *testing.T) {
	// test content:
	// init,write, read
	m := map[string]string{"name": "王凯旋", "age": "24", "gender": "男"}
	fmt.Println(m)
	m["name"] = "陈刚容"
	fmt.Println(m)
	delete(m, "name")
	fmt.Println(m)
}
func TestMap1(t *testing.T) {
	// test content:
	// init,write, read
	m := make(map[string]string)
	m["cgr"] = "陈刚容"
	fmt.Println(m)
	m["wkx"] = "王凯旋"
	fmt.Println(m)
	delete(m, "wkx")
	fmt.Println(m)
}

func TestMap2(t *testing.T) {
	// test content:
	// read not exist val
	m := make(map[string]string)
	m["wkx"] = "王凯旋"
	name := m["wkx"]
	fmt.Println(name)
	name1 := m["wkx1"]
	if name1 == "" {
		fmt.Println("name not exist")
	} else {
		fmt.Println("name exist")
	}
	fmt.Println("end")
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

// 问题：如何获取interface类型变量的成员
// 回答：fmt.Sprintf
func TestInterfaceMember(t *testing.T) {
	var i interface{}
	i = 1
	i = "nihao"

	i1 := fmt.Sprintf("%s", i)
	i = true
	fmt.Println(i)
	fmt.Println(i1)
}

func TestSet(t *testing.T) {
	// 初始化
	s := New()

	s.Add(1)
	s.Add(1)
	s.Add(2)
	fmt.Println("list of all items", s.List())

	s.Clear()
	if s.IsEmpty() {
		fmt.Println("0 item")
	}

	s.Add(1)
	s.Add(2)
	s.Add(3)
	fmt.Println("list of all items", s.List())

	if s.Has(2) {
		fmt.Println("2 does exist")
	}

	s.Remove(2)
	s.Remove(3)
	fmt.Println("list of all items", s.List())
}
