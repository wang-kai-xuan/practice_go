package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

type Person struct {
	Name string
	Age  uint
}

func main() {
	p := Person{"wangkaixuan", 25}
	var buffer bytes.Buffer

	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(&p)
	if err != nil {
		fmt.Println(err)
		log.Panic(" encode error!")
	}

	fmt.Printf("encode:%v\n", buffer.Bytes())

	decoder := gob.NewDecoder(bytes.NewReader(buffer.Bytes()))
	var me Person
	err = decoder.Decode(&me)
	if err != nil {
		log.Panic("decode error!")
	}
	fmt.Printf("decode:%v\n", me)
}
