package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	cli, err := rpc.DialHTTP("tcp", "127.0.0.1:10086")
	if err != nil {
		fmt.Println(err)
	}

	var data int
	err = cli.Call("Demo1.GetResult", 101, &data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
}
