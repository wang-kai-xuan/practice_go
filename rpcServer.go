package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
)

type Demo1 int

func (demo *Demo1) GetResult(req int, res *int) error {
	fmt.Println("recv data:", req)
	*res = req + 100
	return nil
}
func main() {
	demo1 := new(Demo1)
	rpc.Register(demo1)
	rpc.HandleHTTP()

	ln, err := net.Listen("tcp", ":10086")
	if err != nil {
		fmt.Println(err)
	}
	http.Serve(ln, nil)
}
