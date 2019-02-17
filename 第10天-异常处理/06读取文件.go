package main

import (
	"fmt"
	"os"
)

func main() {
	fp, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("打开文件出错")
		return
	}
	defer fp.Close()
	buf := make([]byte, 20)
	fp.Read(buf)
	fmt.Println(string(buf))

}
