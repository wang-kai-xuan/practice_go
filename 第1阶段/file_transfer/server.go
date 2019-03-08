package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func recvFile(fileName string, conn net.Conn) {
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("os.Create err", err)
		return
	}
	defer f.Close()

	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		fmt.Println(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件接受完毕")
				return
			} else {
				fmt.Println("conn.Read err", err)
				return
			}
			return
		}
		f.Write(buf[:n])
	}

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
	conn, err := lister.Accept()
	if err != nil {
		fmt.Println("net.Accept err", err)
		return
	}
	defer conn.Close()

	// 读取client发送的文件名
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err", err)
		return
	}

	// 保存文件名
	fileName := string(buf[:n])
	conn.Write([]byte("ok"))
	// 接受文件内容
	recvFile(fileName, conn)
}
