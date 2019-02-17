package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	fp, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("打开文件出错")
		return
	}
	defer fp.Close()
	r := bufio.NewReader(fp)
	for {
		str, err := r.ReadString('\n')
		fmt.Println(str)
		if err == io.EOF {
			break
		}
	}
}
