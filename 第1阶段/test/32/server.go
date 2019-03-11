package main

import (
	"fmt"
	"net"
	"strings"
)

func handle(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err", err)
		return
	}
	data := string(buf[:n])
	str := strings.Split(data, " ")
	var result []string
	for i := len(str) - 1; i >= 0; i-- {
		result = append(result, str[i])
	}
	conn.Write([]byte(strings.Join(result, " ")))
}
func main() {
	// 创建tcp监听
	lister, err := net.Listen("tcp", "127.0.0.1:8005")
	if err != nil {
		fmt.Println("net.Listen err", err)
		return
	}
	defer lister.Close()

	// 等待client连接
	for {
		conn, err := lister.Accept()
		if err != nil {
			fmt.Println("net.Accept err", err)
			return
		}
		go handle(conn)
	}
}
