package main

import (
	"fmt"
	"practice_go/microServices/proto"
)

func main() {
	text := &proto.Demo1{
		Name:   "wangkaixuan",
		Weight: []int32{150, 140},
		Height: 175,
		Memo:   "hello world!",
	}

	fmt.Println(text)
}
