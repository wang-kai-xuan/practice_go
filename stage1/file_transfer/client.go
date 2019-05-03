package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func sendFile(conn net.Conn, filePath string) {
	// 打开文件
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("os.Open err", err)
		return
	}
	defer f.Close()

	// 发送完整文件内容
	buf := make([]byte, 4096)
	for {
		n, err := f.Read(buf)
		fmt.Println(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件发送完成")
			} else {
				fmt.Println("os.Read err", err)
				return
			}
			return
		}
		conn.Write(buf[:n])
	}
}
func main() {
	// 获取文件信息
	list := os.Args
	if len(list) != 2 {
		fmt.Println("usage: xxx.go file_name")
		return
	}

	fileName := list[1]
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		fmt.Println("os.Stat err", err)
		return
	}
	fmt.Println(fileInfo.Name())
	fmt.Println(fileInfo.Size())

	// 连接server
	conn, err := net.Dial("tcp", "127.0.0.1:8005")
	if err != nil {
		fmt.Println("net.Dial err", err)
		return
	}
	defer conn.Close()

	// 发送文件信息给server
	_, err = conn.Write([]byte(fileInfo.Name()))
	if err != nil {
		fmt.Println("conn.Write err", err)
		return
	}

	// 读取server确认收到消息, 判断是不是ok
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err", err)
		return
	}
	if "ok" == string(buf[:n]) {
		sendFile(conn, fileName)
	}
}
