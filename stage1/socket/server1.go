package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("listen err:", err)
		return
	}
	defer listener.Close()

	fmt.Println("服务器等待client链接：")
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("accept err:", err)
		return
	}
	defer conn.Close()

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("read err:", err)
		return
	}
	fmt.Println("client 发送的数据是：", string(buf[:n]))
}
