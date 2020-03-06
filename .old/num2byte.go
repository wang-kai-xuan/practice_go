package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
)

func Uint64ToByte(num uint64) []byte {
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buffer.Bytes()
}
func main() {
	var num uint64
	num = 15
	b := Uint64ToByte(num)
	fmt.Printf("num=%d,byte=%x\n", num, b)
}
