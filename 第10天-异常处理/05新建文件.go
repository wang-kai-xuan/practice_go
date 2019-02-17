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
	defer fp.Close()
	fp.WriteString("我叫王凯旋")
}
