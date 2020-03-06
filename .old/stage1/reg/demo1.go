package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "wangkaixuan 47 wkx abc wxx 3.14 5.11"
	reg := regexp.MustCompile(`w.x`)
	res := reg.FindAllStringSubmatch(str, -1)
	fmt.Println(res)
}
