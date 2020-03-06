package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("Dial err:", err)
		return
	}
	defer conn.Close()

	_, err = conn.Write([]byte("我是王凯旋==="))
	if err != nil {
		fmt.Println("write err:", err)
		return
	}
}
