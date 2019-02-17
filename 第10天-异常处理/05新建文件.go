package main

import (
	"fmt"
	"os"
)

func main() {
	fp, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("创建文件失败")
		return
	}
	fp.Close()
}
