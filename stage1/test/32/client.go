package main

import (
	"fmt"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:8005")
	if err != nil {
		fmt.Println("net.Dial err", err)
		return
	}
	defer conn.Close()

	// 发送文件信息给server
	_, err = conn.Write([]byte("this is a socket test"))
	if err != nil {
		fmt.Println("conn.Write err", err)
		return
	}
	buf := make([]byte, 1024)
	n, _ := conn.Read(buf)
	fmt.Println(string(buf[:n]))
}
