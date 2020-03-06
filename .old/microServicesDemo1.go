package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"practice_go/microServices/prototext"
)

func main() {
	text := &prototext.Demo1{
		Name:   "wangkaixuan",
		Weight: []int32{150, 140},
		Height: 175,
		Memo:   "hello world!",
	}

	fmt.Println(text)

	data, err := proto.Marshal(text)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)

	newText := &prototext.Demo1{}

	err = proto.Unmarshal(data, newText)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(newText)

}
