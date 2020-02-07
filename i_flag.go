package main

import (
	"flag"
	"fmt"
	"strings"
)

var port = flag.Int("port",0,"port number")
var ip = flag.String("ip","127.0.0.1","ip addresss")

func main(){
	flag.Parse()
	fmt.Println(*port)
	fmt.Println(*ip)
	fmt.Println(strings.Join(flag.Args(),","))
}

/*
➜  practice_go git:(master) ✗ go run i_flag.go -port 1234 -ip 255.244.244.244 wangkai xua
1234
255.244.244.244
wangkai,xua
*/